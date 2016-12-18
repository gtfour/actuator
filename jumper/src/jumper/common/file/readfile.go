package file
import "strings"
import "io/ioutil"

func ReadFile(filename string)(lines []string,err error){
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return lines, err
    }
    lines = strings.Split(string(content), "\n")
    return lines, nil
}

func ReadDir(dirname string)(filenames []string,err error){



}
