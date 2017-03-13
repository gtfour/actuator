package analyze

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
