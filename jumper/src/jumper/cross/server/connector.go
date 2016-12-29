package server

import "jumper/cross"
import "jumper/cross/entrance/mongodb_edge"

func CreateConnector(g *cross.Garreth)(cross.Database,error){
    selected_dbtype := g.GetDbType()
    if !cross.IsOk(selected_dbtype,valid_db_types) {return nil, cross.Selected_dbtype_is_not_ok_on_server_side  }
    if selected_dbtype == cross.MONGODB {
        return mongodb_edge.GetDatabase(g)
    }
    return nil,nil
}
