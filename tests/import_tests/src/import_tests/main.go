package main

import "import_tests/a"
import "import_tests/b"
import "fmt"

func main() {
    _ = b.Test
    fmt.Printf("\n%d\n",a.A)
    b.Increase()
    fmt.Printf("\n%d\n",a.A)
    fmt.Printf("\nGetA: %d\n",a.GetA())
}
