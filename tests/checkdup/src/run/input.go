package run

import "flag"

var PROPS = GetProps()

func GetProps()(props map[string]interface{}){

    in_file_ptr          := flag.String("infile","none","Input file")
    in_file_offset_ptr   := flag.Int64("offset",0,"Input file offset")
    entry_ptr            := flag.String("entry","hello","New entry")
    flag.Parse()
    //out_file_ptr       := flag.String("outfile","out.txt","Out file")
    in_file              := *in_file_ptr
    offset               := *in_file_offset_ptr
    entry                := *entry_ptr

}
