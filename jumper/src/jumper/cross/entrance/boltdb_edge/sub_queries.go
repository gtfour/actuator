package boltdb_edge

import "errors"
// import "strconv"
import "encoding/json"
import "github.com/boltdb/bolt"
import "jumper/cross"
import "jumper/common/maps"

func(d *Database)CreateNew(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    _,_,err=q.Validate()
    if err!=nil{ return }
    //
    key_byte,err_key     := json.Marshal(q.KeyBody)
    query_byte,err_query := json.Marshal(q.QueryBody)
    if err_key!=nil || err_query!=nil {
        return nil, cross.EncodeError
    }
    //
    err=d.db.Update(func(tx *bolt.Tx) error {
        table := tx.Bucket([]byte(q.Table))
        if table==nil { return cross.TableDoesntExist }
        // 
        multi   := q.Multi
        queries := q.Queries
        if multi && len(queries)>0  {

        }

        //
        entry := table.Get(key_byte)
        if entry == nil {
            if q.Type == cross.CREATE_NEW_IFNOT {
                err:=table.Put(key_byte, query_byte)
                return err
            } else if q.Type == cross.UPDATE || q.Type == cross.EDIT {
                return cross.EntryDoesntExist
            }
            return err
        } else {
            if q.Type == cross.CREATE_NEW_IFNOT {
                return cross.EntryAlreadyExist
            } else if q.Type == cross.UPDATE || q.Type == cross.EDIT {
                err:=table.Put(key_byte, query_byte)
                return err
            }
        }
        return err
    })
    return nil,err
}

func(d *Database)Update(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    //
    match_by_key,match_by_value,err:=q.Validate()
    _,_=match_by_key,match_by_value
    if err!=nil{return}
    //
    return
}

func(d *Database)QueryInsert(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    //
    match_by_key,match_by_value,err:=q.Validate()
    _,_=match_by_key,match_by_value
    if err!=nil{return}
    return
}

func(d *Database)Get(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){

    // run queries with map types of KeyBody or/and QueryBody
    result_slice := make([]map[string]interface{},0)
    var match_by_key    bool
    var match_by_value  bool

    /*if q.KeyBody != nil || q.QueryBody == nil {
        match_by_key   = true
    }
    if q.KeyBody == nil || q.QueryBody != nil {
        match_by_value = true
    }
    if match_by_key == false && match_by_value == false {
        return nil, cross.KeyAndValueEmpty
    }*/
    match_by_key,match_by_value,err=q.Validate()
    if err!=nil{return}

    err = d.db.View(func(tx *bolt.Tx) error {

        table := tx.Bucket([]byte(q.Table))
        if table==nil { return cross.TableDoesntExist }
        err=table.ForEach(func(key, value []byte)(error){

            var key_satisfied    bool = false
            var value_satisfied  bool = false

            search_result_slice       := make(map[string]interface{}, 0)
            key_map                   := make(map[string]interface{}, 0)
            query_map                 := make(map[string]interface{}, 0)

            err_key                   := json.Unmarshal(key,   &key_map   )
            err_value                 := json.Unmarshal(value, &query_map )

            if err_key != nil || err_value != nil {
                return cross.EncodeError
            }
            if match_by_key {
                key_satisfied   = maps.CompareMap(q.KeyBody, key_map)
            }
            if match_by_value {
                value_satisfied = maps.CompareMap(q.QueryBody, query_map)
		}
            if (match_by_key && (match_by_value == false) && key_satisfied) || (match_by_value && (match_by_key == false) && value_satisfied) || (match_by_key && match_by_value && key_satisfied && value_satisfied ) {
                search_result_slice["key"]   = key_map
                search_result_slice["value"] = query_map
                result_slice                 = append(result_slice, search_result_slice)
                if q.Type == cross.GET || q.Type == cross.CHECK_EXIST {
                    return nil
                }
            }
            return nil
         })
         return err
    })
    if len(result_slice) == 0 {
        result_slice=nil
        err = cross.EntryDoesntExist
    }
    if q.Type == cross.CHECK_EXIST {
        result_slice=nil
    }
    if q.Type == cross.GET {
        if len(result_slice)>1 {
            result_slice_single:=make([]map[string]interface{},1)
            result_slice_single=append(result_slice_single, result_slice[0])
            return &result_slice_single,err
        }
    }
    return &result_slice, err
    //return
}

func(d *Database)GetAll(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    return
}

func(d *Database)Remove(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    match_by_key,match_by_value,err := q.Validate()
    if err!=nil              { return }
    //if match_by_key == false { return nil, errors.New("cross:can't determinate what data should be removed because key is nil ")  }

    key_byte,err_key     := json.Marshal(q.KeyBody)
    //query_byte,err_query := json.Marshal(q.QueryBody)
    if err_key!=nil {
        return nil, cross.EncodeError
    }
    err = d.db.Update(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table==nil { return cross.TableDoesntExist }
        if match_by_key && !match_by_value {
            err=table.Delete(key_byte)
            return err
        } else {
            err=table.ForEach(func(key, value []byte)(error){
                var value_satisfied  bool =   false
                //key_map                   :=  make(map[string]interface{}, 0)
                query_map                 :=  make(map[string]interface{}, 0)

                //err_key                   :=  json.Unmarshal(key,   &key_map   )
                err_value                 :=  json.Unmarshal(value, &query_map )
                if err_value != nil {   //|| err_key != nil {
                    return cross.EncodeError
                }
                value_satisfied = maps.CompareMap(q.QueryBody, query_map)
                if value_satisfied {
                    err=table.Delete(key)
                    return err
                }
                return nil

            });
            return err
        }
    });
    return nil, err
}

func(d *Database)CreateNewTableIfDoesntExist(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    err = q.CheckTableName()
    if err!= nil && len(q.TableList) == 0 {
        return
    }
    if len(q.TableList) == 0 {
        q.TableList=append(q.TableList, q.Table)
    }
    err=d.db.Update(func(tx *bolt.Tx) error {
        glob_state:="cross:\n"
        perfect:=true
        for i:= range q.TableList {
            table_name:=q.TableList[i]
            state:="success\n"
            _,err    := tx.CreateBucketIfNotExists([]byte(table_name))
            if err!=nil { state="failed\n" ; perfect = false  }
            glob_state=glob_state+"\n"+table_name+":"+state
        }
        if !perfect { return errors.New(glob_state) }
        return nil
    });

    return
}

func (d *Database)CheckTableExist(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){

    err = q.CheckTableName()
    if err!= nil {
        return
    }
    err = d.db.View(func(tx *bolt.Tx) error {
        table := tx.Bucket([]byte(q.Table))
        if table==nil { return cross.TableDoesntExist } else { return nil }
    })
    return nil,err
}




func (d *Database)RemoveTable(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    err = q.CheckTableName()
    if err!= nil && len(q.TableList) == 0 {
        return
    }
    if len(q.TableList) == 0 {
        q.TableList=append(q.TableList, q.Table)
    }
    err=d.db.Update(func(tx *bolt.Tx) error {
        glob_state  :="cross:\n"
        perfect     :=true
        for i:= range q.TableList {
            table_name:=q.TableList[i]
            state:="success\n"
            err:=tx.DeleteBucket([]byte(table_name))
            if err!=nil { state="failed\n" ; perfect = false  }
            glob_state=glob_state+"\n"+table_name+":"+state
        }
        if !perfect { return errors.New(glob_state) }
        return err
    });
    return nil, err
}

func (d *Database)ModifyPair(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    key_exist,value_exist,err:=q.Validate()
    if !key_exist   { return nil,cross.KeyIsEmpty   }
    if !value_exist { return nil,cross.ValueIsEmpty }
    if err!=nil     { return nil,err }
    //
    key_byte,err_key     := json.Marshal(q.KeyBody)
    //query_byte,err_query := json.Marshal(q.QueryBody)
    if err_key!=nil {
        return nil, cross.EncodeError
    }
    err=d.db.Update(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table==nil{ return cross.TableDoesntExist  }
        bucket:=table.Bucket(key_byte)
        if bucket==nil{
            entry:=table.Get(key_byte)
            if entry == nil {
                return cross.EntryDoesntExist
            } else {
                entry_map  := make(map[string]interface{}, 0)
                err_entry  := json.Unmarshal(entry, &entry_map )
                if err_entry == nil {
                    updated_map:=make(map[string]interface{}, 0)
                    if q.Type == cross.ADD_PAIR {
                        updated_map,err = maps.UpdateMap(q.QueryBody,entry_map)
                    } else {
                        updated_map,err = maps.RemoveFromMap(q.QueryBody,entry_map)
                    }
                    if err == nil {
                         updated_byte,err_up     := json.Marshal(updated_map)
                         if err_up!=nil {
                             return cross.EncodeError
                         }
                         err=table.Put(key_byte, updated_byte)
                         return err
                    }
                } else {
                    return err_entry
                }
            }
        } else {
            for key,value := range q.QueryBody {
                keyBYTE,   err_key       := json.Marshal(key)
                valueBYTE, err_value     := json.Marshal(value)
                if err_key==nil && err_value==nil {
                    if q.Type == cross.ADD_PAIR {
                        err=bucket.Put(keyBYTE, valueBYTE)
                    } else {
                        //entryBYTE:=bucket.Get(keyBYTE)
                        // expiremental trick. have to test
                        //if string(entryBYTE) == string(valueBYTE) {
                        //    _=bucket.Delete(keyBYTE)
                        //}
                        err=bucket.Delete(keyBYTE)

                    }
                    if err!=nil{return err}
                }
            }
        }
        return nil
    });
    return nil, err
}

func (d *Database)GetPair(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){

    key_exist,value_exist,err:=q.Validate()

    if !key_exist   { return nil, cross.KeyIsEmpty   }
    if !value_exist { return nil, cross.ValueIsEmpty }
    if err!=nil     { return nil, err }

    key_byte,err_key     := json.Marshal(q.KeyBody)
    if err_key!=nil {
        return nil, cross.EncodeError
    }
    err=d.db.View(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table==nil{ return cross.TableDoesntExist  }
        bucket:=table.Bucket(key_byte)
        if bucket==nil{
        } else {
        }
        return err
    });
    return
}


func (d *Database)AppendToArray(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){

    key_exist,value_exist,err := q.Validate()

    if !key_exist   { return nil, cross.KeyIsEmpty   }
    if !value_exist { return nil, cross.ValueIsEmpty }
    if err!=nil     { return nil, err                }
    key_byte,err_key     := json.Marshal( q.KeyBody   )
    query_byte,err_query := json.Marshal( q.QueryBody )
    if err_query != nil || err_key != nil {
        return nil, cross.EncodeError
    }
    err=d.db.Update(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table==nil{ return cross.TableDoesntExist  }
        bucket:=table.Bucket(key_byte)
        if bucket==nil {
             entry := table.Get(key_byte)
             if entry == nil {
                return cross.EntryDoesntExist
            } else {

            }
         } else {
            //
            // when bucket identified by key exists
            //
            // decimal          := 10
            table_stats := table.Stats()
            table_size  := table_stats.KeyN
            new_index   := table_size
            // bucket_size_str  := strconv.FormatInt(int64(bucket_size), decimal)
            key_map          := make(map[string]interface{}, 0)
            key_map["index"] =  new_index
            //new_key_byte     := []byte(key_map)
            new_key_byte,err_tx := json.Marshal(key_map)
            if err_tx != nil {
                return err_tx
            }
            err_tx          = table.Put(new_key_byte, query_byte)
            return err_tx
            //
            //
            //
        }
        return nil // !!!
    });
    return

}

func (d *Database)RemoveFromIncludedArray(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){

    key_exist,value_exist,err:=q.Validate()
    if !key_exist   { return nil, cross.KeyIsEmpty   }
    if !value_exist { return nil, cross.ValueIsEmpty }
    if err!=nil     { return nil, err                }
    key_byte,err_key     := json.Marshal(q.KeyBody)
    if err_key!=nil {
        return nil, cross.EncodeError
    }
    err=d.db.Update(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table==nil{ return cross.TableDoesntExist  }
        bucket:=table.Bucket(key_byte)
        if bucket==nil {
            entry:=table.Get(key_byte)
            if entry == nil {
                return cross.EntryDoesntExist
            } else {

            }
        } else {

        }
        return nil // !!!


    });
    return

}

func (d *Database)BucketSize(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    result_slice := make([]map[string]interface{},0)
    err=d.db.View(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table==nil { return cross.TableDoesntExist }
        stats  := table.Stats()
        size   := stats.KeyN
        result := make(map[string]interface{}, 0)
        result["size"] = size
        result_slice = append(result_slice, result)
        return nil
    });
    return &result_slice, err
}


