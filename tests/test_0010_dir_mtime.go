package main
import ("os" ; "fmt")
import "reflect"

func main() {

    fi, _:=os.Stat("/tmp/test.txt")

    mtime:=fi.ModTime()

    test:=fi.Sys()

    fmt.Println(reflect.TypeOf(test))

    val := reflect.Indirect(reflect.ValueOf(test))

    fmt.Println(val.Type().Field(1).Name)

    //fmt.Println(test.Ino)

    fmt.Println(mtime)

    fmt.Println("---")

    fmt.Println(test)

    fmt.Println(test.Ino)

}
