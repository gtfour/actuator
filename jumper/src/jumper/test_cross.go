package main

import "fmt"
import "jumper/cross"
import "jumper/cross/client"
import "jumper/common/gen"

func main() {
    //
    //
    // initianlizing db connection
    dbtype             := "bolt"
    garreth,err        := cross.CreateConnectorTemplate(dbtype)
    if err!=nil {
        fmt.Printf("\n%s is not appropriate\n",dbtype)
    }
    garreth.SetPath("/tmp/cross2.db")
    database,err := client.CreateConnector(garreth)
    defer database.Close()
    fmt.Printf("\ndb:%v open error:%v\ndb connect error:%v\n",database,err,database.Connect())
    //
    // "create_new"-query 
    //
    new_entry_id,err                 := gen.GenId()
    create_key_body                  := make(map[string]interface{},0)
    create_key_body["Id"]            =  new_entry_id
    create_query_body                := make(map[string]interface{},0)
    create_query_body["SourceType"]  =  "file"
    create_query_body["SourcePath"]  =  "/etc/passwd555.so"
    //
    myTempSlice                      := []string {"a","b","c"}
    create_query_body["myTempSlice"] = myTempSlice
    //
    create_query                     :=  cross.Query{Table:"dynimas", Type:cross.CREATE_NEW_IFNOT}
    create_query.QueryBody           =   create_query_body
    create_query.KeyBody             =   create_key_body
    //
    // "get"-query 
    //
    get_query_body                 := make(map[string]interface{},0)
    get_query_body["SourceType"]   =  "file"
    get_query                      := cross.Query{Table:"dynimas", Type:cross.GET_ALL}
    get_query.QueryBody            =  get_query_body
    //
    // "table_create"-query
    //
    maketable_query                := cross.Query{TableList:[]string{"dynimas","activas"}, Type:cross.CREATE_NEW_TABLE_IF_DOESNT_EXIST}
    //
    // "table_check"-query
    //
    table_check_query              := cross.Query{Table:"rytas", Type:cross.CHECK_TABLE_EXIST}
    //
    // "remove"-query
    //
    remove_query                         := cross.Query{Table:"dynimas", Type:cross.REMOVE}
    remove_query_body                    := make(map[string]interface{},0)
    remove_query_body["SourceType"]      =  "dir"
    remove_query.QueryBody               =  remove_query_body
    //
    fmt.Printf("\n---Remove query: %v ---\n",remove_query)
    fmt.Printf("\n---Create query: %v ---\n",create_query)
    //
    get_size_query                         := cross.Query{Table:"dynimas", Type:cross.TABLE_SIZE}
    //
    // "append_to_slice"-query
    //
    appendToSliceQuery                    := cross.Query{Table:"dynimas", Type:cross.APPEND_TO_ARRAY}
    newSlicePropMap                       := make(map[string]interface{}, 0)
    //
    // newEntryIdStr := fmt.Sprintf( "%v", create_query_body)
    // newEntryIdStr                         := fmt.Sprintf( "%v", create_key_body)
    //
    newSlicePropMap["entry_id"]           =  create_key_body
    newSlicePropMap["slice_name"]         =  "myTempSlice"
    appendToSliceQuery.KeyBody            =  newSlicePropMap
    appendToSliceQuery.QueryBody          =  make(map[string]interface{},0)
    appendToSliceQuery.QueryBody["value"] =  "privet" // value to append to slice with name myTempSlice included to mape or bucket identified by create_key_body 
    //
    // "get_slice"-query
    //
    getSliceQuery                 := cross.Query{Table:"dynimas", Type:cross.GET_ARRAY}
    getSliceQuery.KeyBody         =  newSlicePropMap
    //
    // Running queries
    //
    r4,e4 := database.RunQuery( &maketable_query )
    r0,e0 := database.RunQuery( &get_size_query )
    r1,e1 := database.RunQuery( &create_query )
    r2,e2 := database.RunQuery( &remove_query )
    r3,e3 := database.RunQuery( &get_query )
    r5,e5 := database.RunQuery( &table_check_query )
    r6,e6 := database.RunQuery( &get_size_query )
    _, e7 := database.RunQuery( &appendToSliceQuery )
    targetSlice, e8 := database.RunQuery( &getSliceQuery )
    fmt.Printf("Make Tables Query Result:\n%v\nError:%v\n"              ,r4,e4)
    fmt.Printf("Check table %v size: %v Error: %v\n"                       ,get_size_query.Table,r0,e0)
    fmt.Printf("Create new entry:\n%v\nError:%v\n"                      ,r1,e1)
    fmt.Printf("Remove entries by value field:\n%v\nError:%v\n"         ,r2,e2)
    fmt.Printf("Get Query Result:\nError:%v\n",e3)
    for i:= range (*r3) {
        myres:=(*r3)[i]
        fmt.Printf("%v\n",myres)
    }
    fmt.Printf("Check table exist:\n%vError:%v\n"                       ,r5,e5)
    fmt.Printf("Check table %v size: %v Error:%v\n"                       ,get_size_query.Table,r6,e6)
    fmt.Printf("--\nappending to slice err : %v\n--", e7)
    fmt.Printf("--\ngetting slice err : %v\n--Get Result:    %v   \n", e8,targetSlice)
    //
    //
    //
}
