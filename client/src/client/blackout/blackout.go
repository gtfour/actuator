package blackout

import "bytes"
import "fmt"
import "log"
import "os/exec"
import "strings"
import "client/cuda"



func Blackout() {
        cmd := exec.Command("ifconfig","eth0")
        cmd.Stdin = strings.NewReader("some input")
        var out bytes.Buffer
        cmd.Stdout = &out
        err := cmd.Run()
        if err != nil {
                log.Fatal(err)
        }
        lines:= strings.Split(out.String(),"\n")
        for i:=range lines {
            line:=lines[i]
            cuda.DebugPrintCharCounter(line)
            lineAsArray := strings.Split(line,"")
            delims,data:=cuda.GetIndexes(lineAsArray)
            fmt.Printf("\nBefore delims: %v\n data: %v \n" , delims , data)
            delims,data=cuda.PathFilter(lineAsArray,delims,data)
            fmt.Printf("\nAfter delims: %v\n data: %v \n" , delims , data)

        }
}

