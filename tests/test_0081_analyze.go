package main

import "fmt"
import "strings"
import "jumper/cuda/analyze"

func main() {

    myTestString              := "      move or cause to move so as to cover an opening."
    myTestStringArray         := strings.Split(myTestString,"")
    selector                  := []int { 3,10 }
    //
    myTestStringArraySelected := analyze.GetFixedArrayChars(myTestStringArray,  selector)
    fmt.Printf("string: %v\nselector:%v\nstring selected:%v\nsize of selected string:%v\n",myTestString,selector,myTestStringArraySelected, len(myTestStringArraySelected))
    //
    //
    sectionExample         :=  "[ base vi1   ]"
    sectionExampleAsArray  :=  strings.Split(sectionExample, "")
    sectionNameIndexes     :=  []int{ 0+1, 13-1 }
    fmt.Printf("\n<sectionNameIndexes: %v>\n", sectionNameIndexes)
    sectionNameArray                := analyze.GetFixedArrayChars( sectionExampleAsArray, sectionNameIndexes )
    fmt.Printf("\n<sectionNameArray:|%v|>\n", sectionNameArray)
    sectionNameWithoutSpacesIndexes := analyze.RemoveSpaces(sectionNameArray, 2)
    fmt.Printf("\n<sectionNameWithoutSpacesIndexes: %v>\n", sectionNameWithoutSpacesIndexes)

    //
    //

}
