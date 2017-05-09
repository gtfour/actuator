package main

import "fmt"
import "strings"
import "jumper/cuda/analyze"
import "jumper/cuda/filtering"

func main() {

    myString     := "=puppet==x"
    lineSlice    := strings.Split(myString,"")
    analyze.DebugPrintCharCounter(myString)
    delims,datas := analyze.GetIndexes(lineSlice)
    // // ndelims,ndatas := filtering.ColonFilter( lineSlice, delims  , datas  )
    // // fmt.Printf("\n:: Using Colon Filter: delims: %v  datas: %v ::\n", ndelims,ndatas)
    // // ndelims,ndatas = filtering.UrlFilter( lineSlice, ndelims  , ndatas  )
    // // fmt.Printf("\n:: Using Url Filter: delims: %v  datas: %v ::\n", ndelims,ndatas)
    //
    // EqualSignFilter
    //
    ndelims,ndatas := filtering.EqualSignFilter( lineSlice, delims  , datas  )
    fmt.Printf("\n:: Using Equal Sign Filter: delims: %v  datas: %v ::\n", ndelims,ndatas)

}
