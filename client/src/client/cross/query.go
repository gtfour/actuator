package cross

import "encoding/json"
import "github.com/boltdb/bolt"
import "client/common/types"

type Query struct {
    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}
}

func (s *Storage)RunQuery(q *Query)(result_slice_addr *[]map[string]interface{}, err error){

    result_slice:=make([]map[string]interface{},0)
    //result_slice = &result_slice_full
    //c      := d.Session.DB(d.dbname).C(q.Table)

    if q.Type == types.CREATE_NEW_TABLE {
        err:=s.Db.Update(func(tx *bolt.Tx) error {
            _,err    := tx.CreateBucketIfNotExists([]byte(q.Table))
            return err
        });
        return nil, err
    } else if q.Type == types.CHECK_TABLE_EXIST  {
        err := s.Db.View(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(q.Table))
            if b==nil{ return table_doesnt_exist } else{
                return nil
            }
        });
        return nil, err
    } else if q.Type == types.REMOVE_TABLE {
        err:=s.Db.Update(func(tx *bolt.Tx) error {
            err:=tx.DeleteBucket([]byte(q.Table))
            return err
        });
        return nil, err
    }  else if q.Type == types.CREATE_NEW {
        if q.QueryBody != nil && q.KeyBody != nil {
            //err         = c.Insert(q.QueryBody)
            key_byte,err_key     := json.Marshal(q.KeyBody)
            query_byte,err_query := json.Marshal(q.QueryBody)
            if err_key!=nil || err_query!=nil {
                return nil, encode_error
            }
            //
            err=s.Db.Update(func(tx *bolt.Tx) error {
                b := tx.Bucket([]byte(q.Table))
                if b==nil { return table_doesnt_exist }
                err:=b.Put(key_byte, query_byte)
            })
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
                //result:=make(map[string]interface{})
                //err     =  c.Find(bson.M(q.KeyBody)).One(&result)
                //result_slice = append(result_slice, result)
                //if err == nil {
                //    return &result_slice, err
                //} else {
                //    return nil, err
                //}
                
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
