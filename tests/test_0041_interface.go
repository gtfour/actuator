package main
import "fmt"

type AbTest interface{

    GetUsers()([]string)

}

type Mongodb struct {
    Users []string
}

type Postgresql struct {
    Users []string
}

func (d *Mongodb)GetUsers()([]string) {
    return d.Users
}

func (d *Postgresql)GetUsers()([]string) {
    return d.Users
}






func GetUsers(t AbTest)([]string) {
    return t.GetUsers()
}

func CreateConnection ( contype  string ) ( t AbTest ) {

    if contype == "mongodb" {
        return &Mongodb{Users:[]string{"user1","user2"}}
    }
    if contype == "postgresql" {
        return &Postgresql{Users:[]string{"user3","user4"}}
    }
    return nil
}

func main() {

    mongodb_instance    := CreateConnection("mongodb")
    postgresql_instance := CreateConnection("postgresql")
    fmt.Printf("Mongo users:  %v", mongodb_instance.GetUsers())
    fmt.Printf("Postgresql users:  %v", postgresql_instance.GetUsers())

}
