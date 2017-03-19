package analyze

import "fmt"
import "strings"

var section_brackets_square   =  [2]string {"[","]"}
var section_brackets_triangle =  [3]string {"<",">","</"}
var section_brackets_curly    =  [2]string {"{","}"}

func Escape_Section ( entry string ) ( name, tag []int , section_type int ) {

    entryAsArray := strings.Split(entry,"")

    square         := 0
    triangle       := 1
    curly          := 2

    opening        := 0
    closing        := 1
    opening_slashed:= 2

    // When section has square type
    square_section_opening_index := strings.Index(entry,section_brackets_square[opening])
    square_section_closing_index := strings.Index(entry,section_brackets_square[closing])

    if square_section_opening_index  == 0 && square_section_closing_index  == (len(entry)-1) {

        return []int {1, square_section_closing_index}, tag , square

    }

    // When section has trianle type

    triangle_section_opening_index         := strings.Index(entry, section_brackets_triangle[opening])
    triangle_section_opening_slashed_index := strings.Index(entry, section_brackets_triangle[opening_slashed])
    triangle_section_closing_index         := strings.Index(entry, section_brackets_triangle[closing])

    if (triangle_section_opening_index == 0 || triangle_section_opening_slashed_index == 0) && (strings.Index(entry,section_brackets_triangle[closing]) == (len(entry)-1)) {

            opening_index:=1
            //replacing 1 to 2 because "</"-has two characters but "<"-just one
            if (triangle_section_opening_slashed_index == 0){ opening_index=2 }
            space_indexes:=GetSeparatorIndexes(entry, " ")
            var name_index  = []int {}
            var tag_index   = []int {}
            if len(space_indexes)>0 {
                first_space_index:=space_indexes[0]
                name_index  =[]int {opening_index, first_space_index-1}
                tag_index   =[]int {first_space_index+1, triangle_section_closing_index-1}
            } else {
                name_index =[]int {opening_index, triangle_section_closing_index-1}
                tag_index  =[]int {0,0}
            }
            return name_index, tag_index, triangle

    }
    // When section has curly type
    cleaned_entry_indexes:=RemoveSpaces(entryAsArray,2)
    cleaned_entry:=entry[cleaned_entry_indexes[0]:cleaned_entry_indexes[1]+1]
    cleaned_entry_asArray:=strings.Split(cleaned_entry,"")
    nametag_indexes := RemoveSpaces(cleaned_entry_asArray, 2)
    fmt.Printf("\nnametag_indexes:|%v|\n", nametag_indexes)
    nametag_end_index := strings.Index(entry, cleaned_entry)+(nametag_indexes[1]-nametag_indexes[0])
    fmt.Printf("\nnametag_tag_endindex:|%v|\n", nametag_end_index)

    curly_section_opening_index:=strings.Index(cleaned_entry,section_brackets_curly[opening])
    if curly_section_opening_index == (len(cleaned_entry)-1) {
        if curly_section_opening_index == 0 {
            return []int {0,0} , []int {0,0} , curly
        } else {
             fmt.Printf("\n^^Cleaned Entry: |%s| ^^\n",cleaned_entry)
             spaces:=GetSeparatorIndexes(cleaned_entry, " ")
             var first_space_index   int
             //var second_space_index  int
             var name_index  = []int {}
             var tag_index   = []int {}
             if len(spaces) > 0  { first_space_index  = spaces[0] }
             //if len(spaces) > 1  { second_space_index = spaces[1] }

             cleaned_entry_start_index:=strings.Index(entry, cleaned_entry)
             fmt.Printf("\ncleaned_entry|%v|   cleaned_entry_start_index|%v|  spaces|%v|\n",cleaned_entry, cleaned_entry_start_index, spaces)

             name_index = []int {cleaned_entry_start_index,cleaned_entry_start_index+first_space_index-1}
             tag_index  = []int {cleaned_entry_start_index+first_space_index+1,nametag_end_index}
             return name_index, tag_index, curly
             //if len(spaces)<=2 {
                 //return name_index, tag_index, curly
             //} else {
                 // take first element is CleanedEntry first index
                 // CleanedEntry means entry  without leading and ending spaces 
                 // Example: entry        : "     server   {   "
                 //          cleaned_entry: "server   {"
                 //          spaces       : [6,7,8] 
                 // Another Example:  entry                : "       if a>2 && b==3 {"
                 //                   cleaned_entry        : "if a>2 && b==3 {"
                 //                   spaces               : [2,6,9,14]
                 //                   spaces[len(spaces)-1]:  14   // it is last space in cleaned entry
                 //last_space_inside_cleaned_entry:=spaces[len(spaces)-1]

                 //if curly_section_opening_index-last_space_inside_cleaned_entry== 1 {
                 //    tag_index = []int {cleaned_entry_start_index+first_space_index+1, nametag_end_index}
                 //} else {
                 //tag_index = []int {cleaned_entry_start_index+first_space_index+1, nametag_end_index}
                 //}
                 //return name_index, tag_index, curly

             //}
        }
    }
    return

}
