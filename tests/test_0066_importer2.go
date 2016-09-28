package main

import (
    "fmt"
    "go/ast"
    "go/importer"
    "go/parser"
    "go/token"
    "go/types"
    "log"
)


func main() {
    fset := token.NewFileSet()

    f, err := parser.ParseFile(fset, "/actuator/wapour/src/wapour/settings/settings.go", nil, 0)
    if err != nil {
        log.Fatal(err)
    }

    conf := types.Config{Importer: importer.Default()}

    pkg, err := conf.Check("wapour/settings", fset, []*ast.File{f}, nil)
    if err != nil {
        log.Fatal(err) // type error
    }

    fmt.Printf("Package  %q\n", pkg.Path())
    fmt.Printf("Name:    %s\n", pkg.Name())
    fmt.Printf("Imports: %s\n", pkg.Imports())

    my_var:=pkg.Scope().Lookup("SERVER_ADDR").(*types.Const)
    var server_addr string
    server_addr=my_var.Val().String()
    fmt.Printf("Scope:   %s\n\n\n%s\n", pkg.Scope(),server_addr)
}
