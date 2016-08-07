package compas

// do not forget to match line number
// uncertain module

// abcData and xyzData is data parsed by cuda

import "fmt"

type changes []Change

type CmpDataProp struct {


}

type CmpDataEntryProp struct {

    oldLineNumber  int
    newLineNumber  int
    statuses       []int
    value          []string

}

func Trite(abcData [][]string, xyzData [][]string) {
    for i := range abcData {
        abcLen:=len(abcData)
        xyzLen:=len(xyzData)
        if (i<=abcLen-1)&&(i<=xyzLen-1){
            abcLine:=abcData[i]
            xyzLine:=xyzData[i]
            status:=TriteLine(abcLine, xyzLine)
            fmt.Printf("\nStatus:%d",status)
        }
    }
}

func TriteLine(abcDataLine []string, xyzDataLine []string)(status int) {
    //sleepStreamIndex:=0
    aLineLen:=len(abcDataLine)
    bLineLen:=len(xyzDataLine)
    fmt.Printf("\nabcLen:%d  xyzLen:%d\n",aLineLen,bLineLen)
    if aLineLen == bLineLen {
       //for i:=0 ; i < len(abcDataLine) ; i++  {
       //     abcLinePart := abcDataLine[i]
       //     xyzDataLine := xyzDataLine[i]
       //    if abcLinePart == xyzDataLine {
       //        sleepStreamIndex = i+1
       //    }
       // }
       ssIndex:=GetSleepStreamIndex(abcDataLine, xyzDataLine)
       if (ssIndex == len(abcDataLine)) {
           fmt.Printf("\nLine is equal\n")
           return EQUAL_ENTRY
       }
    } else {
        ssIndex:=GetSleepStreamIndex(abcDataLine, xyzDataLine)
        if(bLineLen > aLineLen){
            newPartLen:=bLineLen-ssIndex
            newPart:=xyzDataLine[newPartLen]
            fmt.Printf("New part: %s",string(newPart))
            return APPENDED_TO_THE_END_OF_LINE
        }
        //fmt.Printf("S")
    }
    return status
}


func GetSleepStreamIndex(aData []string, bData []string)(ssIndex int) {

    minLen := 0
    if len(aData)>len(bData) { minLen = len(bData) } else { minLen = len(aData) }
    for i:=0 ; i < minLen ; i++ {
        aPart:=aData[i]
        bPart:=bData[i]
        if(aPart == bPart) {
            ssIndex+=1
        }
    }
    return ssIndex
}


func cutCheck(abcData [][]string, xyzData [][]string)(){

}

func insertCheck(abcData [][]string, xyzData [][]string)(){

}
