package dusk

import "gopkg.in/mgo.v2/bson"
import "wengine/core/types/db_types"

type Query struct {

    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}

}


func(d *MongoDb)RunQuery(q Query)(result map[string]interface{}, err error){
    c      := d.Session.DB(d.dbname).C(q.Table)
    if q.Type == db_types.CREATE_NEW || q.Type == db_types.UPDATE || q.Type == db_types.EDIT   {


    } else if q.Type == db_types.GET || q.Type == db_types.GET_ALL  || q.Type == db_types.CHECK_EXIST {


    } else  if q.Type == db_types.REMOVE {


    } else {
        return nil, incorrect_query_type
    }

    err    =  c.Find(bson.M(q.KeyBody)).One(&result)
    if err != nil {
        return result,err
    }
    return result, err
}
