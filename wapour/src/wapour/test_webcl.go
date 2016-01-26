package main
import . "wapour/api/webclient"
import "fmt"

func main() {

    p,err:=Init("wengine","http://127.0.0.1:9000","Anna","SecretPassword123")
    fmt.Printf("Proxy:%v\nError:%v",p,err)

}
