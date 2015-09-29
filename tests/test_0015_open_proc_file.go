package main

import "fmt"
import "os"
import "bufio"
import "io"
import "time"
//import "io/ioutil"

func main() {

    //filename:="/proc/1/task/1/cwd/proc/kmsg"      //empty not readable file
    filename:="/proc/1/task/1/cwd/proc/partitions"  //empty readable file

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

            if is_readable == true { is_completed:=<-readable ; if is_completed==false { defer file.Close()  } }

        default:

            fmt.Println("file is not readable")
            file.Close()

    }
    for i:=range content {

        fmt.Printf("%s",content[i])

    }
}



func read_file( file *os.File, readable chan<- bool, content *[]string )(err error){

    buffered_reader:=bufio.NewReader(file)
    eof := false
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
    return nil
}
