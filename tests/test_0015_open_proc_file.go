package main

import "fmt"
import "os"
import "bufio"
import "io"
import "time"
//import "io/ioutil"

func main() {

    filename:="/proc/1/task/1/cwd/proc/kmsg"
    //filename:="/proc/1/task/1/cwd/proc/partitions"
    file, err := os.Open(filename)

      if err!=nil {
        return
    }

    buffered_reader:=bufio.NewReader(file)

    eof := false

    mng:=make


    for lino := 1; !eof; lino++ {
        fmt.Println("for phase")

        select{


        case line, err := buffered_reader.ReadString('\n'):
            fmt.Println("--")
            fmt.Println(line)
            fmt.Println("--")
            if err == io.EOF {

                err = nil
                eof = true

             } else if err != nil {
        default:
             

            return

        }
    }
//content, _ := ioutil.ReadFile(filename)
//fmt.Println(content)
}

func timer(mng chan bool )(){

    time.Sleep(1000 * time.Millisecond)
    mng <- true

}
