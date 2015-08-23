package deb_check

import "io"
import "os"
import "bufio"
import "strings"

// StatusEntry

func ParseFile(filename string) (statusfile StatusFile, err error) {


    file, err := os.Open(filename)
    status_file:=StatusFile{}

    if err!=nil {

        return status_file,err

    }

    buffered_reader:=bufio.NewReader(file)
    eof := false

    for lino := 1; !eof; lino++ {

        status_entry:=StatusEntry{}

        line, err := buffered_reader.ReadString('\n')

        if err == io.EOF {
            err = nil
            eof = true 
        } else if err != nil {
            return status_file, err
        }

        if ( strings.HasPrefix(line, "Package") || strings.HasPrefix(line, "Architecture") || strings.HasPrefix(line, "Version")){
            status_entry.ParseField(line)
        }
        if status_entry.Complete {

          status_file.InstalledPackages=append(status_file.InstalledPackages,status_entry)
          status_entry=StatusEntry{}

        }

    }

    return status_file,nil


}

func (status_entry *StatusEntry)ParseField(line string){

    words:=strings.Split(line,": ")

    if len(words)==2 {

        if words[0]== "Package" { status_entry.Name = words[1] }
        if words[0]== "Architecture" { status_entry.Architecture = words[1] }
        if words[0]== "Version" { status_entry.Version = words[1] ; status_entry.Complete = true}

}
}
