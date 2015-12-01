package main
import "client_side/cuda"
import "fmt"


func main() {

    word_set,complete :=cuda.QuotesParse(`"'Hello'"+""`)
    test:= cuda.GetSeparatorIndexes(`"'Hello'"+""`,`"`)
    word:= "hello my_name waat aha-ha --hi"
    fmt.Printf(":: word \n%s\n::",word)
    cuda.GetWordIndexes(word)
    fmt.Printf("\n%v\n%v\n",word_set,complete)
    fmt.Printf(" -- %v -- ",test)


}
