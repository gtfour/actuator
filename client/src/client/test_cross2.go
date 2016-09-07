package main

import "fmt"
import "client/cross"
import "client/common"
import "client/common/types"

var database = cross.Database

func main(){

   new_id,_:=common.GenId()
   key_body   := make(map[string]interface{},0)
   query_body := make(map[string]interface{},0)
   key_body["Id"]          = new_id
   query_body["SourcePath"] ="/etc/passwd"
   query_body["SourceType"] ="file"
   create_new_query:=cross.Query{Table:types.DYNIMAS_T,KeyBody:key_body,QueryBody:query_body,Type:types.CREATE_NEW}

   get_query_body:=make(map[string]interface{},0)
   get_query_body["SourceType"]="file"
   get_query:=cross.Query{Table:types.DYNIMAS_T,QueryBody:query_body,Type:types.GET}

    res1,err1:=database.RunQuery(create_new_query)
    res2,err2:=database.RunQuery(get_query)

    fmt.Printf("Res1:\n%v\n%v\n##################\n",res1,err1)
    fmt.Printf("Res2:\n%v\n%v\n",res2,err2)


}

