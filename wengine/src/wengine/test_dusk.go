package main
import "wengine/dusk"
//import .  "wengine/core/utah"
import "fmt"

func main() {

    d:=dusk.OpenDatabase("mongo","wengine","OpenStack123","127.0.0.1","wengine")
    //user:=&User{Name:"Mike",SecondName:"Livshieshch"}
    //d.CreateUser(user)
    user,err:=d.GetUser("60F8FEE2-A6B9-45CF-24CA-B2795002C779")
    fmt.Printf("--\n%v\n%v\n--",user,err)
}
