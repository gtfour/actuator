package custom

import "client/cuda"


var CUSTOM_PARSERS = CreateParserList()



func CreateParserList ()(cuda.ParserList) {

    pl := make(cuda.ParserList,0)
    pl.Append(shitty_parser)
    return pl

}


// You can add here custom parser's

