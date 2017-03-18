package analyze

import "unicode"
import "strings"

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

        //fmt.Printf("--\n%i\n--",i)
        //
        _ = i

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

func GetSeparatorIndexes (entry, sep string) (indexes []int) {

    chars := strings.Split(entry,"")
    for char := range chars {
        if chars[char] == sep {
            indexes=append(indexes, char)
        }
    }
    return indexes
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
