package main

import "fmt"
import "os"

type Prop struct {

    Inode               uint64
    InoFound            bool
    IsDir               bool
    IsEmpty             bool
    IsReadable          bool
    IsRegular           bool
    Dir                 string
    Mtime               string
    MtimeAvailable      bool
    HashSum             string
    HashSumType         string //md5
    HashSumAvailable    bool
    Perm                string
    Owner               string
    OwnerGroup          string
    Size                int64
    DirContent          []string
    DirContentAvailable bool


}


func GetProp (path string) (p *Prop,err error){


    p               =  &Prop{}
    file, err       :=  os.Open(path)
    defer           file.Close()
    if( err!=nil )  { return p,err  }
    file_stat, err  :=  os.Stat(path)
    //defer           file_stat.Close()
    if( err!=nil )  { return p,err  }

    fmt.Printf("\n -- Getting prop --\n")

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
    fmt.Printf("\n<< Printing error >>\n")
    if err != nil { fmt.Printf("\n error not null %v\n",err) } else { fmt.Printf("\n error is nil  \n")  }
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
            fmt.Printf("\n This dir is Dir \n")
            if err == nil { p.DirContent = content ; p.DirContentAvailable = true } else { p.DirContentAvailable = false }
        }

        mtime_struct     := file_stat.ModTime()
        p.Mtime          =  string(mtime_struct.Format("2006-01-02T15:04:05.999999999Z07:00"))
        p.MtimeAvailable =  true
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

func main() {

test,err := GetProp("/tmp/test")




}

