package solis

//import "crypto/rsa"
import "fmt"
import "client/settings"
import "client/cross"
import "client/common"
import "client/common/types"

var database = cross.Database
var pubkey_path     = settings.PUBKEY_PATH
var privatekey_path = settings.PRIVATEKEY_PATH

func GenPrivateKey()() {


}

func SaveKeyToFile()() {


}

func LoadKeyFromFile()(){


}

func Preparing()(error){

    err:=CheckSolisDbTable()
    if err!=nil {
        err:=CreateSolisDbTable()
        if err!=nil{
            return err
        }
    } else {
        err:=CheckHostId()
        if err!= nil {
            err=SetNewHostId()
            if err!=nil{
                return err
            }
        }
    }
    fmt.Printf("\n Finished \n")
    return nil

}

func CheckSolisDbTable()(error){

    check_table_query    := cross.Query{Table:types.SOLIS_T, Type:types.CHECK_TABLE_EXIST}
    _,err                := database.RunQuery(check_table_query)
    //fmt.Printf("\nChecking Solis Table Err: %v\n",err)
    return err

}

func CreateSolisDbTable()(error){

    create_table_query    := cross.Query{Table:types.SOLIS_T, Type:types.CREATE_NEW_TABLE}
    _,err                := database.RunQuery(create_table_query)
    //fmt.Printf("\nCreating Solis Table Err: %v\n",err)
    return err

}

func CheckHostId()(error){

    key_body:=make(map[string]interface{},0)
    key_body["key"]="host_id"
    check_query:=cross.Query{Table:types.SOLIS_T,Type:types.CHECK_EXIST,KeyBody:key_body}
    _,err:=database.RunQuery(check_query)
    return err


}

func SetNewHostId()(error){
    new_id,_            := common.GenId()
    key_body            := make(map[string]interface{},0)
    query_body          := make(map[string]interface{},0)
    key_body["key"]     = "host_id"
    query_body["query"] = new_id
    create_query:=cross.Query{Type:types.CREATE_NEW,Table:types.SOLIS_T,KeyBody:key_body,QueryBody:query_body}

    _,err:=database.RunQuery(create_query)
    return err
}

