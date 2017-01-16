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


    //
    // add pair
    // 
    addpair_query            := cross.Query{Table:"dynimas",Type:cross.ADD_PAIR}
    key_queryA               := make(map[string]interface{},0)
    value_queryA             := make(map[string]interface{},0)
    key_queryA["Id"]         =  "3E4F8EA5-92FF-7C94-0770-9C504D8EEF88"
    first_line               := []string{"a","z"}
    second_line              := []string{"b","y"}
    value_queryA["Data"]     =  [][]string{first_line,second_line}
    addpair_query.KeyBody    =  key_queryA
    addpair_query.QueryBody  =  value_queryA
    //
    // remove pair
    //
    removepair_query           := cross.Query{Table:"dynimas",Type:cross.REMOVE_PAIR}
    key_queryC                 := make(map[string]interface{},0)
    value_queryC               := make(map[string]interface{},0)
    key_queryC["Id"]           =  "3E4F8EA5-92FF-7C94-0770-9C504D8EEF88"
    value_queryC["SourceType"] =  "file"
    removepair_query.KeyBody   =  key_queryC
    removepair_query.QueryBody =  value_queryC
    //
    //
    //
    getentry_query           := cross.Query{Table:"dynimas",Type:cross.GET}
    key_queryB               := make(map[string]interface{},0)
    key_queryB["Id"]         =  "3E4F8EA5-92FF-7C94-0770-9C504D8EEF88"
    getentry_query.KeyBody   =  key_queryB



    maketable_query                := cross.Query{TableList:[]string{"dynimas","activas"}, Type:cross.CREATE_NEW_TABLE_IF_DOESNT_EXIST}
    res0,err0:=database.RunQuery(&maketable_query)

    fmt.Printf("Make table \nResult:\n%v\nError:%v\n",res0,err0)

    resA,errA:=database.RunQuery(&addpair_query)
    fmt.Printf("Add pair query:\n%v\nResult:\n%v\nError:%v\n",addpair_query,resA,errA)

    resB,errB:=database.RunQuery(&getentry_query)
    fmt.Printf("==>Get query:\n%v\nResult:\n%v\nError:%v\n",getentry_query,resB,errB)

    resC,errC:=database.RunQuery(&removepair_query)
    fmt.Printf("Remove pair query:\n%v\nResult:\n%v\nError:%v\n",removepair_query,resC,errC)

    resB,errB=database.RunQuery(&getentry_query)
    fmt.Printf("==>Get query:\n%v\nResult:\n%v\nError:%v\n",getentry_query,resB,errB)
}
