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
    if q.Type == db_types.CREATE_NEW {
        if q.QueryBody != nil {
            err         = c.Insert(q.QueryBody)
            return nil, err
        } else {
            return nil, empty_query
        }
    } else if q.Type == db_types.UPDATE || q.Type == db_types.EDIT   {

        if q.KeyBody   == nil { return nil,empty_key   }
        if q.QueryBody == nil { return nil,empty_query }

    } else if q.Type == db_types.GET || q.Type == db_types.GET_ALL  || q.Type == db_types.CHECK_EXIST {
        if q.KeyBody != nil {
            if q.Type == db_types.GET_ALL {
                err     =  c.Find(bson.M(q.KeyBody)).All(&result)
            } else {
                err     =  c.Find(bson.M(q.KeyBody)).One(&result)
            }
            if err == nil {
                return result, err
            } else {
                return nil, err
            }
        } else {
            return nil, empty_key
        }
    } else  if q.Type == db_types.REMOVE {
        if q.KeyBody != nil {
            err    =  c.Remove(bson.M(q.KeyBody))
            return nil, err
        } else {
            return nil, empty_key
        }
    } else if q.Type == db_types.INSERT_ITEM || q.Type == db_types.REMOVE_ITEM {
        if q.KeyBody   == nil { return nil,empty_key   }
        if q.QueryBody == nil { return nil,empty_query }

        if q.Type      == db_types.INSERT_ITEM {
            err            =  c.Update(bson.M(q.KeyBody), bson.M{"$push":bson.M(q.QueryBody)})
        } else {
            err            =  c.Update(bson.M(q.KeyBody), bson.M{"$pull":bson.M(q.QueryBody)})
        }


    } else {
        return nil, incorrect_query_type
    }

    err    =  c.Find(bson.M(q.KeyBody)).One(&result)
    if err != nil {
        return result,err
    }
    return result, err
}
