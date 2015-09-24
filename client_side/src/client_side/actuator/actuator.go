package main
//package actuator 
//
// actuator
// client side

import ( "crypto/md5" ; "io" ; "os" ; "errors" )
import ( "path/filepath")
import "fmt"


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

func IsDir(path string)(isdir bool,err error) {

    fmt.Printf("Path: %s\n",path)
    file, err := os.Open(path)
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

        var file_type string

        if file_type, ok := file_mode.String()[0] ; ok != false { return false, is_not_regular }

        fmt.Printf("file mode: %s ", file_mode.String()[0])

        if file_mode.IsRegular()==false { fmt.Printf("%s Is not  regular ", path) }

        isdir = false

        if ((file_type == "-") && (isdir==false))||((file_type=="L") && (isdir==false)) { return false, is_not_regular }

    }

    return
}

func Get_mtime(path string)(mtime string) {

    fi, _:=os.Stat(path)
    mtime_struct:=fi.ModTime()
    return string(mtime_struct.Format("2006-01-02T15:04:05.999999999Z07:00"))

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

        if err==is_dir_error{

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

    //
    var result []byte
    file_struct=File{}

    if isdir,err:=IsDir(path) ; (isdir==true && err==nil ) {return  file_struct, is_dir_error } 
    if (err!=nil ) {return  file_struct, err }

    file, _:= os.Open(path)

    defer file.Close()

    hash := md5.New()

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

        dir_struct , _ :=Get_md5_dir("/tmp/test")

        for file := range dir_struct.Files {

            file_struct := dir_struct.Files[file]

            fmt.Printf("Filename: %s MD5Sum:  %x\n",file_struct.Path,file_struct.Sum)

        }
        fmt.Println(":: mtime ::")
        fmt.Println(Get_mtime("/tmp/does_not_exist"))

    }
