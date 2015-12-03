package cuda

//import "os"
//import "bufio"
//import "io"
import "strings"
import "fmt"
import "regexp"

var comments                  =  []string {`//` , `#`}
var delimiters                =  []string {":", "="}
var enum_delimiters           =  []string {",",";"}
var word_delimiters           =  []string {"-","_"}
var brackets                  =  []string {"[","]","<","/>",">","{","}",")","("}
var section_brackets_square   =  [2]string {"[","]"}
var section_brackets_triangle =  [3]string {"<",">","/>"}
var quotes                    =  [3]string {`"`, "'", "`"}
// ⣳ 


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

func GetQuotesIndexes (line string) ( indexes []int) {

    lineAsArray:=strings.Split(line,"")
    for char:= range lineAsArray {
        for q:=range quotes {
            quote:=quotes[q]
            if lineAsArray[char] == quote {
                indexes = append(indexes, char)
            }
        }
    }
    return
}


func GroupByQuotes (lineAsArray []string, quotes_indexes []int) (quotes_pairs [][]int) {

    pending := make(map[string]int)
    for char:= range quotes_indexes {
        quote:=quotes_indexes[char]
        quoteInPending:=false
        quote_value:=lineAsArray[quote]
        for key, _ := range pending { if quote_value == key { quoteInPending = true } }
        if quoteInPending == false {
            pending[quote_value] = quote
        } else {
            var new_pair = []int  { pending[quote_value],quote } // 1 added for ignoring quotes  themselves
            quotes_pairs = append(quotes_pairs, new_pair)
            delete(pending, quote_value)
        }
    }

    return
}

func QuotesSpreading ( entry string) ( word_set [3]string, complete [3]bool  ) {

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

func EqualSignEscape (entry string) (words_indexes [][]int) {

    wordSplittedByEqualSign:= strings.Split(entry, "=")
    offset:=0
    for i := range wordSplittedByEqualSign {
        word:=wordSplittedByEqualSign[i]
        word_index:=strings.Index(entry[offset:], word)
        var new_pair =  []int {(word_index+offset),(len(word)+offset)}
        words_indexes=append(words_indexes, new_pair)
        offset+=len(word)+1 // 1 is equal sign
    }
    return words_indexes

}

func SectionNameEscape ( entry string ) ( name, tag string , section_type int ) {

    square         :=0
    //triangle       :=1
    //curly          :=2

    opening        :=0
    closing        :=1
    //closing_slashed:=2


    if strings.Index(entry,section_brackets_square[opening]) == 0 && strings.Index(entry,section_brackets_square[closing]) == (len(entry)-1) {

        return name, "", square

    }
    //opening        :=0
    //closing        :=1
    //closing_slashed:=2
    return

}

func DebugCharCounter (line  string) (heads, foots []string) {

    lineAsArray:=strings.Split(line,"")

    head:=""
    foot:=""
    for i:=0 ; i<len(lineAsArray) ; i++ {

        delim:=""
        delim_template:=" %s%s "
        for z:=2;z<=len(fmt.Sprint(i));z++ {delim+=" "}
        head+=fmt.Sprintf(delim_template, lineAsArray[i], delim)
        foot+=fmt.Sprintf("|%d|",i)
        if (i%10==0)&&(i!=0) || (i+1==len(line))  { heads=append(heads,head) ; foots=append(foots,foot) ; head="" ; foot="" }

    }
    return heads, foots
}

func RemoveSpaces(entry string, remove_type int)([]int) {


    leading    :=0
    closing    :=1
    both       :=2


    lineAsArray:=strings.Split(entry, "")
    leadingChar:=0
    closingChar:=len(entry)-1
    leadReady:=false
    closeReady:=false
    for char := range lineAsArray {
        if (remove_type==leading || remove_type==both) && lineAsArray[char] != " " {
                if leadReady != true {

                    leadingChar=char
                    leadReady=true
                    if remove_type==leading { break }

                }
        }

        closing_char:=len(lineAsArray)-1-char

        if (remove_type==closing || remove_type==both) && (lineAsArray[closing_char]!=" ")  {
                 if closeReady != true {
                     closingChar=closing_char+1
                     closeReady=true
                     if remove_type==closing { break }
                 }

        }
        if closeReady && leadReady { break }
    }
    if closingChar<leadingChar { return []int {0,0} }
    return []int {leadingChar,closingChar}
}

/*func ParseFile( filename string ) ( err error ) {

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

}*/
