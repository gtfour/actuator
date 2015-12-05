package main
import "fmt"
import "strings"

func main() {

    test:="hello	hello"
    fmt.Printf("\norig:|%s|\ntab_replaced|%s|\n",test,strings.Replace(test, "	", " ", -1) )

}
