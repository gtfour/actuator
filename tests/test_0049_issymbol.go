package main
import . "unicode"
import "fmt"

func main() {

    /*mystring     := "a"
    myunderscore := "_"
    mydigit      := "1"
    var mybyte1,mybyte2,mybyte3 byte
    if len(mystring) == 1 {
        mybyte1=mystring[0]
    }
    if len(myunderscore) == 1 {
        mybyte2=myunderscore[0]
    }
    if len(mydigit) == 1 {
        mybyte3=mydigit[0]
    }

    fmt.Printf("%v   %v   %v\n",IsLetter(rune(mybyte1)),IsLetter(rune(mybyte2)),IsLetter(rune(mybyte3)))
    */
    for _,r := range "1"  { // knows about russian letters
        fmt.Printf("IsLetter:%v IsDigit:%v IsNumber:%v\n",IsLetter(r),IsDigit(r),IsNumber(r))
    }
}
