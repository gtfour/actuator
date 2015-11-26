package cuda

import "os"
import "strings"
import "bufio"
import "io"

func ParseFile( filename string ) ( err error ) {

    file, err   := os.Open(filename)
    if err!=nil {

        return err
    }
    buffered_reader:=bufio.NewReader(file)
    eof := false

    for lino := 1; !eof; lino++ {


        line, err := buffered_reader.ReadString('\n')

        if err == io.EOF {
            err = nil
            eof = true
        } else if err != nil {
            return err
        }

        if ( strings.HasPrefix(line, "Package") || strings.HasPrefix(line, "Status") ||  strings.HasPrefix(line, "Architecture") || strings.HasPrefix(line, "Version")){
        }



    }
    return nil

}
