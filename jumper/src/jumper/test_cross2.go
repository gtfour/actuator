package main

import "fmt"
import "jumper/cross"
import "jumper/cross/client"

func main(){


    dbtype             := "bolt"
    garreth,err        := cross.CreateConnectorTemplate(dbtype)
    if err!=nil {
        fmt.Printf("\n%s is not appropriate\n",dbtype)
    }
    garreth.SetPath("/tmp/cross2.db")
    database,err := client.CreateConnector(garreth)
    database.Connect()
    defer database.Close()

    addpair_query           := cross.Query{Table:"dynimas",Type:cross.ADD_PAIR}
    key_queryA               := make(map[string]interface{},0)
    value_queryA             := make(map[string]interface{},0)
    key_queryA["Id"]         =  "B0E34597-AB16-8CDD-9E89-C8D1EC2DC134"
    first_line              := []string{"a","z"}
    second_line             := []string{"b","y"}
    value_queryA["Data"]     =  [][]string{first_line,second_line}
    addpair_query.KeyBody   =  key_queryA
    addpair_query.QueryBody =  value_queryA


    getentry_query              := cross.Query{Table:"dynimas",Type:cross.GET}
    key_queryB                  := make(map[string]interface{},0)
    key_queryB["Id"]            = "B0E34597-AB16-8CDD-9E89-C8D1EC2DC134"
    getentry_query.KeyBody = key_queryB


    resA,errA:=database.RunQuery(&addpair_query)
    fmt.Printf("Add pair query:\n%v\nResult:\n%v\nError:%v\n",addpair_query,resA,errA)

    resB,errB:=database.RunQuery(&getentry_query)
    fmt.Printf("Get query:\n%v\nResult:\n%v\nError:%v\n",getentry_query,resB,errB)
}
