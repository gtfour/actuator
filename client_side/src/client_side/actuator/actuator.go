//package actuator
package main
//
// actuator
// client side

import ( "crypto/md5" ; "io" ; "os" ; "errors" )
import ( "path/filepath")
import "time"
import "bufio"
import "fmt"
//
import _ "net/http/pprof"
import "net/http"

// Now it skips symlinks and other shit like a pipes and character devices


type File struct {
    Path string
    Dir string
    Sum []byte
    Type string
}

type Directory struct {

    Path string
    Files []File
    SubDirs []string

}

var is_dir_error = errors.New("is_dir")

var is_not_regular = errors.New("isnt_reg")

var is_not_readable = errors.New("isnt_read")

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
    return

}

func ReadFileWithTimeoutControll ( file *os.File, readable chan<- bool, content *[]string )(err error){

    buffered_reader:=bufio.NewReader(file)
    eof := false
    for lino := 1; !eof; lino++ {
        //fmt.Printf("linenum %d",lino)
        if lino ==2 {  readable<-true }
        line, err := buffered_reader.ReadString('\n')
        *content=append(*content,line)

            if err == io.EOF {
                err = nil
                eof = true
             } else if err != nil {
            readable<-false
            return err
        }
    }

    readable<-false
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
    if err!=nil {return mtime,err}
    mtime_struct:=fi.ModTime()
    return string(mtime_struct.Format("2006-01-02T15:04:05.999999999Z07:00")),nil

}


func Get_md5_dir(path string)(dir_struct Directory,err error){


    //var dir_struct Directory
    dir, err := os.Open(path)
    dir_struct.Path = path
    //dir_struct.Files = [] File {}

    if err != nil {
        return  dir_struct, err
    }

    dir_content , err := dir.Readdirnames(-1)

    if err != nil {
        return  dir_struct, err
    }

    defer dir.Close()


    for file:= range dir_content{

        file_struct, err := Get_md5_file(path+"/"+dir_content[file])


        if err==nil {

            //go func() { 

                var subdir_added bool
                dir_struct.Files=append(dir_struct.Files,file_struct)
                for i:=range dir_struct.SubDirs { if (dir_struct.SubDirs[i]==path) {subdir_added=true} }
                if subdir_added==false { dir_struct.SubDirs=append(dir_struct.SubDirs,path) }

            //}()
        }

        if err==is_dir_error {

            another_dir_struct, _:=Get_md5_dir(path+"/"+dir_content[file])

            var subdir_added bool

            for i:=range dir_struct.SubDirs { if (dir_struct.SubDirs[i]==(path+"/"+dir_content[file])) {subdir_added=true} }

            if subdir_added==false { dir_struct.SubDirs=append(dir_struct.SubDirs,(path+"/"+dir_content[file])) }

            for another_file:= range another_dir_struct.Files{

               dir_struct.Files=append(dir_struct.Files,another_dir_struct.Files[another_file])

            }
        }

    }
    return dir_struct,nil
    
    // os.Readdirnames

}


func Get_md5_file(path string)(file_struct File, err error){

    IsEmpty(path)
    //
    var result []byte

    file_struct=File{}

    isdir,err:=IsDir(path)



    if (isdir==true && err==nil ) { return file_struct, is_dir_error }

    if ( err!=nil ) { return file_struct, err }


    is_readable := RegularFileIsReadable(path)
    if is_readable==false { fmt.Printf("<Not readable %s>",path)  ;  return file_struct, is_not_readable }
    

    file, err := os.Open(path)

    defer file.Close()

    hash := md5.New()

    fmt.Println("Io copy starting %s",path)

    if _,err = io.Copy(hash, file); err !=nil {

        return file_struct, err
    }

    mdsum:=hash.Sum(result)

    file_struct.Path = path
    file_struct.Sum = mdsum
    file_struct.Dir = filepath.Dir(path)

    return file_struct , nil

}



func main() {

        go func() {
	    fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
        }()


        dir_struct , _ :=Get_md5_dir("/proc/1")



        for file := range dir_struct.Files {

            file_struct := dir_struct.Files[file]

            fmt.Printf("Filename: %s MD5Sum:  %x\n",file_struct.Path,file_struct.Sum)

        }
        fmt.Println(":: mtime ::")
        fmt.Println(Get_mtime("/tmp/does_not_exist"))

    }
