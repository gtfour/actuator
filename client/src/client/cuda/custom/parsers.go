package custom
import "client/cuda"

var shitty_parser = cuda.Parser{Name:"shitty",
                           Call:func( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
                               return delims,data
                           },
                    }
