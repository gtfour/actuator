package b

import "fmt"
import "import_tests2/mytypes"

func Compare(h int)() {

    if h == mytypes.HELLO_1 {
        fmt.Printf("\nHello 1\n")
    } else if  h == mytypes.HELLO_2 {
        fmt.Printf("\nHello 2\n")
    } else {
        fmt.Printf("\nNone\n")
    }
}
