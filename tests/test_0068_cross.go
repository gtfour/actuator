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
    g.SetPath("/tmp/hello.txt")
    database,err := client.CreateConnector(g)
    myq:=&cross.Query{ Type:cross.GET_ALL }
    _,_=database.RunQuery(myq)



}
