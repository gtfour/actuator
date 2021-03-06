package blacktop

import "io"
import "os"
import "bufio"
import "encoding/json"

import "client/activa"


func MakeRoller(motion *activa.Motion)(*Roller,error) {

    data:=motion.MotionData
    var r Roller
    err    := json.Unmarshal(data, &r)
    if err == nil {
        return &r,nil
    } else {
        return nil,err
    }
}


func (r *Roller)LayDown(motion *activa.Motion)(error) {

    return nil



}


func ReadFileLines(path string)(lines []string ,err error) {

    inFile, err := os.Open(path)
    if err != nil {
        return lines, err
    }
    defer inFile.Close()
    reader := bufio.NewReader(inFile)

    lines=make([]string,0)
    eof := false
    for !eof {
        var line string
        line, err = reader.ReadString('\n')
        if err == io.EOF {
            err = nil   // io.EOF isn't really an error
            eof = true  // this will end the loop at the next iteration
            break
        } else if err != nil {
            return lines,err  // finish immediately for real errors
        }
        lines = append(lines, line)
    }
    return lines, err
}
