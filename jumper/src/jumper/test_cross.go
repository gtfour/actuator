package main

import "fmt"
import "jumper/cross"
import "jumper/cross/client"
import "jumper/common/gen"

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
    // create new query 
    new_entry_id,err                :=  gen.GenId()
    create_key_body                 :=  make(map[string]interface{},0)
    create_key_body["Id"]           =   new_entry_id
    create_query_body               :=  make(map[string]interface{},0)
    create_query_body["SourceType"] =  "file"
    create_query_body["SourcePath"] =  "/etc/passwd.so"
    create_query                    := cross.Query{Table:"dynimas", Type:cross.CREATE_NEW_IFNOT}
    create_query.QueryBody          =  create_query_body
    create_query.KeyBody            =  create_key_body


    // get query 
    get_query_body               := make(map[string]interface{},0)
    get_query_body["SourceType"] =  "file"
    get_query                    := cross.Query{Table:"dynimas", Type:cross.GET_ALL}
    get_query.QueryBody          =  get_query_body
    // table create query
    maketable_query              := cross.Query{TableList:[]string{"dynimas","activas"}, Type:cross.CREATE_NEW_TABLE_IF_DOESNT_EXIST}
    table_check_query            := cross.Query{Table:"rytas", Type:cross.CHECK_TABLE_EXIST}



    // Runing queries
    res0,err0 := database.RunQuery(&create_query)
    res1,err1 := database.RunQuery(&get_query)
    res2,err2 := database.RunQuery(&maketable_query)
    res3,err3 := database.RunQuery(&table_check_query)
    fmt.Printf("Create new entry:\n%v\nError:%v\n"         ,res0,err0)
    fmt.Printf("Get Query Result:\n%v\nError:%v\n"         ,res1,err1)
    fmt.Printf("Make Tables Query Result:\n%v\nError:%v\n" ,res2,err2)
    fmt.Printf("Check table exist:\n%vError:%v\n"          ,res3,err3)



}
