package main
import "fmt"

type Test func(string)(string)

func (t Test) Bfuck (int) {
    fmt.Printf("%s -- %s\n",t("2"),t("d"))
}

func main() {

    var test21 Test

    test21 = func(hello string)(hello_ex string) { return hello+"ex"  }
    test22 := Test(func(hello string)(hello_ex string) { return hello+"ex"  })

    test21.Bfuck(2)
    test22.Bfuck(2)


}

