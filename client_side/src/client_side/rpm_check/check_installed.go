package main

import "client_side/protodb"
import "github.com/golang/protobuf/proto"
// proto3_proto
//import "github.com/golang/protobuf/proto/proto3_proto"
import "fmt"

var packages_db_file = "/var/lib/rpm/Packages"
var environment_dir = "/var/lib/rpm"

func main(){

    environment_config:=&protodb.EnvironmentConfig{Create:false}

    environment,_:=protodb.OpenEnvironment(environment_dir,environment_config)

    transaction:=protodb.NoTransaction

    database,_ :=protodb.OpenDatabase(environment,transaction,packages_db_file,&protodb.DatabaseConfig{Create:false})

      defer database.Close()

    database_type,_:=database.Type()

    cursor,err:=database.Cursor(transaction)

    if err!= nil {

       return

    }

    var message proto.Message
    var cursor_message proto.Message
    err=cursor.First(cursor_message)
    fmt.Println(" ------ ")
    //fmt.Println(err)
    //fmt.Println(cursor)
    fmt.Println(" ------ ")
    //fmt.Println(cursor_message)
    err=database.Get(transaction,false,message)
    fmt.Println(" ------ ")
    fmt.Println(err)
    fmt.Println(message)
    fmt.Println(" ------ ")
    fmt.Println(database_type)

}


