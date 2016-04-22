package main
import ("os" ; "fmt")
import "reflect"
import "syscall"



func main() {


    fi, err:=os.Stat("/tmp")

    if err!=nil {return}

    mtime:=fi.ModTime()

    test:=fi.Sys()

    test2,found :=test.(*syscall.Stat_t)

    if found==false { fmt.Println("error") ; return }

    fmt.Println(test2.Ino)

    fmt.Println(reflect.TypeOf(test))

    val := reflect.Indirect(reflect.ValueOf(test))

    //value:=test.Value

    fmt.Printf("Second interface field : %s ",val.Type().Field(1).Name)


    //fmt.Println(test.Ino)

    fmt.Println(mtime)

    fmt.Println("---")


}
