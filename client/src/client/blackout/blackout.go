package blackout

import "bytes"
//import "fmt"
import "log"
import "strings"
import "os/exec"
import "encoding/json"
//import "client/cuda"

import "client/activa"



func MakeBounce(motion *activa.Motion)(*Bounce, error) {

    data:=motion.MotionData
    var b Bounce
    err    := json.Unmarshal(data, &b)
    if err == nil {
        return &b,nil
    } else {
        return nil,err
    }
}

func(b *Bounce)Trail(motion *activa.Motion)(error){
    return nil
}


func Blackout() {
        cmd := exec.Command("ifconfig","-a")
        cmd.Stdin = strings.NewReader("n\n")
        var out bytes.Buffer
        cmd.Stdout = &out
        err := cmd.Run()
        if err != nil {
                log.Fatal(err)
        }
        lines:= strings.Split(out.String(),"\n")
        for _=range lines {
            //line:=lines[i]
            //cuda.DebugPrintCharCounter(line)
            /*lineAsArray := strings.Split(line,"")
            delims,data:=cuda.GetIndexes(lineAsArray)
            fmt.Printf("\nBefore delims: %v\n data: %v \n" , delims , data)
            delims,data=cuda.PathFilter(lineAsArray,delims,data)
            fmt.Printf("\nAfter delims: %v\n data: %v \n" , delims , data)
            */
        }
}
