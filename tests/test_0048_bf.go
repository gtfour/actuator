package main
import "time"
import "fmt"

type Test struct {
    FieldOne string
}

func Init ()(*Test){
    t := &Test{}
    go func() {
        time.Sleep(400 * time.Millisecond)
        t.FieldOne = "Hello"
    }()
    return t
}


func main(){
    test := Init()
    fmt.Printf("\n%v\n",test)
    time.Sleep(400 * time.Millisecond)
    fmt.Printf("\n%v\n",test)
}
