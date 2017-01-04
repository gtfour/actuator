package main

import "fmt"
import "jumper/cross"
import "jumper/cross/client"

func main() {

    dbtype       := "bolt"
    g,err        := cross.CreateConnectorTemplate(dbtype)
    if err!=nil {
        fmt.Printf("\n%s is not appropriate\n",dbtype)
    }
    g.SetPath("/tmp/cross.db")
    database,err := client.CreateConnector(g)
    fmt.Printf("\ndb:%v open error:%v\ndb connect error:%v\n",database,err,database.Connect())

    //myq:=&cross.Query{ Type:cross.GET_ALL }

    get_query_body               := make(map[string]interface{},0)
    get_query_body["SourceType"] = "file"
    get_query                    := cross.Query{Table:"dynimas",QueryBody:get_query_body,Type:cross.GET_ALL}


    res,err:=database.RunQuery(&get_query)
    fmt.Printf("Get Query Result:\n%v\nError:%v\n",res,err)



}
