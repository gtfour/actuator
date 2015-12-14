package cuda
import "strings"
import "fmt"

/*
var splitted_by_space  int = 0
var splitted_by_colon  int = 1
*/
/* when it facing with an of the delimiters keep it in mind and try to find next . then group text between 
both delimiters to array  */

var ABC = []string  { "A","a","B","b","C","c","D","d","E","e","F","f","G","g","H","h","I","i","J","j","K","k","L","l","M","m","N","n","O","o","P","p","Q","q","R","r","S","s","T","t","U","u","V","v","W","w","X","x","Y","y","Z","z" }

var NUMBERS                     = []string { "0","1","2","3","4","5","6","7","8","9" }
var WORD_DELIM                  = []string {"_","-"}
var PATH_DELIM                  = []string {"/"}


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

func GetKeyByValue(signs map[int]string, string_value string) (key int) {


    for key, value :=range signs {
        if value == string_value {
            return key
        }

    }
    return -1
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

//var SQ_CU
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



func AcidPriorityAnalyzer (entry string) () {
    //lineSplitBySpace:=
    //lineSplitByQuote:=
    //lineSplitByColon:=

}

func AcidSequencer(entry string) ()  {


}

//func 

/*

var SECTION_SQUARE_OPEN = int  0 







*/

func SortByNested ( entry string ) () {




}

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

func PopArrray( double [][]int) (single []int) {
    if len(double)  ==  1 {
        single=double[0]
    }
    return single
}

func PrepareData ( lineAsArray []string ) ([][]int) {

    //var cleanData = [][]int {}
    var delims    = [][]int {}
    //var words     = [][]int {}
    var delimPair = []int   {-1, -1}

    fmt.Printf("\n[ Array len %d  ]\n",len(lineAsArray))
    for i:= range lineAsArray {
        fmt.Printf("\n==================      for counter -->>  %d ==================\n",i)
        char:=lineAsArray[i]
        if IsSymbolIn(char,ABC,NUMBERS,WORD_DELIM) == false {
            fmt.Printf("\nspecial character: %s >>\n",char)
            //fmt.Printf("\n next symbol is word: %v >>\n", IsSymbolIn(lineAsArray[i+1],ABC,NUMBERS,WORD_DELIM))
            fmt.Printf("\nlen(lineAsArray)-2) : %d\n",(len(lineAsArray)-2))
            if delimPair[0] == -1 {
                delimPair[0]= i
            }
                delimPair[1] = i
                fmt.Printf("\nlen(lineAsArray)-1) : %d\n",(len(lineAsArray)-1))
                if ((i==(len(lineAsArray)-1)) || ((i<=len(lineAsArray)-2) && (IsSymbolIn(lineAsArray[i+1],ABC,NUMBERS,WORD_DELIM) == true))) {
                    fmt.Printf("\n     -- Condition matched \n")
                    delims=append(delims, delimPair)
                    delimPair=[]int{-1, -1}
                }
        } else {


        }
    }
    return delims
}
