package cuda
import "fmt"

var left_direction   int = 1100
var right_direction  int = 1001
var both_directions  int = 2002

type Symbol struct {
    Value           string
    SearchDirection int
    MaxCount        int
    Accepter        func(string)(bool)
    Breaker         func(string)(bool)
}

var URL_SPEC_CHARS = []string {"%","=",":","/","@","?","#"}

type Cyclone struct {
    // line prop
}

func DataHeaderSelector(first_table_string []string)(data [][]int, isTableHeader bool ) {
    return
}

/*func UrlSelector(str []string, delim []int,  data_before []int , data_after []int)(data [][]int, isUrl bool ) {
    fmt.Printf("Delim:%v StrPart:%v", delim,str[delim[0]:delim[1]])
    return
}

func UrlMatcher(str []string, delim []int ) {

    match:=str[delim[0]:delim[1]]
    fmt.Printf("match:%v  str:%v  delim:%v",match,str,delim )


}*/

func StringArrayIsEqual (abc , def []string) (bool) {

    return true


}

func UrlFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {

    fmt.Printf("\n line: %v \n  delims: %v \n  data: %v \n",lineAsArray,delims,data)
    return ndelims,ndata

}

func ArrayInArrayIndexes (abc []string, phrases ...[]string )(indexes [][]int) {
    // In case when phrases have same prefix result is filling by longest phrase 
    // freaks attack

    if (len(abc) < 1 )||(len(phrases) < 1){return}
    //  first_match := -1
    //  last_match  := -1
    for i := 0; i < len(abc); i++  {
        symbol:=abc[i]
        //var found [][]int
        for p := range phrases {
            local_found:=make( []int, 2 )
            phrase:=phrases[p]
            if len(phrase) > 1 {
                //for z:= range phrase {
                    zsymbol := phrase[0]
                    if symbol == zsymbol {
                        //xi:=i
                        local_found[0] =  i
                        counter        := 1
                        xi             := i+1
                        for  ; counter < len(phrase) ;  {
                            //xi         :=  i
                            xsymbol    := abc[xi]
                            if xsymbol != phrase[counter] { break /*; z=z+counter-1*/ }
                            if counter >= len(phrase)-1 { local_found[1] = xi ; indexes = append(indexes, local_found) ; break }
                            xi         += 1
                            counter    += 1
                        }
                        //break
                    } else {
                        //if z == (len(phrase)-1) {

                        //}
                    }
                //}
            }
        }
    }
    return
}

//func AnalyzeDelims()
