package client

import "jumper/cross"

func CreateConnector(dbtype string)(*cross.Database,error){
    connectorTemplate,err:=cross.CreateConnectorTemplate(dbtype)
    if err!=nil {return nil,err}
    if !cross.IsOk(connectorTemplate.Dbtype, valid_db_types) {return nil, cross.Selected_dbtype_is_not_ok_on_client_side  }
    return nil,nil



}

