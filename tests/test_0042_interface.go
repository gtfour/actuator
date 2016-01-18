package main
import "fmt"
//import "reflect"


type Test interface {
    Hello(string)
}
type Test1 struct {
    Field1 string
}
type Test2 struct {
    Field2 string
}

func(t1 Test1) Hello() (string ) {
    return "hello1"
}

func (t2 Test2) Hello() (string) {
    return "hello2"
}

func CreateTest (ttype string) (t *Test ) {

    if ttype == "test1" {
        return &Test1{}
    }
    if ttype == "test2" {
        return &Test2{}
    }
        return nil
}

func main() {

    test1 := CreateTest("test1")
    test2 := CreateTest("test2")
    fmt.Printf("\n 1 |%s|\n",test1.Field1)
    fmt.Printf("\n 2 |%s|\n",test2.Field2)

}
