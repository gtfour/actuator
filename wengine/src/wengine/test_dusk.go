package main
import "wengine/dusk"
import "gopkg.in/mgo.v2/bson"
//import .  "wengine/core/utah"
import "fmt"

func main() {

    d:=dusk.OpenDatabase("mongo","wengine","OpenStack123","127.0.0.1","wengine")
    //user:=&User{Name:"Mike",SecondName:"Livshieshch"}
    //d.CreateUser(user)
    //existing_user,err:=d.GetUserById("60F8FEE2-A6B9-45CF-24CA-B2795002C779")
    //fmt.Printf("--\n%v\n%v\n--",existing_user,err)
    query:=make(map[string]interface{})
    query["name"] = "Mike"
    query["secondname"] = "Livshieshch"
    existing_user,err:=d.GetUser(query)
    fmt.Printf("==\n%v\n==\n%v\n==bson==\n%v",existing_user,err,bson.M(query))
}
