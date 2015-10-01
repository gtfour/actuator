package main

import "fmt"
import "os"
import "bufio"
import "io"
import "time"
//import "io/ioutil"

func main() {

    //filename:="/proc/1/task/1/cwd/proc/kmsg"      //empty not readable file
    //filename:="/proc/1/task/1/cwd/proc/partitions"  //empty readable file
    filename:="/tmp/test222/test333/hello.txt" //simple empty file

    //
    file, err := os.Open(filename)
      if err!=nil {
        return
    }
    readable:=make(chan bool,1)
    var content []string
    go read_file(file,readable,&content)
    time.Sleep(1 * time.Millisecond)
    select {
        case is_readable:=<-readable:

            fmt.Println("first signal recieved:")
            if is_readable == true { is_completed:=<-readable ; if is_completed==false {  fmt.Println("file is readable +")  ; defer file.Close()  } }

        default:

            fmt.Println("file is not readable -")
            file.Close()

    }
    for i:=range content {

        fmt.Printf("%s",content[i])

    }
}



func read_file( file *os.File, readable chan<- bool, content *[]string )(err error){

    buffered_reader:=bufio.NewReader(file)
    eof := false
    fmt.Println("Reading the file")
    for lino := 1; !eof; lino++ {
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
    fmt.Println("Finishing")
    return nil
}
