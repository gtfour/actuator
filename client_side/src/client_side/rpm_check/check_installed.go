package main

import "client_side/protodb"
import "github.com/golang/protobuf/proto"
import "github.com/golang/protobuf/proto/proto3_proto"
import "fmt"

var packages_db_file = "/var/lib/rpm/Dirnames"
var environment_dir = "/var/lib/rpm"

func main(){

    environment_config:=&protodb.EnvironmentConfig{Create:false}

    environment,_:=protodb.OpenEnvironment(environment_dir,environment_config)

    transaction:=protodb.NoTransaction

    //database_config:=DatabaseConfig{Create:false}

    database,_ :=protodb.OpenDatabase(environment,transaction,packages_db_file,&protodb.DatabaseConfig{Create:false})

      defer database.Close()

    database_type,_:=database.Type()

    //message:=proto.Message {}
    message:=proto.proto3_proto.ProtoMessage()

    records:=database.Get(transaction,false,message)

        fmt.Println(records)

        fmt.Println("------")

        fmt.Println(message)

        fmt.Println("------")

        fmt.Println(database_type)


}


