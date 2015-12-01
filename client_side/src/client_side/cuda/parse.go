package cuda

import "os"
import "bufio"
import "io"
import "strings"
import "fmt"
import "regexp"

var comments         =  []string {`//` , `#`}
var delimiters       =  []string {":", "="}
var enum_delimiters  =  []string {",",";"}
var word_delimiters  =  []string {"-","_"}
var brackets         =  []string {"[","]","<","/>",">","{","}",")","("}
var section_brackets =  []string {"[","]","<","/>",">"}
var quotes           =  [2]string {`"`, "'"}


type Section struct {


}


func Parser(entry string) ( interface{} ) {

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

func GetSeparatorIndexes (entry, sep string) (indexes []int) {

    chars := strings.Split(entry,"")
    for char := range chars {
        if chars[char] == sep {
            indexes=append(indexes, char)
        }
    }
    return indexes
}

func GetWordIndexes (entry string) (indexes []int) {

    reg:=regexp.MustCompile("[[:alnum:]-]+")
    match:= reg.FindAllString(entry, -1)
    match_indexes := reg.FindAllStringIndex(entry, -1)
    for i := range match_indexes{
        fmt.Printf("\n%s>>\n",entry[match_indexes[i][0]:match_indexes[i][1]])
    }
    for i := range match {
        fmt.Printf("\n--%s--\n",match[i])

    }
    return indexes
}

func QuotesParse ( entry string) ( word_set [2]string, complete [2]bool  ) {

    //var single_quotes_count        int32
    //var double_quotes_count        int32
    // var another = [2]int8 {1,0}
    for i:=range quotes {
        quote:=quotes[i]
        if strings.Count(entry, quote)%2 == 0 {
            complete[i] = true
        } else {
            complete[i] = false
        }
        word_set[i] = strings.Replace(entry, quote, "", -888)

    }
    return word_set, complete
}

func EqualParse (entry string) () {



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
