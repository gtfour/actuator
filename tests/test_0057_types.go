package main
import "fmt"

type Hello int


func (h *Hello)Inc()() {
    new_value:=&Hello{}
    h=&new_value+1
}

func main(){
    var test Hello
    test.Inc()
    fmt.Printf("\n%d\n",test)


}
