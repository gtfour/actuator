package main

import "fmt"
import "strings"
import "jumper/cuda/filtering"
import "jumper/cuda/analyze"

func main(){
    //
    //

    //
    first_line     := "deb http://ru.archive.ubuntu.com/ubuntu/ xenial-backports main restricted universe multiverse"
    second_line    := "deb [arch=amd64] https://download.docker.com/linux/ubuntu xenial stable"
    third_line     := "gpgcheck=1"
    //
    firstSlice     := strings.Split(first_line,  "")
    secondSlice    := strings.Split(second_line, "")
    thirdSlice     := strings.Split(third_line,  "")
    //
    delims1,datas1 := analyze.GetIndexes( firstSlice  )
    delims2,datas2 := analyze.GetIndexes( secondSlice )
    delims3,datas3 := analyze.GetIndexes( thirdSlice  )
    //
    //
    dataKey1,dataValue1 := filtering.EqualSignRocker( firstSlice,  delims1, datas1 )
    dataKey2,dataValue2 := filtering.EqualSignRocker( secondSlice, delims2, datas2 )
    dataKey3,dataValue3 := filtering.EqualSignRocker( thirdSlice,  delims3, datas3 )
    //
    //
    fmt.Printf("\ndataKey1: %v dataValue1: %v\n",dataKey1,dataValue1)
    fmt.Printf("\ndataKey2: %v dataValue2: %v\n",dataKey2,dataValue2)
    fmt.Printf("\ndataKey3: %v dataValue3: %v\n",dataKey3,dataValue3)
    //

    //
    //
}
