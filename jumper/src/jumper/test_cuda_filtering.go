package main

import "fmt"
import "strings"
import "jumper/cuda/analyze"
import "jumper/cuda/filtering"

func main() {

    myString     := "puppet:x:999:998:puppetserver daemon:/opt/puppetlabs/server/data/puppetserver:/usr/sbin/nologin"
    lineSlice    := strings.Split(myString,"")
    analyze.DebugPrintCharCounter(myString)
    delims,datas := analyze.GetIndexes(lineSlice)
    ndelims,ndatas :=filtering.ColonFilter( lineSlice, delims  , datas  )
    fmt.Printf("\n:: Using Colon Filter: delims: %v  datas: %v ::\n", ndelims,ndatas)

}
