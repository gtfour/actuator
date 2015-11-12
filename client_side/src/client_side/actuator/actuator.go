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
//
//import _ "net/http/pprof"
//import "net/http"
import "fmt"

// Now it skips symlinks and other shit like a pipes and character devices
// Bug with opening /proc/1/task/1/cwd/proc/kcore" still does not fixed !!!

type inodes []uint64
type strings []string

var OPEN_FILE_TIMEOUT time.Duration = 5 // Remember that OPEN_FILE_TIMEOUT digit  is dividing for two parts in RegularFileIsReadable

type Prop struct {

    Inode               uint64
    InoFound            bool
    IsDir               bool
    IsEmpty             bool
    IsReadable          bool
    IsRegular           bool
    Dir                 string
    Mtime               string
    HashSum             string
    HashSumType         string //md5
    HashSumAvailable    bool
    Perm                string
    Owner               string
    OwnerGroup          string
    Size                int64
    DirContent          []string


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

func GetProp (path string) (p Prop,err error){

    p               =  Prop{}
    file, err       :=  os.Open(path)
    defer           file.Close()
    if( err!=nil )  { return p,err  }
    file_stat, err  :=  os.Stat(path)
    //defer           file_stat.Close()
    if( err!=nil )  { return p,err  }

    stat_interface     :=  file_stat.Sys()
    stat_object,found  :=  stat_interface.(*syscall.Stat_t)
    if found==false { p.InoFound=false  } else { p.InoFound=true  ; p.Inode = stat_object.Ino }

    if RegularFileIsReadable(file) == nil {
        p.IsReadable       = true
        p.HashSumAvailable = true
    } else {
        p.IsReadable       = false
        p.HashSumAvailable = false
    }

    file_info,err    :=  file.Stat()
    if err==nil {
        p.Size       =  file_info.Size()
        if p.Size==0 { p.IsEmpty=true  } else { p.IsEmpty = false }

        file_mode :=  file_info.Mode()
        if file_mode.IsDir()==true {
            p.IsDir = true
        } else {
            file_type := string(file_mode.String()[0])
            if ( file_type == "-" ) {  p.IsRegular = true } else { p.IsRegular = false  }
        }
        if p.IsDir == true  {
            content,err  := file.Readdirnames(-1)
            if err == nil { p.DirContent = content }
        }

        mtime_struct  := file_stat.ModTime()
        p.Mtime       =  string(mtime_struct.Format("2006-01-02T15:04:05.999999999Z07:00"))
        if p.HashSumAvailable == false { p.HashSum = p.Mtime  }
    }

    if p.HashSumAvailable == true && p.IsReadable==true && p.IsRegular == true && p.IsDir == false {

        var result []byte
        hash := md5.New()
        if _,err = io.Copy(hash, file); err !=nil {
            p.HashSumAvailable = false
        }else { p.HashSum = string(hash.Sum(result))  }
    }
    p.Dir = filepath.Dir(path)
    return p,nil

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
                            fmt.Printf("\nswitch to mtime\n")
                            file.Close()
                            return Have_to_switch_to_mtime

                    }
            }

        default:
            // In case when we are not recieving second signal in time modification marker getting  method is switching to 'lazy mode'
            // We just getting file modification time
            fmt.Printf("\nswitch to mtime\n")
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


func ( directory *Directory ) Get_md5_dir (path string) (err error){

    //var dir_struct Directory

    
    prop , err := GetProp(path)
    if err != nil {
        return  err
    }

    directory.Path = path
    directory.Prop = &prop



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

        file_struct  :=  &File{}
        file_path    :=  path+"/"+directory.Prop.DirContent[file]
        err=file_struct.GetFileProp( file_path )

        if err==nil || err==Have_to_switch_to_mtime || err == Permission_denied  {
                //var subdir_added bool
                //if err == Have_to_switch_to_mtime {fmt.Printf("\n -- Switched to mtime marker mode -- %s -- \n",path+"/"+dir_content[file])}
                if err == Have_to_switch_to_mtime || err == Permission_denied  {
                    file_struct.Prop.HashSumAvailable = false
                }
                directory.Files=append(directory.Files,file_struct)
                //for i:=range directory.SubDirs { if (directory.SubDirs[i]==path) {subdir_added=true ; break  } }
                //if subdir_added==false { directory.SubDirs=append(directory.SubDirs,path) }
        }

        if err == Is_dir_error {

            another_dir                   :=  &Directory{}
            another_dir.DiscoveredInodes  =   directory.DiscoveredInodes
            subdir_path                   :=  path+"/"+dir_content[file]
            err                           =   another_dir.Get_md5_dir( subdir_path )

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


func (file_struct *File) GetFileProp (path string) (err error){


    prop,err         := GetProp(path)
    file_struct.Prop = &prop
    file_struct.Path =  path

    if err!= nil { return err }

    return nil

}
