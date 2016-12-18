package file
import "fmt"
import "strings"
import "io/ioutil"
import "jumper/actuator"

func ReadFile(filename string)(lines []string,err error){
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return lines, err
    }
    lines = strings.Split(string(content), "\n")
    return lines, nil
}

func ReadDir(dirname string)(filenames []string,err error){

    mode:=actuator.SAFE_OPENING_MODE
    prop:=actuator.GetProp(dirname,mode)
    fmt.Printf("directory prop:\n%v\n",prop)
    if prop.Error == false && prop.DirContentAvailable == true {
        return prop.DirContent,nil
    } else {
        return nil, prop_error
    }
}
