package main
import "wengine/dusk"
import .  "wengine/core/utah"
//import "fmt"

func main() {

    d:=dusk.OpenDatabase("mongo","wengine","OpenStack123","127.0.0.1","wengine")
    user:=&User{Name:"John",SecondName:"Smith"}
    d.CreateUser(user)
}
