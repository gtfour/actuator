package main
import "regexp"
import "fmt"
func main() {

    entry    := "[section]"
    nRx   := regexp.MustCompile(`[[],[]]`)
    match    := nRx.FindAllStringIndex(entry, -1)

    fmt.Printf("%s", match)


}
