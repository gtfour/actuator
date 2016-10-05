package parse
//
import "unicode"
//
// same functions are existing inside package "client/cuda"
//
var WORD_DELIM                  = []string {"_","-"}

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
