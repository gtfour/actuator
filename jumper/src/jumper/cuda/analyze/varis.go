package analyze

var delimiters                =  []string {":", "="}
var sequence                  =  []string {",",";"}
var word_delimiters           =  []string {"-","_"}
var brackets                  =  []string {"[","]","<","/>",">","{","}",")","("}
// var section_brackets_square   =  [2]string {"[","]"}
// var section_brackets_triangle =  [3]string {"<",">","</"}
// var section_brackets_curly    =  [2]string {"{","}"}
// TODO: add ident section type . Example ifconfig and dmidecode

var LEADING int                  = 0
var CLOSING int                  = 1
var BOTH    int                  = 2
