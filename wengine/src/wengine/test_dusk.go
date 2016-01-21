package main
import "wengine/dusk"
//import "gopkg.in/mgo.v2/bson"
//import .  "wengine/core/utah"
import "fmt"

func main() {

    d:=dusk.OpenDatabase("mongo","wengine","OpenStack123","127.0.0.1","wengine")
    //user:=&User{Name:"Makush",SecondName:"Lvoloch"}
    //d.CreateUser(user)
    //existing_user,err:=d.GetUserById("60F8FEE2-A6B9-45CF-24CA-B2795002C779")
    //fmt.Printf("--\n%v\n%v\n--",existing_user,err)
    //query:=make(map[string]interface{})
    //query["name"] = "Mike"
    //query["secondname"] = "Livshieshch"
    //existing_user,err:=d.GetUser(query)
    //fmt.Printf("==\n%v\n==\n%v\n==bson==\n%v",existing_user,err,bson.M(query))
    //d.RemoveUsersById("a","b","159E2D96-0AFF-3EBC-D01C-C2E3F3AD16A9")
    //token,err:=d.CreateToken("AF35CEFC-1AEA-A399-7448-C2EF4B80E77F")
    fmt.Printf("token exists %v",d.TokenIsExist("AF35CEFC-1AEA-A399-7448-C2EF4B80E77F", "8D52B9F2-2E19-427F-4E72-04AF9BF91571"))
}
