package main
import "client_side/cuda"
import "fmt"


func main() {

    word_set,complete :=cuda.QuotesParse(`"'Hello'"+""`)
    test:= cuda.SeparatorIndexes(`"'Hello'"+""`,`"`)
    fmt.Printf("\n%v\n%v\n",word_set,complete)
    fmt.Printf(" -- %v -- ",test)


}
