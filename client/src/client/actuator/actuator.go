package actuator
//package main
//
// actuator
// client side

import ( "crypto/md5" ; "io" ; "os" ; "errors" )
import ( "path/filepath")
import "time"
import "bufio"
import "syscall"
import "reflect"
import "client/evebridge"
//
//import _ "net/http/pprof"
//import "net/http"
import "fmt"

// Now it skips symlinks and other shit like a pipes and character devices
// Bug with opening /proc/1/task/1/cwd/proc/kcore" still does not fixed !!!

type inodes []uint64
type strings []string

var OPEN_FILE_TIMEOUT time.Duration = 10 // Remember that OPEN_FILE_TIMEOUT digit  is dividing for two parts in RegularFileIsReadable
var LAZY_OPENING_MODE int = 01
var SAFE_OPENING_MODE int = 02

//type CompNote struct {

//    Field    string
//    Before   string
//    After    string


//}

type Prop struct {
    Inode               uint64
    InoFound            bool
    IsDir               bool
    IsEmpty             bool
    IsReadable          bool
    IsRegular           bool
    Type                string
    Dir                 string
    Mtime               string
    MtimeAvailable      bool
    HashSum             string
    HashSumType         string //md5
    HashSumAvailable    bool
    Perm                string
    Uid                 uint32
    Gid                 uint32
    Owner               string
    OwnerGroup          string
    Size                int64
    DirContent          []string `ignore`
    DirContentAvailable bool
    Error               bool
    Fd                  *os.File `ignore`
    FdCheck             *os.File `ignore`
    Ready               bool     `ignore`
}


type File struct {
    Path  string
    Dir   string
    Prop  *Prop
}

type Directory struct {

    Path             string
    Inode            uint64
    Files            []*File
    SubDirs          strings
    DiscoveredInodes inodes
    Prop             *Prop

}

var   cant_open_file          =  errors.New("cant_open")
var   Is_dir_error            =  errors.New("is_dir")
var   is_not_regular          =  errors.New("isnt_reg")
var   is_not_readable         =  errors.New("isnt_read")
var   ino_not_found           =  errors.New("ino_not_found")
var   dup_inode               =  errors.New("dup_inode")
var   Have_to_switch_to_mtime =  errors.New("switch_to_mtime")
var   Permission_denied       =  os.ErrPermission
                                                            // we are switching to MTIME method // 



func (array inodes) IncludeValue ( value uint64 ) ( includes bool ){

    for i:=range array { if value==array[i] { includes=true ; break } }
    return

}

func ( array strings ) IncludeValue ( value string ) (includes bool) {


    for i:=range array { if value==array[i] { includes=true ; break } }
    return

}

func GetProp (path string, mode int) (p *Prop){

    p                   =  &Prop{}
    file, err           := os.Open(path)
    file_same, err_same := os.Open(path)
    p.Fd = file
    p.FdCheck = file_same
    defer file.Close()
    defer file_same.Close()
    if( err!=nil || err_same!=nil )  {  p.Error = true  ; return p  }
    file_stat, err  :=  os.Stat(path)
    if( err!=nil )  {  p.Error = true  ; return p  }

    stat_interface     :=  file_stat.Sys()
    stat_object,found  :=  stat_interface.(*syscall.Stat_t)
    if found==false { p.InoFound=false  } else { p.InoFound=true  ; p.Inode = stat_object.Ino }
    p.Uid  = stat_object.Uid
    p.Gid  = stat_object.Gid

    file_info,err    :=  file.Stat()
    if err==nil {
        p.Size       =  file_info.Size()
        if p.Size==0 { p.IsEmpty=true  } else { p.IsEmpty = false }

        file_mode :=  file_info.Mode()
        if file_mode.IsDir()==true {
            p.IsDir = true
        } else {
            file_type := string(file_mode.String()[0])
            p.Type = file_type
            if mode == SAFE_OPENING_MODE {
                if RegularFileIsReadable(file_same) == nil {
                    p.IsReadable       = true
                    p.HashSumAvailable = true
                } else {
                    p.IsReadable       = false
                    p.HashSumAvailable = false
                }
            // new code
            } else {
                p.IsReadable       = true
                p.HashSumAvailable = true
            }
            // new code
            if ( file_type == "-" ) {  p.IsRegular = true } else { p.IsRegular = false  }
        }
        if p.IsDir == true  {
            content,err  := file.Readdirnames(-1)
            //fmt.Printf("\n This dir is Dir \n")
            if err == nil { p.DirContent = content ; p.DirContentAvailable = true } else { p.DirContentAvailable = false }
        }
        p.Perm           =  string(file_mode.Perm())

        mtime_struct     := file_stat.ModTime()
        p.Mtime          =  string(mtime_struct.Format("2006-01-02T15:04:05.999999999Z07:00"))
        p.MtimeAvailable =  true
        //if p.HashSumAvailable == false { p.HashSum = p.Mtime  }
    }

    if p.HashSumAvailable == true && p.IsReadable==true && p.IsRegular == true && p.IsDir == false {

        var result []byte
        hash := md5.New()
        if _,err = io.Copy(hash, file); err !=nil {
            p.HashSumAvailable = false
        }else { p.HashSum = string(hash.Sum(result))  }
    }
    p.Dir = filepath.Dir(path)
    return p

}




func RegularFileIsReadable (file *os.File) (err error) {

    // thanks for "postman" from golang@cjr for this great idea
    // this function is preventing blocking during reading files like a /proc/1/task/1/cwd/proc/kmsg


    manage_chn:=make(chan bool,1)

    var content []string

    go ReadFileWithTimeoutControll( file, manage_chn, &content)

    first_timeout_period :=OPEN_FILE_TIMEOUT/2
    time.Sleep(first_timeout_period * time.Millisecond)

    select {
        case <-manage_chn:
            select {
                case <-manage_chn:
                    defer file.Close()
                    return nil
                default:
                    time.Sleep((OPEN_FILE_TIMEOUT-first_timeout_period) * time.Millisecond)
                    select {
                        case <-manage_chn:
                            defer file.Close()
                            return nil
                        default:
                            // In case when we are not recieving second signal in time modification marker getting  method is switching to 'lazy mode'
                            // We just getting file modification time
                            //fmt.Printf("\nswitch to mtime\n")
                            file.Close()
                            return Have_to_switch_to_mtime

                    }
            }

        default:
            // In case when we are not recieving second signal in time modification marker getting  method is switching to 'lazy mode'
            // We just getting file modification time
            //fmt.Printf("\nswitch to mtime\n")
            file.Close()
            // set check method to GetMtime
            return Have_to_switch_to_mtime
    }

    return nil

}


func ReadFileWithTimeoutControll ( file *os.File, readable chan<- bool, content *[]string )(err error){

    buffered_reader:=bufio.NewReader(file)

    eof := false

    var read_start_signal_sent bool

    for lino := 1; !eof; lino++ {
        if lino ==2 {  readable<-true ; read_start_signal_sent=true  }
        line, err := buffered_reader.ReadString('\n')
        *content=append(*content,line)

        if err == io.EOF {
            err = nil
            eof = true
        } else if err != nil {

            if ( !read_start_signal_sent  ) {
                readable<-true
                readable<-false } else { readable<-false  }
            return nil
        }
    }

    if ( !read_start_signal_sent )  { readable<-true ; readable<-false  } else { readable<-false  }

    return nil
}


func ( directory *Directory ) GetHashSumDir (path string, mode int) (err error){

    //var dir_struct Directory
    //fmt.Println("\n == Starting  ==\n")
    prop  := GetProp(path, mode)
    if prop.Error == true {
        //fmt.Println("\n Exiting  \n")
        return  err
    }

    //fmt.Printf("\n Dir content is available : %t  \n", prop.DirContentAvailable)

    directory.Path = path
    directory.Prop = prop



    // check inode number 
    // prevent looping while discovering subdirectories

    //directory.Dir = filepath.Dir(path)


    if directory.Prop.InoFound == false { return ino_not_found }

    if ( directory.DiscoveredInodes.IncludeValue( directory.Prop.Inode ) ) {


                return dup_inode } else {

                directory.DiscoveredInodes = append( directory.DiscoveredInodes, directory.Prop.Inode )

    }

    var subdir_added bool

    for i:=range directory.SubDirs { if (directory.SubDirs[i]==path) { subdir_added=true ; break  } }

    if subdir_added==false { directory.SubDirs=append(directory.SubDirs, path) }


    //

    for file:= range directory.Prop.DirContent{

        fstruct            :=  &File{}
        file_path          :=  path+"/"+directory.Prop.DirContent[file]
        fstruct.Path       =   file_path
        fstruct.Prop   =   GetProp( file_path, mode )
        if fstruct.Prop.Error == true  { continue }

        //fmt.Println("\n :Debug:  \n")
        //fmt.Printf("\n%s\n",file_path)
        //fmt.Println("\n : : \n")

        if fstruct.Prop.IsRegular == true && fstruct.Prop.MtimeAvailable == true && fstruct.Prop.IsDir == false  {
                directory.Files=append(directory.Files,fstruct)
        }

        if fstruct.Prop.IsDir == true && fstruct.Prop.DirContentAvailable == true  {

            another_dir                   :=  &Directory{}
            another_dir.DiscoveredInodes  =   directory.DiscoveredInodes
            subdir_path                   :=  path+"/"+directory.Prop.DirContent[file]
            err                           =   another_dir.GetHashSumDir( subdir_path , mode)

            if (err!=nil) { continue }

            //var subdir_added bool
            //for i:=range directory.SubDirs { if ( directory.SubDirs[i] == subdir_path ) {subdir_added=true ; break  } }
            //if subdir_added==false { directory.SubDirs = append(directory.SubDirs, subdir_path ) }

            for another_file:= range another_dir.Files {

                directory.Files = append(directory.Files, another_dir.Files[another_file])

            }

            for subdir := range another_dir.SubDirs {

                directory.SubDirs = append( directory.SubDirs, another_dir.SubDirs[subdir] )

          }
        }
    }
    return nil
    // os.Readdirnames
}

func CompareProp(old_prop,new_prop *Prop, path string)(cnotes evebridge.CompNotes) {


    valueOld:=reflect.ValueOf(old_prop).Elem()
    valueNew:=reflect.ValueOf(new_prop).Elem()

    //fmt.Printf("\n Compare prop: %s %s %s\n",path,valueOld.Kind(),valueNew.Kind())
    //fmt.Printf("\n"+fmt.Sprint(valueOld.Kind())+" -- "+fmt.Sprint(valueNew.Kind())+"\n")
    if (new_prop.IsDir == true ) {
        cnotes.SourceType = "dir"
    } else if ( new_prop.IsRegular == true ) {
        cnotes.SourceType = "file"
    }

    field:=reflect.TypeOf(old_prop).Elem()

    old_field_count := valueOld.NumField()

    for i := 0; i <= old_field_count-1; i++  {

        if string(field.Field(i).Tag)!="ignore" &&  fmt.Sprint(valueOld.Field(i).Interface())!=fmt.Sprint(valueNew.Field(i).Interface()) {


             cnote:=evebridge.CompNote{Field:field.Field(i).Name,Before:fmt.Sprint(valueOld.Field(i).Interface()),After:fmt.Sprint(valueNew.Field(i).Interface())}
             cnotes.List=append(cnotes.List, cnote)


        }

    }
    cnotes.Path = path
    return cnotes
}

func Initial(prop *Prop, path string) (cnotes evebridge.CompNotes) {

    value:=reflect.ValueOf(prop).Elem()
    fields_count := value.NumField()
    field:=reflect.TypeOf(prop).Elem()

    for i := 0; i <= fields_count-1; i++  {
        cnote:=evebridge.CompNote{ Field:field.Field(i).Name,
                                   Before:"",
                                   After:fmt.Sprint(value.Field(i).Interface()) }
        cnotes.List=append(cnotes.List, cnote)
    }
    cnotes.Path = path
    return cnotes
}


