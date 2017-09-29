package main

import "fmt"
import "jumper/cuda/analyze"
import "jumper/cuda/handling"

func main() {
    //
    //
    myString2               := "09/28/17 04:00:20 my_script.py ERROR   : Unhandled exception in _my_func: `Traceback (most recent call last):"
    keyWords                :=  []string{"Traceback (most recent call last)"}
    name, tag, section_type :=  analyze.EscapeIndentSection(myString2, keyWords)
    fmt.Printf("\nSection Parse:   %v %v  %d\n---",name,tag,section_type)
    analyze.DebugPrintCharCounter(myString2)
    breaker2                := handling.GetSectionBreaker( myString2, name, tag, section_type )
    fmt.Printf("\nstring: %v\n",myString2)
    fmt.Printf("checking breaker: %v\n",breaker2(`    import "fvfv"`)           )
    fmt.Printf("checking breaker: %v\n",breaker2(`                ^`)           )
    fmt.Printf("checking breaker: %v\n",breaker2(`SyntaxError: invalid syntax`) )
    fmt.Printf("\n--- --- ---\n")
    //
    //
}
