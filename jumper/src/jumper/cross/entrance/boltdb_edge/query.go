package boltdb_edge

import "encoding/json"
import "github.com/boltdb/bolt"
import "jumper/cross"


func (s *Storage)RunQuery(q cross.Query)(result_slice_addr *[]map[string]interface{}, err error){


    result_slice:=make([]map[string]interface{},0)
    //result_slice = &result_slice_full
    //c      := d.Session.DB(d.dbname).C(q.Table)

    if q.Type == cross.CREATE_NEW_TABLE {
        err:=s.db.Update(func(tx *bolt.Tx) error {
            _,err    := tx.CreateBucketIfNotExists([]byte(q.Table))
            return err
        });
        return nil, err
    } else if q.Type == cross.CHECK_TABLE_EXIST  {
        err := s.db.View(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(q.Table))
            if b==nil{ return cross.TableDoesntExist } else{
                return nil
            }
        });
        return nil, err
    } else if q.Type == cross.REMOVE_TABLE {
        err:=s.db.Update(func(tx *bolt.Tx) error {
            err:=tx.DeleteBucket([]byte(q.Table))
            return err
        });
        return nil, err
    } else if q.Type == cross.CREATE_NEW {
        if q.QueryBody != nil && q.KeyBody != nil {
            //err         = c.Insert(q.QueryBody)
            key_byte,err_key     := json.Marshal(q.KeyBody)
            query_byte,err_query := json.Marshal(q.QueryBody)
            if err_key!=nil || err_query!=nil {
                return nil, cross.EncodeError
            }
            //
            err=s.db.Update(func(tx *bolt.Tx) error {
                table := tx.Bucket([]byte(q.Table))
                if table==nil { return cross.TableDoesntExist }

                entry := table.Get(key_byte)
                if entry == nil {
                    err:=table.Put(key_byte, query_byte)
                    return err
                } else {
                    return cross.EntryAlreadyExist
                }

                err:=table.Put(key_byte, query_byte)
                return err
            })
            return nil, err
        } else {
            return nil, cross.EmptyQuery
        }

    } else if q.Type == cross.CREATE_NEW_IFNOT {


    } else if q.Type == cross.REPLACE {
        if q.QueryBody != nil && q.KeyBody != nil {
            key_byte,err_key     := json.Marshal(q.KeyBody)
            query_byte,err_query := json.Marshal(q.QueryBody)
            if err_key!=nil || err_query!=nil {
                return nil, cross.EncodeError
            }
            err=s.db.Update(func(tx *bolt.Tx) error {
                table := tx.Bucket([]byte(q.Table))
                if table==nil { return cross.TableDoesntExist }
                entry := table.Get(key_byte)
                if entry == nil {
                    return cross.EntryDoesntExist
                } else {
                    err:=table.Put(key_byte, query_byte)
                    return err
                }
            })
        }
    } else if q.Type == cross.UPDATE || q.Type == cross.EDIT   {
        if q.KeyBody   == nil { return nil, cross.EmptyKey   }
        if q.QueryBody == nil { return nil, cross.EmptyQuery }
        //
        //err = c.Update(bson.M(q.KeyBody), bson.M{"$set": bson.M(q.QueryBody)})
        return nil, err
        //
    } else if q.Type == cross.GET || q.Type == cross.GET_ALL  || q.Type == cross.CHECK_EXIST {
        //result_slice_addr,err := s.RunQueryGet(q)
        //return result_slice_addr, err
    } else  if q.Type == cross.REMOVE {
        if q.KeyBody != nil {
            //err    =  c.Remove(bson.M(q.KeyBody))
            return nil, err
        } else {
            return nil, cross.EmptyKey
        }
    } else if q.Type == cross.INSERT_ITEM || q.Type == cross.REMOVE_ITEM {
        if q.KeyBody   == nil { return nil, cross.EmptyKey   }
        if q.QueryBody == nil { return nil, cross.EmptyQuery }

        if q.Type      == cross.INSERT_ITEM {
            //err            =  c.Update(bson.M(q.KeyBody), bson.M{"$push":bson.M(q.QueryBody)})
            return nil,err
        } else {
            //err            =  c.Update(bson.M(q.KeyBody), bson.M{"$pull":bson.M(q.QueryBody)})
            return nil,err
        }
    } else {
        return nil, cross.IncorrectQueryType
    }
    // err    =  c.Find(bson.M(q.KeyBody)).One(&result)
    // if err != nil {
    //     return result,err
    // }
    return &result_slice, err
}
