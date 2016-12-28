package client

import "jumper/cross"
import "jumper/cross/entrance/boltdb_edge"

func CreateConnector(dbtype string)(*cross.Database,error){
    connectorTemplate,err:=cross.CreateConnectorTemplate( dbtype )
    if err!=nil { return nil,err }
    selected_dbtype := connectorTemplate.GetDbType()
    if !cross.IsOk( selected_dbtype ) {return nil, cross.Selected_dbtype_is_not_ok_on_client_side  }
    return nil,nil
}
