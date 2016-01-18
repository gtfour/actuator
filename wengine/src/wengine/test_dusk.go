package main
import "wengine/dusk"
import .  "wengine/core/utah"

func main() {

    d:=dusk.OpenDatabase("mongo","root","OpenStack123","127.0.0.1","wengine")
    user:=&User{Id:"123",Name:"John"}
    d.CreateUser(user)
}
