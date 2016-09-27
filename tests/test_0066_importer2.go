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

    // Parse the input string, []byte, or io.Reader,
    // recording position information in fset.
    // ParseFile returns an *ast.File, a syntax tree.
    f, err := parser.ParseFile(fset, "/actuator/wapour/src/wapour/settings/settings.go", nil, 0)
    if err != nil {
        log.Fatal(err) // parse error
    }

    // A Config controls various options of the type checker.
    // The defaults work fine except for one setting:
    // we must specify how to deal with imports.
    conf := types.Config{Importer: importer.Default()}

    // Type-check the package containing only file f.
    // Check returns a *types.Package.
    pkg, err := conf.Check("wapour/settings", fset, []*ast.File{f}, nil)
    if err != nil {
        log.Fatal(err) // type error
    }

    fmt.Printf("Package  %q\n", pkg.Path())
    fmt.Printf("Name:    %s\n", pkg.Name())
    fmt.Printf("Imports: %s\n", pkg.Imports())
    fmt.Printf("Scope:   %s\n SERVER_ADDR %v\n", pkg.Scope(),pkg.Scope().Lookup("SERVER_ADDR"))
}
