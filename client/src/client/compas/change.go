package compas

type Change struct {
  directive int
  keyData   []string
  newData   []string
  lineNum   int
}


func (c *Change)Perform(data [][]string)(new_data [][]string) {
    return
}

func (changes *changes)Verify(oldData [][]string, newData [][]string) {
    return
}

func GetChanges([][]string,[][]string)(chs changes,err error){
    return chs,err
}
