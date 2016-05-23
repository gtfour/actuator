package custom

var shitty_parser = Parser{Name:"shitty",
                           Call:func( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
                                    return delims,data
                           }
                    }
