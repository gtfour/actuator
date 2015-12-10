package main
import "fmt"

func Greeting(who ...string) {

    for i:=range who {

        fmt.Printf("\n%s",who[i])

    }



}

func main() {

    Greeting("Sandra","John")





}
