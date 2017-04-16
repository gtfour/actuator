package main

import "fmt"
import "jumper/cuda/analyze"

func main(){
    //
    // name          []int
    // tag           []int
    // section_type  int
    //
    name, tag, section_type := analyze.EscapeSection("updates")
    fmt.Printf("Section Parse:   %v %v  %d\n---",name,tag,section_type)
    //
    name, tag, section_type = analyze.EscapeSection("[updates]")
    fmt.Printf("\nSection Parse:   %v %v  %d\n---",name,tag,section_type)
    //
    name, tag, section_type = analyze.EscapeSection("<dev>")
    fmt.Printf("\nSection Parse:   %v %v  %d\n---",name,tag,section_type)
    //
    name, tag, section_type = analyze.EscapeSection("</dev>")
    fmt.Printf("\nSection Parse:   %v %v  %d\n---",name,tag,section_type)
    //
}
