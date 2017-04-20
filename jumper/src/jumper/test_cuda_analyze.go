package main

import "fmt"
import "jumper/cuda/analyze"
import "jumper/cuda/handling"

func main(){
    //
    // name          []int
    // tag           []int
    // section_type  int
    //
    //
    fmt.Printf("\n--- --- ---\n")
    myString1:="updates"
    name, tag, section_type := analyze.EscapeSection(myString1)
    fmt.Printf("Section Parse:   %v %v  %d\n---",name,tag,section_type)
    analyze.DebugPrintCharCounter(myString1)
    breaker1 := handling.GetSectionBreaker( myString1, name, tag , section_type  )
    fmt.Printf("\nget breaker for this string: %v\n",breaker1("bull-shit"))
    fmt.Printf("\n--- --- ---\n")
    //
    //
    myString2:="[updates]"
    name, tag, section_type = analyze.EscapeSection(myString2)
    fmt.Printf("\nSection Parse:   %v %v  %d\n---",name,tag,section_type)
    analyze.DebugPrintCharCounter(myString2)
    breaker2 := handling.GetSectionBreaker( myString2, name, tag, section_type )
    fmt.Printf("\nget breaker for this string: passing empty line : must be true: %v\n",breaker2(""))
    fmt.Printf("\n--- --- ---\n")
    //
    //
    myString3:=`<div class="print-logo-wrapper" type="button">`
    name, tag, section_type = analyze.EscapeSection(myString3)
    fmt.Printf("\n Section Parse:   %v %v  %d\n---",name,tag,section_type)
    fmt.Printf("\n Section Name : %v \n",myString3[name[0]:name[1]])
    analyze.DebugPrintCharCounter(myString3)
    breaker3 := handling.GetSectionBreaker( myString3, name, tag, section_type )
    fmt.Printf("\nget breaker for this string: passing close-tag </div> : must be true: %v\n",breaker3("</div>"))
    fmt.Printf("\n--- --- ---\n")
    //
    //
    myString4:="</div>"
    name, tag, section_type = analyze.EscapeSection(myString4)
    fmt.Printf("\nSection Parse:   %v %v  %d\n---",name,tag,section_type)
    analyze.DebugPrintCharCounter(myString4)
    breaker4 := handling.GetSectionBreaker( myString4, name, tag, section_type )
    fmt.Printf("\nget breaker for this string: %v\n",breaker4(myString4))
    fmt.Printf("\n--- --- ---\n")
    //
    //
}
