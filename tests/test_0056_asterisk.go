package main
import "fmt"

type A struct {
    array []int
}


func (a *A) AppendAsterisk (value int) {
    a.array = append(a.array, value)
}

func (a A) Append (value int) {
     a.array = append(a.array, value)
}

func (a *A) GetAsterisk ()([]int) {
    return a.array
}

func (a A) Get()([]int) {
    return a.array
}

func main() {

    t1:=&A{}
    t2:=&A{}
    t1.AppendAsterisk(1)
    t2.Append(1)
    fmt.Printf("Append:\nt1:%v\nt2:%v\n---\n",t1,t2)
    fmt.Printf("Get:\nt1:%v\nt2:%v\n",t1.Get(),t2.Get())


}
