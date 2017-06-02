package boltdb_edge

import "encoding/json"
import "github.com/boltdb/bolt"
import "jumper/cross"


func (d *Database)RunQuery(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){


    result_slice:=make([]map[string]interface{},0)
    //result_slice = &result_slice_full
    //c      := d.Session.DB(d.dbname).C(q.Table)
    switch query_type:=q.Type; query_type {
        case cross.CREATE_NEW_TABLE:
            err:=d.db.Update(func(tx *bolt.Tx) error {
                _,err    := tx.CreateBucketIfNotExists([]byte(q.Table))
                return err
            });
            return nil, err
        case cross.CHECK_TABLE_EXIST:
            err := d.db.View(func(tx *bolt.Tx) error {
                b:=tx.Bucket([]byte(q.Table))
                if b==nil{ return cross.TableDoesntExist } else{
                    return nil
                }
            });
            return nil, err
        case cross.REMOVE_TABLE:
            err:=d.db.Update(func(tx *bolt.Tx) error {
                err:=tx.DeleteBucket([]byte(q.Table))
                return err
            });
            return nil, err
        case cross.CREATE_NEW, cross.CREATE_NEW_IFNOT, cross.UPDATE, cross.EDIT:
            res,err:=d.CreateNew(q)
            return res,err
        case cross.REPLACE:
            if q.QueryBody != nil && q.KeyBody != nil {
                key_byte,err_key     := json.Marshal(q.KeyBody)
                query_byte,err_query := json.Marshal(q.QueryBody)
                if err_key!=nil || err_query!=nil {
                    return nil, cross.EncodeError
                }
                err=d.db.Update(func(tx *bolt.Tx) error {
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
        case cross.GET, cross.GET_ALL, cross.CHECK_EXIST:
            //result_slice_addr,err := s.RunQueryGet(q)
            //return result_slice_addr, err
            //fmt.Printf("Database:\n%v\n",d)
            res,err := d.Get(q)
            return res,err
        case cross.REMOVE:
            res,err:=d.Remove(q)
            return res,err
        case cross.INSERT_ITEM, cross.REMOVE_ITEM:
            //if q.KeyBody   == nil { return nil, cross.EmptyKey   }
            //if q.QueryBody == nil { return nil, cross.EmptyQuery }

            if q.Type      == cross.INSERT_ITEM {
                //err            =  c.Update(bson.M(q.KeyBody), bson.M{"$push":bson.M(q.QueryBody)})
                return nil,err
            } else {
                //err            =  c.Update(bson.M(q.KeyBody), bson.M{"$pull":bson.M(q.QueryBody)})
                return nil,err
            }
        case cross.CREATE_NEW_TABLE_IF_DOESNT_EXIST:
            res,err:=d.CreateNewTableIfDoesntExist(q)
            return res,err
        case cross.ADD_PAIR, cross.REMOVE_PAIR:
            res,err:=d.ModifyPair(q)
            return res,err
        case cross.GET_PAIR:
            res,err:=d.GetPair(q)
            return res,err

        case cross.APPEND_TO_ARRAY:
            res,err:=d.AppendToArray(q)
            return res, err
        case cross.REMOVE_ELEMENT_FROM_INCLUDED_ARRAY:
            res,err:=d.RemoveFromIncludedArray(q)
            return res, err
        case cross.TABLE_SIZE:
            res,err := d.BucketSize(q)
            return res, err
        default:
            return nil, cross.IncorrectQueryType
        // err    =  c.Find(bson.M(q.KeyBody)).One(&result)
        // if err != nil {
        //     return result,err
        // }
        }
    return &result_slice, err
}
