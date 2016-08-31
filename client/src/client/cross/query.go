package cross

import "client/common/types"

type Query struct {
    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}
}

func (s *Storage)RunQuery(query *Query)(result_slice_addr *[]map[string]interface{}, err error){

    result_slice:=make([]map[string]interface{},0)
    //result_slice = &result_slice_full
    c      := d.Session.DB(d.dbname).C(q.Table)
    if q.Type == types.CREATE_NEW {
        if q.QueryBody != nil {
            err         = c.Insert(q.QueryBody)
            return nil, err
        } else {
            return nil, empty_query
        }
    } else if q.Type == types.UPDATE || q.Type == types.EDIT   {
        if q.KeyBody   == nil { return nil,empty_key   }
        if q.QueryBody == nil { return nil,empty_query }
        //
        err = c.Update(bson.M(q.KeyBody), bson.M{"$set": bson.M(q.QueryBody)})
        return nil, err
        //
    } else if q.Type == types.GET || q.Type == types.GET_ALL  || q.Type == types.CHECK_EXIST {
        if q.KeyBody != nil {
            if q.Type == types.GET_ALL {
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
            return nil, empty_key
        }
    } else  if q.Type == types.REMOVE {
        if q.KeyBody != nil {
            err    =  c.Remove(bson.M(q.KeyBody))
            return nil, err
        } else {
            return nil, empty_key
        }
    } else if q.Type == types.INSERT_ITEM || q.Type == types.REMOVE_ITEM {
        if q.KeyBody   == nil { return nil,empty_key   }
        if q.QueryBody == nil { return nil,empty_query }

        if q.Type      == types.INSERT_ITEM {
            err            =  c.Update(bson.M(q.KeyBody), bson.M{"$push":bson.M(q.QueryBody)})
            return nil,err
        } else {
            err            =  c.Update(bson.M(q.KeyBody), bson.M{"$pull":bson.M(q.QueryBody)})
            return nil,err
        }
    } else {
        return nil, incorrect_query_type
    }

    // err    =  c.Find(bson.M(q.KeyBody)).One(&result)
    // if err != nil {
    //     return result,err
    // }
    // return result, err





}

