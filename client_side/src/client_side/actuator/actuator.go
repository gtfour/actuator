package actuator 
//
// actuator
// client side

import ( "fmt" ; "crypto/md5" ; "io" ; "os" ; "errors" )
//import ( "path/filepath")


type File struct {
    Path string
    Sum []byte
}

type Directory struct {

    Path string
    Files []File

}

var is_dir_error = errors.New("is_dir")

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


    for file:= range dir_content{

        file_struct, err := Get_md5_file(path+"/"+dir_content[file])

        if err==nil {

            //go func() { 

                dir_struct.Files=append(dir_struct.Files,file_struct)

            //}()
        }

        if err==is_dir_error{

            another_dir_struct, _:=Get_md5_dir(path+"/"+dir_content[file])

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
    file, err := os.Open(path)
    if err != nil {

        return file_struct, err 

    }
    file_info , err := file.Stat()
    if err != nil { 
        
        return file_struct, err

    }
    file_mode :=  file_info.Mode()
    if file_mode.IsDir()==true {
  
        return file_struct, is_dir_error
 
    }
    //if mode.IsDir() {
    //    fmt.Println("FIle is directory")
    //}
    defer file.Close()
    hash := md5.New()
    if _,err = io.Copy(hash, file); err !=nil {

        return file_struct, err
    }
    mdsum:=hash.Sum(result)

    file_struct.Path = path
    file_struct.Sum = mdsum

    return file_struct , nil

}



func main() {

        dir_struct , _ :=Get_md5_dir("/var/lib/rpm")

        for file := range dir_struct.Files {

            file_struct := dir_struct.Files[file]

            fmt.Printf("Filename: %s MD5Sum:  %x\n",file_struct.Path,file_struct.Sum)

        }

    }
