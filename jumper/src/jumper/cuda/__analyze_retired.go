package cuda

import "fmt"
import "strings"
import "unicode"

/*
var splitted_by_space  int = 0
var splitted_by_colon  int = 1
*/
/* when it facing with an of the delimiters keep it in mind and try to find next . then group text between 
both delimiters to array  */

//
//
// has been moved into analyzing package
//
//

/*
var ABC = []string  { "A","a","B","b","C","c","D","d","E","e","F","f","G","g","H","h","I","i","J","j","K","k","L","l","M","m","N","n","O","o","P","p","Q","q","R","r","S","s","T","t","U","u","V","v","W","w","X","x","Y","y","Z","z" }

var NUMBERS                     = []string { "0","1","2","3","4","5","6","7","8","9" }
var WORD_DELIM                  = []string {"_","-"}
var PATH_DELIM                  = []string {"/"}
var URL                         = []string {"://"}

var DOUBLE_SIGNS_LIST = []string { "[", "]", "<" , ">", "</" ,  "(" , ")", "{", "}", "'", "'" , `"`,`"`, "`","`" }
var SINGLE_SIGNS_LIST = []string { "="," ",":" }


var OPEN_SECTION_SQUARE      int = 0   //   [
var CLOSE_SECTION_SQUARE     int = 1   //   ]
var OPEN_SECTION_TRIANGLE    int = 2   //   <
var OPEN_SL_SECTION_TRIANGLE int = 3   //   </
var CLOSE_SECTION_TRIANGLE   int = 4   //   >
var OPEN_SECTION_ROUND       int = 5   //   (
var CLOSE_SECTION_ROUND      int = 6   //   )
var OPEN_SECTION_CURLY       int = 7   //   {
var CLOSE_SECTION_CURLY      int = 8   //   {

var SINGLE_QUOTE             int = 8   //   '
var DOUBLE_QUOTE             int = 9   //   "
var GRAVE_QUOTE              int = 10   //   `

var HYPHEN                   int = 11  // -
var MINUS                    int = 11  // -
var PLUS                     int = 12  // +
var UNDERSCORE               int = 13  // _
var EQUAL                    int = 14  // = 
var COLON                    int = 15  // : 
var SEMICOLON                int = 16  // ;
var COMMA                    int = 17  // ,
var DOT                      int = 18  // .

var SLASH                    int = 19  // /
var BACKSLASH                int = 20  // \
var PIPE                     int = 21  // |
var ASTERISK                 int = 22  // *
var NUMBER                   int = 23  // #
var PENIS                    int = 24  // o|o
var DOLLAR                   int = 25  // $
var AMPERSAND                int = 26  // &
var SUFFIX                   int = 27  // ^
var PERCENT                  int = 28  // %
var MAIL                     int = 29  // @
var EXCLAM                   int = 30  // !
var TILDE                    int = 31  // ~



var RIGHT                            int = 123  // :) ahaha blya
var LEFT                             int = 321  // :)

var NOT_FOUND                        int = 2000 //  

var JUST_WORDS_DIGITS                int = 3001
var JUST_WORDS_DIGITS_HYPHENS        int = 3002
var JUST_WORDS_DIGITS_HYPHENS_PATHES int = 3003

var FINISH_WITH_SPACE                int = 4000
var FINISH_WITH_SAME_SIGN            int = 4001


//var SIGNS = make(map[int]string)
//SIGNS[OPEN_SECTION_SQUARE]="["
//SIGNS := map[int]string{ EXCLAM : "!" }





func SignMap()(signs map[int]string) {
    signs=make(map[int]string)
    signs[OPEN_SECTION_SQUARE]      ="["
    signs[CLOSE_SECTION_SQUARE]     ="]"
    signs[OPEN_SECTION_TRIANGLE]    ="<"
    signs[OPEN_SL_SECTION_TRIANGLE] ="</"
    signs[CLOSE_SECTION_TRIANGLE]   =">"
    signs[OPEN_SECTION_ROUND]       ="("
    signs[CLOSE_SECTION_ROUND]      =")"
    signs[OPEN_SECTION_CURLY]       ="{"
    signs[CLOSE_SECTION_CURLY]      ="}"

    signs[SINGLE_QUOTE]             ="'"
    signs[DOUBLE_QUOTE]             =`"`
    signs[GRAVE_QUOTE]              ="`"

    signs[HYPHEN]                   ="-"
    signs[PLUS]                     ="+"
    signs[UNDERSCORE]               ="_"
    signs[EQUAL]                    ="="
    signs[COLON]                    =":"
    signs[SEMICOLON]                =";"
    signs[COMMA]                    =","
    signs[DOT]                      ="."

    signs[SLASH]                    ="/"
    signs[BACKSLASH]                =`\`
    signs[PIPE]                     ="|"
    signs[ASTERISK]                 ="*"
    signs[NUMBER]                   ="#"
    signs[DOLLAR]                   ="$"
    signs[AMPERSAND]                ="&"
    signs[SUFFIX]                   ="^"
    signs[PERCENT]                  ="%"
    signs[MAIL]                     ="@"
    signs[EXCLAM]                   ="!"
    signs[TILDE]                    ="~"

    return


}
*/
//
//
//


/*
func GetKeyByValue(signs map[int]string, string_value string) (key int) {
    for key, value :=range signs {
        if value == string_value {
            return key
        }
    }
    return -1
}
*/

//
//
// functions below has been moved into analyze
//
//
/*
func IsUnicodeLetter(char string)(yes bool) {
    for _,r := range char  { // knows about russian letters
        yes = unicode.IsLetter(r)
    }
    return yes
}

func IsUnicodeDigit(char string)(yes bool) {
    for _,r := range char  {
        yes = unicode.IsDigit(r)
    }
    return yes
}

func IsDigitIn(digit int, digits_sets ...[]int) (yes bool) {
    for i := range digits_sets {
        set := digits_sets[i]
        for s := range set {
            ndigit:=set[s]
            if ndigit == digit {
                yes = true
                break
            }
        }
        if yes == true { break }
    }
    return
}
func IsSymbolIn(char string, symbols_sets ...[]string) (yes bool) {
    for i := range symbols_sets {
        set := symbols_sets[i]
        for s := range set {
            symbol:=set[s]
            if symbol == char {
                yes = true
                break
            }
        }
        if yes == true { break }
    }
    return
}
*/

//
//
//
//
//


//
// moved into analyze/maps.go
//
/*
func ValueExists(signs map[int]string,value string)(found bool ) {
    values:=GetMapValues(signs)
    for i := range values {
        if values[i]==value {
            found=true
        }
    }
    return found
}

func GetMapValues(signs map[int]string)(values []string ){
    for _, value := range signs {
        values=append(values, value)
    }
    return values
}
*/
//
//
//

//var SQ_CU

//
//
// has been moved into analyze module
//
/*
func GetSignsIndexes(entry string)(map[int][]int) {

   sign_map:=SignMap()
   sign_indexes:=make(map[int][]int)
   lineAsArray:=strings.Split(entry,"")
   for i := range lineAsArray {
       char:=lineAsArray[i]
       charSignKey:=GetKeyByValue(sign_map, char)
       if charSignKey > 0 {
           if _, ok := sign_indexes[charSignKey]; ok==false {
               sign_indexes[charSignKey]= []int {}
           }
           sign_indexes[charSignKey]=append(sign_indexes[charSignKey], i)

       }

    }
    return sign_indexes

}
*/
//
//
//
//


/*
func AcidPriorityAnalyzer (entry string) () {
    //lineSplitBySpace:=
    //lineSplitByQuote:=
    //lineSplitByColon:=
}

func AcidSequencer(entry string) ()  {
}
*/
//func 

/*
var SECTION_SQUARE_OPEN = int  0 
*/
/*
func SortByNested ( entry string ) () {
}
*/

//
//
//
//
/*

func GetSignScope( lineAsArray []string, sign int, sign_pos int) (scope [][2]int) {
    switch {
        case sign==EQUAL:
            var first_part [2]int
            var last_part  [2]int
            first_part[0] = 0
            first_part[1] = sign_pos-1
            last_part[0] = sign_pos+1
            last_part[1] = len(lineAsArray)-1
            scope=append(scope,first_part)
            scope=append(scope,last_part)
            return scope
    }
    return scope
}


func GoTillAnyOfSign( lineAsArray []string, signs []int, since int, direction int ) ( index int, code int ) {


    for i:= range lineAsArray {

        i = since
        if since<0 {i=0}

        fmt.Printf("--\n%i\n--",i)

        if direction==RIGHT {

        } else if direction==LEFT {

        } else {
            return -1, NOT_FOUND
        }
    }
    return index, NOT_FOUND

}

func CheckMatchingRx( entry string ) (code int) {
    return code
}
*/
//
//
//
//
/*
func PopArrray( double [][]int) (single []int) {
    if len(double)  ==  1 {
        single=double[0]
    }
    return single
}
*/
//
//
// GetIndexes has been moved into analyze module
//
//
func GetIndexes ( lineAsArray []string ) (delims [][]int , data [][]int) {
    var delimPair = []int   {-1,-1}
    var dataPair  = []int   {-1,-1}
    //var offset int
    for i:= range lineAsArray {
        //offset = i
        char:=lineAsArray[i]
        //if IsSymbolIn(char,ABC,NUMBERS,WORD_DELIM) == false {
        if IsSymbolIn(char,WORD_DELIM) == false && IsUnicodeLetter(char) == false  && IsUnicodeDigit(char) == false {
            if dataPair[0]  != -1 {
                dataPair[1] = i - 1 // make pair with previous element as second member of pair 
                data      = append(data, dataPair)
                dataPair  = []int   {-1,-1}
            }
            if delimPair[0] == -1 {
                delimPair[0]= i
            }
            delimPair[1] = i
            //if ((i==(len(lineAsArray)-1)) || ((i<=len(lineAsArray)-2) && (IsSymbolIn(lineAsArray[i+1],ABC,NUMBERS,WORD_DELIM) == true)) ) {
            if (i==(len(lineAsArray)-1)) || ((i<=len(lineAsArray)-2) && ( IsSymbolIn(lineAsArray[i+1],WORD_DELIM) == true || IsUnicodeLetter(lineAsArray[i+1]) == true  || IsUnicodeDigit(lineAsArray[i+1]) == true)) {

                    // +1 because see /actuator/tests/test_0038_arr.go
                    // delimAsArray:=GetFixedArrayChars(lineAsArray[delimPair[0]:offset+1], []int { 0, (delimPair[1]-delimPair[0]) }) // have to add +1 .but   why !??!? 
                    //delim_split_space:=Escape_Spaces(delimAsArray)
                    delims=append(delims, delimPair)
                    delimPair=[]int{-1, -1}
                    // simplifyig
                    /*
                    if len(delim_split_space) == 1 {
                        delims=append(delims, delimPair)
                        delimPair=[]int{-1, -1}
                    } else if len(delim_split_space) == 0   {
                       // it  seems that there are just a lot of spaces in delimAsArray  and nothing else  
                        // delimPair[1] = delimPair[0] i am going change it to collect all space indexes instead of first 
                        delims=append(delims, delimPair)
                        delimPair=[]int{-1, -1}
                    } else {
                        for sd := range delim_split_space {
                           delim_ss_pair:=delim_split_space[sd]
                           if len(delim_ss_pair) == 2 { delim_ss_pair[0]=delim_ss_pair[0]+delimPair[0] ; delim_ss_pair[1]=delim_ss_pair[1]+delimPair[0]  }
                           delims = append(delims, delim_ss_pair)
                        }
                        delimPair=[]int{-1, -1}
                    }*/
            }
        } else {
            if dataPair[0]  == -1 {
                dataPair[0] = i
            }
            if (i==(len(lineAsArray)-1)) {
                dataPair[1] = i
                data      = append(data, dataPair)
            }
        }
    }
    return delims,data
}
//
//
// Functions below has been moved into analyze module
//
//
/*
func GetFixedArrayChars(lineAsArray []string, selected_indexes[]int) (selected []string) {
    for i := range  lineAsArray {
        char:= lineAsArray[i]
        if len(selected_indexes) == 2 {
            if i>=selected_indexes[0] && i<=selected_indexes[1] {
                selected = append(selected, char)
            }
        } else { break }
    }
    return selected
}

func GetSignPair( sign string )( another_sign string) {
    var DOUBLE_SIGNS_PAIRS = [][2]string { {"[", "]"}, { "<" , ">"}, {"</" , ">"},  {"(" , ")"}, {"{", "}"}, {"'", "'" }, {`"`,`"`}, {"`","`"} }
    var REPLACE01    = [2]int {1,0}
    for pairs := range DOUBLE_SIGNS_PAIRS {
        pair:=DOUBLE_SIGNS_PAIRS[pairs]
        for elem := range pair {
            if sign == pair[elem] {
                second_index:=REPLACE01[elem]
                another_sign = pair[second_index]
                break
            }
        }
    }
    return another_sign
}

func PrepareData(lineAsArray []string, delims_indexes [][]int ) (data [][][]int) {
    for d := range  delims_indexes {
        delim_pair:=delims_indexes[d]
        delimAsArray:=GetFixedArrayChars(lineAsArray, delim_pair)
        delim:=strings.Join(delimAsArray,"")
        parser:=MakeParser(delim)
        subdata:=parser(lineAsArray)
        data=append(data,subdata)
    }
    return data
}
*/
//
//
//
//
//
