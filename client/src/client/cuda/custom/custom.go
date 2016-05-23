package custom

import "errors"

var dup_name     = errors.New("error:parser with following name is already exist")
var name_is_none = errors.New("error:parser name wasn't specified")

var CUSTOM_PARSERS = CreateParserList()

type ParserList []Parser


func CreateParserList ()(ParserList) {

    pl := make(ParserList,0)
    pl.Append(shitty_parser)
    return pl

}

type Parser struct {

    Name string
    Call func( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int)

}

func (pl ParserList) Append (new_parser Parser)(error) {

    if (new_parser.Name == "") { return name_is_none }
    for key := range pl {
        parser := pl[key]
        if ( new_parser.Name == parser.Name ) {
            return dup_name
        }
    }
    pl = append(pl, new_parser)
    return nil
}







// You can add here custom parser's

