package main

import "os"
import "fmt"
import "flag"

//var TARGET_FILE string = "/tmp/simple.txt"

func main(){

    in_file_ptr          := flag.String("infile","none","Input file")
    in_file_offset_ptr   := flag.Int64("offset",0,"Input file offset")
    entry_ptr            := flag.String("entry","hello","New entry")
    flag.Parse()
    //out_file_ptr       := flag.String("outfile","out.txt","Out file")
    in_file              := *in_file_ptr
    offset               := *in_file_offset_ptr
    entry                := *entry_ptr

    err:=Replace(in_file, offset, entry)
    fmt.Printf("\nInput file:%v\nEntry:%v\nOffset:%v\nError:%v\n",in_file,entry,offset,err)

}

func Replace(file_path string , offset int64 , new_entry string)(err error){
    f, err := os.OpenFile(file_path , os.O_RDWR, os.ModePerm)
    if err!= nil {
        return err
    }
    f.Seek(offset, os.SEEK_SET)
    f.Write([]byte(new_entry))
    f.Close()
    return err
}
