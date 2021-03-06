package analyze

import "strings"

var section_brackets_square   =  [2]string {"[","]"}
var section_brackets_triangle =  [3]string {"<",">","</"}
var section_brackets_curly    =  [2]string {"{","}"}


func EscapeSection( entry string ) ( name, tag [2]int , section_type int ) {

    // match tabbed section by keyword

    entryAsArray         :=  strings.Split(entry,"")
    section_type         =   NOT_SECTION
    if len(entryAsArray) == 0 { return }

    //                   
    // // square         := 0
    // // triangle       := 1
    // // curly          := 2
    //

    opening         := 0
    closing         := 1
    opening_slashed := 2

    //
    // check if section has square type
    //
    square_section_opening_index := strings.Index( entry, section_brackets_square[opening] )
    square_section_closing_index := strings.Index( entry, section_brackets_square[closing] )
    if ( square_section_opening_index>=0 && square_section_closing_index >0 &&  square_section_opening_index < square_section_closing_index ) {
        //
        // test
        //
        // cleanEntryIndexes            := RemoveSpaces(entryAsArray, 2)
        // GetFixedArrayChars(lineAsArray []string, selected_indexes[]int) (selected []string)
        //
        sectionNameIndexes              := []int{ square_section_opening_index+1, square_section_closing_index-1 }
        //fmt.Printf("\n<sectionNameIndexes: %v>\n", sectionNameIndexes)
        sectionNameArray                := GetFixedArrayChars( entryAsArray, sectionNameIndexes )
        //fmt.Printf("\n<sectionNameArray: %v>\n", sectionNameArray)
        sectionNameWithoutSpacesIndexes := RemoveSpaces(sectionNameArray, 2)
        //
        // test 
        //
        offsetOpening     := square_section_opening_index + 1
        offsetClosing     := offsetOpening
        squareSectionName := [2]int{ sectionNameWithoutSpacesIndexes[0]+offsetOpening, sectionNameWithoutSpacesIndexes[1]+offsetClosing }
        //
        //
        //
        //fmt.Printf("\n<sectionNameWithoutSpacesIndexes: %v>\n", sectionNameWithoutSpacesIndexes)
        // now we have to remove leading and closing spaces
        //
        //
        if square_section_opening_index  == 0 && square_section_closing_index  == (len(entry)-1) {
            //
            return squareSectionName, tag , SQUARE_SECTION
            //
        }
        //
        //
        //
    }
    //
    // check if section has trianle type
    //
    triangle_section_opening_index         := strings.Index(entry, section_brackets_triangle[opening])
    triangle_section_opening_slashed_index := strings.Index(entry, section_brackets_triangle[opening_slashed])
    triangle_section_closing_index         := strings.Index(entry, section_brackets_triangle[closing])

    if (triangle_section_opening_index == 0 || triangle_section_opening_slashed_index == 0) && (strings.Index(entry,section_brackets_triangle[closing]) == (len(entry)-1)) {

            //
            kindOfTriangleSection  :=  TRIANGLE_SECTION_UNDEFINED
            if triangle_section_opening_index == 0         { kindOfTriangleSection = TRIANGLE_SECTION_STARTING }
            if triangle_section_opening_slashed_index == 0 { kindOfTriangleSection = TRIANGLE_SECTION_ENDING  }
            //
            opening_index:=1
            //
            // replacing 1 to 2 because "</"-has two characters but "<"-just one
            //
            if (triangle_section_opening_slashed_index == 0){
                opening_index = 2
            }
            space_indexes := GetSeparatorIndexes(entry, " ")
            //
            // it seems that here we try to get tag and name indexes
            // i suppose that section like that "</blablabla>" could'nt has any tags inside
            //
            var name_index  = [2]int {}
            var tag_index   = [2]int {}
            if len(space_indexes)>0 {
                first_space_index:=space_indexes[0]
                name_index  =[2]int {opening_index, first_space_index-1}
                tag_index   =[2]int {first_space_index+1, triangle_section_closing_index-1}
            } else {
                name_index =[2]int {opening_index, triangle_section_closing_index-1}
                tag_index  =[2]int {0,0}
            }
            return name_index, tag_index, kindOfTriangleSection

    }
    //
    // check if section has curly type
    //
    cleaned_entry_indexes  :=  RemoveSpaces( entryAsArray, 2 )
    // fmt.Printf("\n---\ncleaned_entry_indexes: %v\n---\nentry: %v\n---", cleaned_entry_indexes, entry )
    cleaned_entry          :=  entry[cleaned_entry_indexes[0]:cleaned_entry_indexes[1]+1]
    cleaned_entry_asArray  :=  strings.Split( cleaned_entry,"" )
    //
    //
    //
    nametag_indexes := RemoveSpaces(cleaned_entry_asArray, 2)
    // fmt.Printf("\nnametag_indexes:|%v|\n", nametag_indexes)
    nametag_end_index := strings.Index(entry, cleaned_entry)+(nametag_indexes[1]-nametag_indexes[0])
    // fmt.Printf("\nnametag_tag_endindex:|%v|\n", nametag_end_index)

    curly_section_opening_index:=strings.Index(cleaned_entry,section_brackets_curly[opening])
    if curly_section_opening_index == (len(cleaned_entry)-1) {
        if curly_section_opening_index == 0 {
            return [2]int {0,0} , [2]int {0,0} , CURLY_SECTION
        } else {
             // fmt.Printf("\n^^Cleaned Entry: |%s| ^^\n",cleaned_entry)
             spaces:=GetSeparatorIndexes(cleaned_entry, " ")
             var first_space_index   int
             //var second_space_index  int
             var name_index  = [2]int {}
             var tag_index   = [2]int {}
             if len(spaces) > 0  { first_space_index  = spaces[0] }
             //if len(spaces) > 1  { second_space_index = spaces[1] }

             cleaned_entry_start_index:=strings.Index(entry, cleaned_entry)
             // fmt.Printf("\ncleaned_entry|%v|   cleaned_entry_start_index|%v|  spaces|%v|\n",cleaned_entry, cleaned_entry_start_index, spaces)

             name_index = [2]int {cleaned_entry_start_index,cleaned_entry_start_index+first_space_index-1}
             tag_index  = [2]int {cleaned_entry_start_index+first_space_index+1,nametag_end_index}
             return name_index, tag_index, CURLY_SECTION
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
    //
    // check if section has tabbed (check if keyword exists inside entry)
    //
    //if len(keyword) > 0 {
    //    first_keyword := keyword[0]
    //    keyword_index := strings.Index(entry, first_keyword)
    //}
    //
    //
    //
    return

}

func EscapeIndentSection(entry string, keywords []string)(name, tag [2]int , section_type int){
    //
    //
    section_type    =  NOT_SECTION
    keywords_len    := len(keywords)
    if keywords_len <= 0 { return }
    if len(entry)   <= 0 { return }
    // last            := keywords_len-1
    found           := false
    for i := range keywords {
        keyword       := keywords[i]
        keyword_index := strings.Index(entry, keyword)
        if keyword_index < 0 {
            continue
        } else {
            found = true
            break
        }
    }
    if !found { return }
    last_char    := len(entry)-1
    name         = [2]int{0, last_char}
    section_type = INDENT_SECTION
    return
    //
    //
}

func SectionCouldBeNested(section_type int)(yes bool){
    // incomplete ! have to add checking curly sections 
    if section_type == TRIANGLE_SECTION_STARTING || section_type == CURLY_SECTION { yes = true }
    return
    //
}

func SectionCouldBeInline(section_type int)(yes bool) {
    if section_type == TRIANGLE_SECTION_STARTING || section_type == CURLY_SECTION { yes = true }
    return
}
