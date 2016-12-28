package client

import "jumper/cross"
import "jumper/cross/entrance/boltdb_edge"

func CreateConnector(g *cross.Garreth)(cross.Database,error){
    selected_dbtype := g.GetDbType()
    if !cross.IsOk(selected_dbtype,valid_db_types) {return nil, cross.Selected_dbtype_is_not_ok_on_client_side  }
    if selected_dbtype == cross.BOLTDB {
        return boltdb_edge.GetDatabase(g)


    }

    return nil,nil
}
