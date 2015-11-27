package cuda

import "os"
import "bufio"
import "io"
import "strings"
import "fmt"

var comments         =  []string {`//` , `#`}
var delimiters       =  []string {":", "="}
var enum_delimiters  =  []string {",",";"}
var word_delimiters  =  []string {"-","_"}
var word_quotes      =  []string {`"`, "'"}
var brackets         =  []string {"[","]","<","/>",">","{","}",")","("}
var section_brackets =  []string {"[","]","<","/>",">"}


type Section struct {


}


func Parse(entry string) ( interface{} ) {

    sub_entries      :=  strings.Split(entry," ")
    sub_entries_len  :=  len(sub_entries)
    switch {
        case sub_entries_len == 1:
            fmt.Printf("sub en len is 1")
        case sub_entries_len == 2:
             fmt.Printf("sub en len is 2")




    }

    return nil

}

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
