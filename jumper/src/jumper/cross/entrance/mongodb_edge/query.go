package mongodb_edge

import "gopkg.in/mgo.v2/bson"
import "jumper/cross"

/*
type Query struct {
    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}
}*/


func(d *MongoDb)RunQuery(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){

    result_slice:=make([]map[string]interface{},0)
    //result_slice = &result_slice_full
    c      := d.session.DB(d.dbname).C(q.Table)
    if q.Type == cross.CREATE_NEW {
        if q.QueryBody != nil {
            err         = c.Insert(q.QueryBody)
            return nil, err
        } else {
            return nil, cross.EmptyQuery
        }
    } else if q.Type == cross.UPDATE || q.Type == cross.EDIT   {
        if q.KeyBody   == nil { return nil, cross.EmptyKey   }
        if q.QueryBody == nil { return nil, cross.EmptyQuery }
        //
	err = c.Update(bson.M(q.KeyBody), bson.M{"$set": bson.M(q.QueryBody)})
        return nil, err
        //
    } else if q.Type == cross.GET || q.Type == cross.GET_ALL  || q.Type == cross.CHECK_EXIST {
        if q.KeyBody != nil {
            if q.Type == cross.GET_ALL {
                err     =  c.Find(bson.M(q.KeyBody)).All(&result_slice)
                if err == nil {
                    return &result_slice, err
                } else {
                    return nil, err
                }
            } else {
                result:=make(map[string]interface{})
                err     =  c.Find(bson.M(q.KeyBody)).One(&result)
                result_slice = append(result_slice, result)
                if err == nil {
                    return &result_slice, err
                } else {
                    return nil, err
                }
            }
        } else {
            return nil, cross.EmptyKey
        }
    } else  if q.Type == cross.REMOVE {
        if q.KeyBody != nil {
            err    =  c.Remove(bson.M(q.KeyBody))
            return nil, err
        } else {
            return nil, cross.EmptyKey
        }
    } else if q.Type == cross.INSERT_ITEM || q.Type == cross.REMOVE_ITEM {
        if q.KeyBody   == nil { return nil, cross.EmptyKey   }
        if q.QueryBody == nil { return nil, cross.EmptyQuery }

        if q.Type      == cross.INSERT_ITEM {
            err            =  c.Update(bson.M(q.KeyBody), bson.M{"$push":bson.M(q.QueryBody)})
            return nil,err
        } else {
            err            =  c.Update(bson.M(q.KeyBody), bson.M{"$pull":bson.M(q.QueryBody)})
            return nil,err
        }
    } else {
        return nil, cross.IncorrectQueryType
    }

    // err    =  c.Find(bson.M(q.KeyBody)).One(&result)
    // if err != nil {
    //     return result,err
    // }
    // return result, err
}

