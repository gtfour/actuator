package cuda

import "strings"
import "fmt"
import "regexp"

var comments                  =  []string {`//` , `#`}
var delimiters                =  []string {":", "="}
var sequence                  =  []string {",",";"}
var word_delimiters           =  []string {"-","_"}
var brackets                  =  []string {"[","]","<","/>",">","{","}",")","("}
var section_brackets_square   =  [2]string {"[","]"}
var section_brackets_triangle =  [3]string {"<",">","</"}
var section_brackets_curly    =  [2]string {"{","}"}
// TODO: add ident section type . Example ifconfig and dmidecode

var LEADING int                  = 0
var CLOSING int                  = 1
var BOTH    int                  = 2

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

func Escape_Spaces (oldlineAsArray []string) (indexes [][]int) {
    // 
    // Duplicate spaces will be present as one
    //
    lineAsArray      := ReplaceTabsToSpaces(oldlineAsArray)
    var pair         =  []int { -1, -1 }
    for i := range lineAsArray {
        char:=lineAsArray[i]
        if char == " " && pair[0]!=-1 {
            pair[1] = i-1
            indexes = append( indexes, []int {pair[0], pair[1]})
            pair = []int { -1, -1 }
        } else if pair[0] == -1 && char != " " {
            pair[0] = i
        }
        if i == len(lineAsArray)-1 && pair[0]!=-1 {
            indexes = append( indexes, []int {pair[0], i})
        }
    }
    return indexes
}

func IsComment(entry string) (comment bool) {

    for i:= range comments {
        if strings.HasPrefix(entry, comments[i]) == true {
            return true
        }
    }
    return false
}

func ReplaceTabsToSpaces ( lineAsArray []string ) ( newlineAsArray []string ) {

    entry:=strings.Join(lineAsArray, "")
    new_entry:=strings.Replace(entry, "	", " ", -1)
    return strings.Split(new_entry,"")

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

func GetQuotesIndexes (lineAsArray []string) ( indexes []int) {

    var quotes                    =  [3]string {`"`, "'", "`"}

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


func GroupByQuotes (lineAsArray []string) (quotes_pairs [][]int) {

    quotes_indexes := GetQuotesIndexes(lineAsArray)

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

func EscapeDoubleSign_functionBuilder( sign_start, sign_end string   ) ( doublesign_escaper  func ( lineAsArray []string ) ([][]int ) ) {

    // returns function to find single pair of text between specialized sign
    doublesign_escaper = func(lineAsArray []string) ( indexes [][]int) {
        pair := []int {-1,-1}
        for i:= range lineAsArray {
            char:=lineAsArray[i]
            if char == sign_start && pair[0] == -1  {
                pair[0] = i+1
                continue
            }
            if char == sign_end && pair[1] == -1  {
                pair[1] = i-1
                continue
            }
        }
        indexes=append(indexes, pair)
        return indexes
    }
    return doublesign_escaper
}

func EscapeSingleSign_functionBuilder( sign  string   ) ( singlesign_escaper  func ( lineAsArray []string ) ([][]int ) ) {

    // returns function to find single pair of text between specialized sign
    singlesign_escaper = func(lineAsArray []string) ( indexes [][]int) {

    entry:=strings.Join(lineAsArray, "")
    wordSplittedByEqualSign:= strings.Split(entry, sign)
    offset:=0
    for i := range wordSplittedByEqualSign {
        word:=wordSplittedByEqualSign[i]
        word_index:=strings.Index(entry[offset:], word)
        var new_pair =  []int {(word_index+offset),(len(word)-1+offset)}
        indexes=append(indexes, new_pair)
        offset+=len(word)+1 // 1 is equal sign

        }

        return indexes
    }
    return singlesign_escaper
}


func QuotesSpreading ( entry string) ( word_set [3]string, complete [3]bool  ) {

    var quotes                    =  [3]string {`"`, "'", "`"}

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

func Escape_EqualSign (lineAsArray []string) (words_indexes [][]int) {

    entry:=strings.Join(lineAsArray, "")
    wordSplittedByEqualSign:= strings.Split(entry, "=")
    offset:=0
    for i := range wordSplittedByEqualSign {
        word:=wordSplittedByEqualSign[i]
        word_index:=strings.Index(entry[offset:], word)
        var new_pair =  []int {(word_index+offset),(len(word)-1+offset)}
        words_indexes=append(words_indexes, new_pair)
        offset+=len(word)+1 // 1 is equal sign
    }
    return words_indexes

}

func Escape_Section ( entry string ) ( name, tag []int , section_type int ) {

    entryAsArray:=strings.Split(entry,"")


    square         :=0
    triangle       :=1
    curly          :=2

    opening        :=0
    closing        :=1
    opening_slashed:=2

    // When section has square type
    square_section_opening_index := strings.Index(entry,section_brackets_square[opening])
    square_section_closing_index := strings.Index(entry,section_brackets_square[closing])

    if square_section_opening_index  == 0 && square_section_closing_index  == (len(entry)-1) {

        return []int {1, square_section_closing_index}, tag , square

    }

    // When section has trianle type

    triangle_section_opening_index         := strings.Index(entry, section_brackets_triangle[opening])
    triangle_section_opening_slashed_index := strings.Index(entry, section_brackets_triangle[opening_slashed])
    triangle_section_closing_index         := strings.Index(entry, section_brackets_triangle[closing])

    if (triangle_section_opening_index == 0 || triangle_section_opening_slashed_index == 0) && (strings.Index(entry,section_brackets_triangle[closing]) == (len(entry)-1)) {

            opening_index:=1
            //replacing 1 to 2 because "</"-has two characters but "<"-just one
            if (triangle_section_opening_slashed_index == 0){ opening_index=2 }
            space_indexes:=GetSeparatorIndexes(entry, " ")
            var name_index  = []int {}
            var tag_index   = []int {}
            if len(space_indexes)>0 {
                first_space_index:=space_indexes[0]
                name_index  =[]int {opening_index, first_space_index-1}
                tag_index   =[]int {first_space_index+1, triangle_section_closing_index-1}
            } else {
                name_index =[]int {opening_index, triangle_section_closing_index-1}
                tag_index  =[]int {0,0}
            }
            return name_index, tag_index, triangle

    }
    // When section has curly type
    cleaned_entry_indexes:=RemoveSpaces(entryAsArray,2)
    cleaned_entry:=entry[cleaned_entry_indexes[0]:cleaned_entry_indexes[1]+1]
    cleaned_entry_asArray:=strings.Split(cleaned_entry,"")
    nametag_indexes := RemoveSpaces(cleaned_entry_asArray, 2)
    fmt.Printf("\nnametag_indexes:|%v|\n", nametag_indexes)
    nametag_end_index := strings.Index(entry, cleaned_entry)+(nametag_indexes[1]-nametag_indexes[0])
    fmt.Printf("\nnametag_tag_endindex:|%v|\n", nametag_end_index)

    curly_section_opening_index:=strings.Index(cleaned_entry,section_brackets_curly[opening])
    if curly_section_opening_index == (len(cleaned_entry)-1) {
        if curly_section_opening_index == 0 {
            return []int {0,0} , []int {0,0} , curly
        } else {
             fmt.Printf("\n^^Cleaned Entry: |%s| ^^\n",cleaned_entry)
             spaces:=GetSeparatorIndexes(cleaned_entry, " ")
             var first_space_index   int
             //var second_space_index  int
             var name_index  = []int {}
             var tag_index   = []int {}
             if len(spaces) > 0  { first_space_index  = spaces[0] }
             //if len(spaces) > 1  { second_space_index = spaces[1] }

             cleaned_entry_start_index:=strings.Index(entry, cleaned_entry)
             fmt.Printf("\ncleaned_entry|%v|   cleaned_entry_start_index|%v|  spaces|%v|\n",cleaned_entry, cleaned_entry_start_index, spaces)

             name_index = []int {cleaned_entry_start_index,cleaned_entry_start_index+first_space_index-1}
             tag_index  = []int {cleaned_entry_start_index+first_space_index+1,nametag_end_index}
             return name_index, tag_index, curly
             //if len(spaces)<=2 {
                 //return name_index, tag_index, curly
             //} else {
                 // take first element is CleanedEntry first index
                 // CleanedEntry means entry  without leading and ending spaces 
                 // Example: entry        : "     server   {   "
                 //          cleaned_entry: "server   {"
                 //          spaces       : [6,7,8] 
                 // Another Example:  entry                : "       if a>2 && b==3 {"
                 //                   cleaned_entry        : "if a>2 && b==3 {"
                 //                   spaces               : [2,6,9,14]
                 //                   spaces[len(spaces)-1]:  14   // it is last space in cleaned entry
                 //last_space_inside_cleaned_entry:=spaces[len(spaces)-1]

                 //if curly_section_opening_index-last_space_inside_cleaned_entry== 1 {
                 //    tag_index = []int {cleaned_entry_start_index+first_space_index+1, nametag_end_index}
                 //} else {
                 //tag_index = []int {cleaned_entry_start_index+first_space_index+1, nametag_end_index}
                 //}
                 //return name_index, tag_index, curly

             //}
        }
    }
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

func DebugPrintCharCounter (line string) {

    heads,foots:=DebugCharCounter(line)
    for i:=range heads {
        fmt.Printf("\n%s\n%s\n",heads[i],foots[i])
    }

}

func RemoveSpaces(lineAsArray []string, remove_type int)([]int) {

    LEADING    :=0
    CLOSING    :=1
    BOTH       :=2

    //lineAsArray:=strings.Split(entry, "")
    leadingChar:=0
    closingChar:=len(lineAsArray)-1
    leadReady:=false
    closeReady:=false
    for char := range lineAsArray {
        if (remove_type==LEADING || remove_type==BOTH) && lineAsArray[char] != " " {
                if leadReady != true {

                    leadingChar=char
                    leadReady=true
                    if remove_type==LEADING { break }

                }
        }

        closing_char:=len(lineAsArray)-1-char

        if (remove_type==CLOSING || remove_type==BOTH) && (lineAsArray[closing_char]!=" ")  {
                 if closeReady != true {
                     closingChar=closing_char // +1
                     closeReady=true
                     if remove_type==CLOSING { break }
                 }

        }
        if closeReady && leadReady { break }
    }
    if closingChar<leadingChar { return []int {0,0} }
    return []int {leadingChar,closingChar}
}

func Escape_Sequence(entry string)(sequences [][]int) {

    return sequences

}

func RemoveDupSpaces ( lineAsArray []string ) (new_entry string) {


    newlineAsArray:=ReplaceTabsToSpaces(lineAsArray)

    for i := range newlineAsArray {
        char:=lineAsArray[i]
        if (char==" ") && (len(lineAsArray)-1>i) && (lineAsArray[i+1]==" ") {} else { new_entry+=char }
    }
    return new_entry
}


func Escape_Colon(entry string)(indexes [][]int) {

    colon_indexes:=GetSeparatorIndexes(entry, ":")
    fmt.Printf("\nColon indexes: %v \n",colon_indexes)
    lineAsArray:=strings.Split(entry,"")
    prev_colon_index:=0
    for i := range colon_indexes {
        offset:=1
        if i == 0 && colon_indexes[i]!=0  { offset = 0 }
        if i<=(len(colon_indexes)-1){
            start := prev_colon_index+offset
            end   := colon_indexes[i]-1
            // Warning: When there is nothing between two colons start may be bigger than end . It may cause error when will try print string
            // with intervals. Example: someword[start:end]
            indexes=append(indexes,[]int { start, end})
        }
        prev_colon_index=colon_indexes[i]
        if i == (len(colon_indexes)-1) {
            start := colon_indexes[i]+offset
            end   := (len(lineAsArray)-1)
            // Warning: When there is nothing between two colons start may be bigger than end . It may cause error when will try print string
            // with intervals. Example: someword[start:end]
            indexes=append(indexes,[]int { start, end})
        }
    }
    return indexes
}

func MakeParser(sign string) (function  func(lineAsArray []string)([][]int)) {

    // return function used for parse entry
    fmt.Printf("\nParserMaker: |%s|\n", sign)

    sign_asArray:=strings.Split(sign,"")
    clean_sign_indexes:=RemoveSpaces(sign_asArray,BOTH)
    clean_sign:=sign[clean_sign_indexes[0]:clean_sign_indexes[1]+1]

    sign=clean_sign
    fmt.Printf("\nParserMaker:Checking sign\n")
    if sign == "" { sign = " " } // if delimiter was space we have to restore it

    if IsSymbolIn(sign, SINGLE_SIGNS_LIST ) {
        fmt.Printf(":: MakeParser :: single :: sign %v ", sign)
        return EscapeSingleSign_functionBuilder(sign)
    }
    if IsSymbolIn(sign, DOUBLE_SIGNS_LIST )  {
        fmt.Printf(":: MakeParser :: double :: sign %v ", sign)
        return EscapeDoubleSign_functionBuilder(sign, GetSignPair(sign))
    }

    return function

}
