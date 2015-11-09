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
// import "fmt"

// Now it skips symlinks and other shit like a pipes and character devices
// Bug with opening /proc/1/task/1/cwd/proc/kcore" still does not fixed !!!

type inodes []uint64
type strings []string

var OPEN_FILE_TIMEOUT time.Duration = 5


type File struct {
    Path                          string
    Dir                           string
    Sum                           []byte
    Type                          string
    Inode                         uint64
    MarkerGetttingModeIsMtime     bool // in case when file size is very big and opening takes more than OPEN_FILE_TIMEOUT
                                       // Remember that OPEN_FILE_TIMEOUT digit  is dividing for two parts in RegularFileIsReadable
}

type Directory struct {

    Path             string
    Inode            uint64
    //Dir              string
    Files            []*File
    SubDirs          strings
    DiscoveredInodes inodes

}

var   cant_open_file          =  errors.New("cant_open")
var   Is_dir_error            =  errors.New("is_dir")
var   is_not_regular          =  errors.New("isnt_reg")
var   is_not_readable         =  errors.New("isnt_read")
var   ino_not_found           =  errors.New("ino_not_found")
var   dup_inode               =  errors.New("dup_inode")
var   Have_to_switch_to_mtime =  errors.New("switch_to_mtime")
                                                            // we are switching to MTIME method // 



func (array inodes) IncludeValue ( value uint64 ) ( includes bool ){

    for i:=range array { if value==array[i] { includes=true ; break } }
    return

}

func ( array strings ) IncludeValue ( value string ) (includes bool) {


    for i:=range array { if value==array[i] { includes=true ; break } }
    return

}



func GetFileIndexNumber(path string)(ino uint64,err error) {

    fi, err:=os.Stat(path)

    if err != nil { return 0, err }

    stat_interface    := fi.Sys()

    stat_object,found := stat_interface.(*syscall.Stat_t) // type assert

    if found==false { return 0, ino_not_found  }

    return stat_object.Ino, nil

}



func RegularFileIsReadable (path string) (err error) {

    // thanks for "postman" from golang@cjr for this great idea
    // this function is preventing blocking during reading files like a /proc/1/task/1/cwd/proc/kmsg

    file, err := os.Open(path)

    defer file.Close()

    if err!=nil {

        return cant_open_file

    }


    //empty,err:= IsEmpty(path)
    //fmt.Printf("File is empty  %t  Err:  %s",empty,err.Error())

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
                            file.Close()
                            return Have_to_switch_to_mtime

                    }
            }

        default:
            // In case when we are not recieving second signal in time modification marker getting  method is switching to 'lazy mode'
            // We just getting file modification time
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



func IsEmpty(path string) (empty bool, err error) {

    file, err := os.Open(path)

    defer file.Close()

    if err != nil {

        return false, err

    }

    file_info , err := file.Stat()

    if err != nil {

        return false,err

    }

    size :=  file_info.Size()

    if size==0 { empty=true  } else { empty=false /* ; fmt.Printf("\nFile path: %s  size : %d \n",path,size)*/ }

    return
}

func IsDir(path string)(isdir bool,err error) {

    file, err := os.Open(path)

    //fmt.Println(path)

    defer file.Close()

    if err != nil {

        return false, err

    }

    file_info , err := file.Stat()

    if err != nil {

        return false,err

    }

    file_mode :=  file_info.Mode()

    if file_mode.IsDir()==true {

        isdir = true

    } else {


        file_type := string(file_mode.String()[0])

        if ( file_type == "-" ) {  err = nil } else { err = is_not_regular }

    }

    return
}

func Get_mtime( path string )( mtime string, err error ) {

    fi, err:=os.Stat(path)

    if err!=nil { return mtime,err }

    mtime_struct:=fi.ModTime()

    return string(mtime_struct.Format("2006-01-02T15:04:05.999999999Z07:00")),nil

}


func ( directory *Directory ) Get_md5_dir (path string) (err error){

    //var dir_struct Directory

    dir, err := os.Open(path)
    defer dir.Close()

    directory.Path = path

    if err != nil {
        return  err
    }

    dir_content , err := dir.Readdirnames(-1)

    if err != nil {
        return  err
    }


    // check inode number 
    // prevent looping while discovering subdirectories

    //directory.Dir = filepath.Dir(path)

    inode,err := GetFileIndexNumber(path)

    if err!=nil { return ino_not_found }

    directory.Inode = inode

    if ( directory.DiscoveredInodes.IncludeValue( directory.Inode ) ) {


                return dup_inode } else {

                directory.DiscoveredInodes = append( directory.DiscoveredInodes, directory.Inode )

    }

    var subdir_added bool

    for i:=range directory.SubDirs { if (directory.SubDirs[i]==path) { subdir_added=true ; break  } }

    if subdir_added==false { directory.SubDirs=append(directory.SubDirs, path) }


    //

    for file:= range dir_content{

        file_struct:=&File{}
        // fmt.Printf("\n---Get md5 : %s \n",path+"/"+dir_content[file]) // /proc/1/task/1/cwd/proc/kcore
        err=file_struct.Get_md5_file(path+"/"+dir_content[file])

        if err==nil || err==Have_to_switch_to_mtime  {
                //var subdir_added bool
                //if err == Have_to_switch_to_mtime {fmt.Printf("\n -- Switched to mtime marker mode -- %s -- \n",path+"/"+dir_content[file])}
                file_struct.MarkerGetttingModeIsMtime = true
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


func (file_struct *File) Get_md5_file (path string) (err error){

    //IsEmpty(path)
    //
    var result []byte

    //<check's 

    isdir, err := IsDir(path)

    if ( isdir==true && err==nil ) { return Is_dir_error }

    if ( err!=nil ) { return err }

    err = RegularFileIsReadable( path ) // check was failed by timeout controll . It fails when opens /proc/kmsg or other strange files

    if err == Have_to_switch_to_mtime {

        //fmt.Printf("\n == Switched to mtime marker == \n")
        file_struct.Path   =  path
        mtime,err          :=  Get_mtime(path)
        //fmt.Printf("\n == Get_mtime  Error %v== \n",err)
        if err!=nil        { return err }
        file_struct.Sum    = []byte(mtime)
        file_struct.Dir    = filepath.Dir(path)
        return             Have_to_switch_to_mtime

    } else if err!= nil { return err }

    //if is_readable==false { /*fmt.Printf("\n<Not readable %s>",path)  ;*/  return is_not_readable }

    // check's>

    new_file, err := os.Open(path)

    defer new_file.Close()

    hash := md5.New()

    //fmt.Printf("\nIo copy starting %s",path)

    if _,err = io.Copy(hash, new_file); err !=nil {

        return err
    }

    mdsum:=hash.Sum(result)

    file_struct.Path = path
    file_struct.Sum = mdsum
    file_struct.Dir = filepath.Dir(path)

    //fmt.Printf("\nIo copy finished  %s",path)

    return nil

}
