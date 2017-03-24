package file

import "os"
import "io"
import "fmt"
import "bufio"
import "strings"
import "io/ioutil"
import "jumper/actuator"




func ReadFile(filename string)(lines []string,err error){
    //
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return lines, err
    }
    lines = strings.Split(string(content), "\n")
    return lines, nil
    //
}

func ReadFileWithOffset( filename string, offset int64 )( lines []string,new_offset int64 ,err error ) {
    if filename != "" {

        file, err := os.Open(filename)
        defer file.Close()

        if err != nil {
            return lines,0,err
        }

        var ORIGIN int     =   0
        _,err =  file.Seek( offset, ORIGIN )
        if err == nil {
            reader        := bufio.NewReader(file)
            file_info,err := file.Stat()
            if err != nil {
                return lines, 0, err
            }
            new_offset = file_info.Size()
            fmt.Printf("--\nFile size: %v\n--", new_offset)
            //
            eof := false
            for !eof {
                var line string
                line, err = reader.ReadString('\n')
                if err == io.EOF || err == nil {
                    err = nil   // io.EOF isn't really an error
                    eof = true  // this will end the loop at the next iteration
                    lines = append(lines, line)
                } else if err != nil {
                    break
                }
            }
            //
        }
        return lines,new_offset,err

    } else {
        err = filename_is_empty
        return
    }
}

func ReadDirContent(dirname string)([]string,error){
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
