package compas

// do not forget to match line number
// uncertain module

// abcData and xyzData is data parsed by cuda

type changes []Change

type CmpDataProp struct {


}

type CmpDataEntryProp struct {

    oldLineNumber  int
    newLineNumber  int
    statuses       []int
    value          []string

}

type Change struct {

  // modify orig
  directives  []int
  data        [][]string
  dataLine    []string
  start


}

//func(c *Cmp)GetChanges()() {
//}

func (c *Change)Perform(data [][]string)(status int) {
    return
}

func (chs *changes)Verify(oldData [][]string, newData [][]string) {



}




func Trite(abcData [][]string, xyzData [][]string) {
    for _= range abcData {

    }
}

func TriteLine(abcDataLine []string, xyzDataLine [][]string)(status int) {

    if len(abcDataLine) == len(xyzDataLine) {

    } else {

    }
    return status

}

func cutCheck(abcData [][]string, xyzData [][]string)(){

}

func insertCheck(abcData [][]string, xyzData [][]string)(){

}
