package main
import (

        "fmt"

        //"log"

        //"gopkg.in/mgo.v2"

        //"gopkg.in/mgo.v2/bson"

        "wengine_parts/repository"
)
func main() {

    fmt.Println("====")

    my_repo_relation:= repository.RpmRelation {Name: "/bin/bash"}

    fmt.Println(my_repo_relation)

}

