//package actuator
package main
//
// actuator
// client side

import ( "crypto/md5" ; "io" ; "os" ; "errors" )
import ( "path/filepath")
import "time"
import "bufio"
import "syscall"
//
import _ "net/http/pprof"
import "net/http"
import "fmt"

// Now it skips symlinks and other shit like a pipes and character devices

type inodes []uint64
type strings []string


type File struct {
    Path  string
    Dir   string
    Sum   []byte
    Type  string
    Inode uint64
}

type Directory struct {

    Path             string
    Inode            uint64
    Files            []*File
    SubDirs          strings
    DiscoveredInodes inodes

}

var is_dir_error = errors.New("is_dir")

var is_not_regular = errors.New("isnt_reg")

var is_not_readable = errors.New("isnt_read")

var ino_not_found = errors.New("ino_not_found")

var dup_inode  = errors.New("dup_inode")



func (array inodes) IncludeValue ( value uint64 ) ( has bool ){


    for i:=range array { if value==array[i] { has=true ; break } }

    return

}

func ( array strings ) IncludeValue ( value string ) (has bool) {


    for i:=range array { if value==array[i] { has=true ; break } }

    return

}



func GetFileIndexNumber(path string)(ino uint64,err error) {

    fi, err:=os.Stat(path)

    if err!=nil {return 0,err}

    stat_interface:=fi.Sys()

    stat_object,found :=stat_interface.(*syscall.Stat_t) // type assert

    if found==false { return 0, ino_not_found  }

    return stat_object.Ino,nil

}



func RegularFileIsReadable (path string) (readable bool) {

    // thanks for "postman" from golang@cjr for this great idea
    // this function is preventing blocking during reading files like a /proc/1/task/1/cwd/proc/kmsg

    file, err := os.Open(path)
      if err!=nil {
        return false
    }
    manage_chn:=make(chan bool,1)
    var content []string
    go ReadFileWithTimeoutControll( file, manage_chn, &content)
    time.Sleep(1 * time.Millisecond)
    select {
        case is_readable:=<-manage_chn:

            if is_readable == true /* true means first line was read  */ { read_is_completed :=<-manage_chn ; if read_is_completed == false /* false means YES  */  {   defer file.Close() ; readable=true  } }

        default:

            file.Close()
            readable=false

    }
    //for i:=range content {
    //    fmt.Printf("%s",content[i])
    //}
    fmt.Printf("\n%s check is done %t\n",path,readable)
    return

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

            if ( !read_start_signal_sent  ) { readable<-true ;  readable<-false } else { readable<-false  }

            return nil
        }
    }
    if ( !read_start_signal_sent )  { readable<-true ; readable<-false  } else { readable<-false  }
    

    return nil
}





func IsEmpty(path string) (empty bool,err error) {

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

    if size==0 { empty=true  } else { empty=false }

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

func Get_mtime(path string)(mtime string,err error) {

    fi, err:=os.Stat(path)

    if err!=nil { return mtime,err }

    mtime_struct:=fi.ModTime()

    return string(mtime_struct.Format("2006-01-02T15:04:05.999999999Z07:00")),nil

}


func ( directory *Directory ) Get_md5_dir (path string) (err error){

    //var dir_struct Directory

    dir, err := os.Open(path)
    directory.Path = path

    //dir_struct.Files = [] File {}

    defer dir.Close()
    if err != nil {
        return  err
    }

    dir_content , err := dir.Readdirnames(-1)

    if err != nil {

        return  err
    }


    // check inode number 
    // prevent looping while discovering subdirectories

    inode,err:=GetFileIndexNumber(path)

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
        err=file_struct.Get_md5_file(path+"/"+dir_content[file])

        if err==nil {
                //var subdir_added bool
                directory.Files=append(directory.Files,file_struct)
                //for i:=range directory.SubDirs { if (directory.SubDirs[i]==path) {subdir_added=true ; break  } }
                //if subdir_added==false { directory.SubDirs=append(directory.SubDirs,path) }
        }

        if err == is_dir_error {

            another_dir := &Directory{}
            another_dir.DiscoveredInodes = directory.DiscoveredInodes
            subdir_path:=path+"/"+dir_content[file]
            err = another_dir.Get_md5_dir( subdir_path )

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

    isdir,err:=IsDir(path)

    if (isdir==true && err==nil ) { return is_dir_error }

    if ( err!=nil ) { return err }

    is_readable := RegularFileIsReadable(path)

    if is_readable==false { /*fmt.Printf("\n<Not readable %s>",path)*/  ;  return is_not_readable }

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

    return nil

}



func main() {

        go func() {
	    fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
        }()


        dir_struct:=&Directory{}
        dir_struct.Get_md5_dir("/tmp/mtime_test")


        counter:=0
        for file := range dir_struct.Files {

            file_struct := dir_struct.Files[file]

            fmt.Printf("Filename: %s MD5Sum:  %x\n",file_struct.Path,file_struct.Sum)
            counter+=1

        }
        fmt.Println(":: mtime ::")
        fmt.Printf("counter %s",counter)

        fmt.Println(Get_mtime("/tmp/does_not_exist"))
        // 
        // test
        file:=&File{}

        fmt.Println("::Checking file::")
        err:=file.Get_md5_file("/proc/1/task/1/cwd/proc/iomem")

        fmt.Printf("Path:%s Sum:%x Dir:%s \n",file.Path,file.Sum,file.Dir)
        fmt.Println(err)

    }
