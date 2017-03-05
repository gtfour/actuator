package file
//import "fmt"
import "strings"
import "io/ioutil"
import "jumper/actuator"

func ReadFile(filename string)(lines []string,err error){
    //
    //
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return lines, err
    }
    lines = strings.Split(string(content), "\n")
    return lines, nil
    //
    //
}

func ReadDir(dirname string)([]string,error){
    //
    // returning filenames inside specified dir
    // 
    // Seems i can use SAFE_OPENING_MODE there because this function will calling only when directory content would be changed and there is
    // no big impact  to performance
    //  
    mode := actuator.SAFE_OPENING_MODE
    prop := actuator.GetProp(dirname, mode)
    // fmt.Printf("directory prop:\n%v\n",prop)
    if prop.Error == false && prop.DirContentAvailable == true {
        return prop.DirContent, nil
    } else {
        return nil, prop_error
    }
}

func ReadDirFiles(dirname string)([]string,error){
    //
    mode := actuator.SAFE_OPENING_MODE
    prop := actuator.GetProp(dirname, mode)
    //
    var directory_files []string
    if prop.Error == false && prop.DirContentAvailable == true {
        if strings.HasSuffix(dirname, "/") == false { dirname = dirname+"/" }
        var directory_items []string
        for i:= range prop.DirContent {
            item            :=  dirname+prop.DirContent[i]
            directory_items =   append(directory_items, item)
        }
        //
        for i:= range directory_items {
            item      := directory_items[i]
            item_prop := actuator.GetProp(item, mode)
            if prop.Error == false &&  item_prop.IsRegular == true && item_prop.IsReadable == true {
                // appending just only filenames , not directory names
                directory_files = append(directory_files, item)
            }

        }
        return directory_files, nil
    } else {
        return nil, prop_error
    }
}
