package main
import "fmt"

func main() {

test_num:= [10]int {0,1,2,3,4,5,6,7,8,9}
test_abc:= [6]string {"a","b","c","d","e"}

for i := range test_num {

    fmt.Printf("%d--\n",test_num[i])
    for z:= range test_abc {
        fmt.Printf("%s>>\n",test_abc[z])
        if test_abc[z] == "c" {break}

    }


}

}
