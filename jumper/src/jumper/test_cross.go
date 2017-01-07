package main

import "fmt"
import "jumper/cross"
import "jumper/cross/client"

func main() {

    // initianlizing db connection
    dbtype       := "bolt"
    garreth,err        := cross.CreateConnectorTemplate(dbtype)
    if err!=nil {
        fmt.Printf("\n%s is not appropriate\n",dbtype)
    }
    garreth.SetPath("/tmp/cross.db")
    database,err := client.CreateConnector(garreth)
    fmt.Printf("\ndb:%v open error:%v\ndb connect error:%v\n",database,err,database.Connect())
    // 


    // get query 
    get_query_body               := make(map[string]interface{},0)
    get_query_body["SourceType"] = "file"
    get_query                    := cross.Query{Table:"dynimas",QueryBody:get_query_body,Type:cross.GET_ALL}
    // table create query
    maketable_query              :=cross.Query{TableList:[]string{"dynimas","activas"}, Type:cross.CREATE_NEW_TABLE_IF_DOESNT_EXIST}
    table_check_query            :=cross.Query{Table:"rytas", Type:cross.CHECK_TABLE_EXIST}



    // Runing queries
    res1,err1:=database.RunQuery(&get_query)
    res2,err2:=database.RunQuery(&maketable_query)
    res3,err3:=database.RunQuery(&table_check_query)
    fmt.Printf("Get Query Result:\n%v\nError:%v\n",res1,err1)
    fmt.Printf("Make Tables Query Result:\n%v\nError:%v\n",res2,err2)
    fmt.Printf("Check table exist:\n%vError:%v\n",res3,err3)



}
