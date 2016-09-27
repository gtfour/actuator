package main

import "fmt"
import "go/importer"

func main() {
    pkg, err := importer.Default().Import("time")
    if err != nil {
        fmt.Printf("error: %s\n", err.Error())
        return
    }
    for _, declName := range pkg.Scope().Names() {
        fmt.Println(declName)
    }
}
