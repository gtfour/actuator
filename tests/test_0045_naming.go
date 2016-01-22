package main
import "fmt"


type _Test struct {
    Name string
}
type Test struct {
    SecondName string
}


func main() {
    test1:=_Test{"Hello1"}
    test2:=Test{SecondName:"Hello2"}
    fmt.Printf("\n%v\n%v",test1,test2)
}
