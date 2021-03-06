package boltdb_edge

import "fmt"
import "errors"
// import "strconv"
import "encoding/json"
import "github.com/boltdb/bolt"
import "jumper/cross"
import "jumper/common/maps"
//
import "jumper/common/flexi"

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


func (d *Database)AppendToSlice(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    //
    //
    key_exist,value_exist,err := q.Validate()
    //
    // query.KeyBody should contains following key fields: 
    // 
    // "entry_id"   : is entry's uniq identifier
    // "slice_name" : slice where new value will be appended( it  seems that just string type of value is allowed )
    //
    if !key_exist   { return nil, cross.KeyIsEmpty   }
    if !value_exist { return nil, cross.ValueIsEmpty }
    if err!=nil     { return nil, err                }
    //
    // key_byte,err_key  := json.Marshal( q.KeyBody )
    //
    // query_byte,err_query := json.Marshal( q.QueryBody )
    // if err_query != nil {
    //    return nil, cross.EncodeError
    // }
    //
    // get entry key and slice name
    //
    entryId,entryIdOk                 := q.KeyBody["entry_id"]
    // newKey, newKeyOk                  := q.KeyBody["new_key"]    // wtf ??? 
    sliceName,sliceNameOk             := q.KeyBody["slice_name"]
    valueToAppend,valueToAppendExists := q.QueryBody["value"]
    //
    // 
    //
    if !entryIdOk           { return nil, cross.EntryIdIsEmpty         }
    if !sliceNameOk         { return nil, cross.SliceNameIsEmpty       }
    if !valueToAppendExists { return nil, cross.NothingIsAppendToSlice }
    //
    //
    //
    err=d.db.Update(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table == nil { return cross.TableDoesntExist  }
        //
        // ???
        // entryIdStr  := fmt.Sprintf( "%v", entryId )
        // entryIdByte := []byte(entryIdStr)
        // EncodeError 
        // ??? 
        //
        entryIdByte,errMarshal := json.Marshal(entryId)
        if errMarshal != nil { return errMarshal }
        // ???
        fmt.Printf("\nAppendToArray\tEntryIdByte: %v\n", entryIdByte)
        // 
        bucket:=table.Bucket(entryIdByte)
        //
        if bucket == nil {
            //
            //
            // entry := table.Get(entryIdByte)
            // if entry == nil {
                //
            //     return cross.EntryDoesntExist
            // } else {
                //
            //}
            // --
            //
            entry_byte := table.Get(entryIdByte)
            if entry_byte == nil { return cross.EntryDoesntExist }
            entry_map  := make(map[string]interface{}, 0)
            err_entry  := json.Unmarshal(entry_byte, &entry_map)
            if err_entry == nil {
                fmt.Printf("\n::>> decoded map\n%v\n<<::\n", entry_map )
                //
                sliceNameStr     := fmt.Sprintf( "%v", sliceName)
                // valueToAppendStr := fmt.Sprintf( "%v", valueToAppend)
                //
                targetSlice, sliceExists     := entry_map[sliceNameStr]
                if sliceExists {
                    // if slice exists then appending
                    newTargetSlice                  := make([]interface{}, 0)
                    newTargetSlice, errOnAppend := flexi.Append( targetSlice, valueToAppend )
                    if errOnAppend != nil { return errOnAppend }
                    //
                    entry_map[sliceNameStr]          = newTargetSlice
                    newEntryByte , errNewEntryEncode := json.Marshal(entry_map)
                    if errNewEntryEncode == nil {
                        // now we have to overwrite existing entry_map . now it should contains updated map
                        return table.Put(entryIdByte, newEntryByte)
                    } else {
                        return errNewEntryEncode
                    }
                    //
                    return nil
                } else {
                    if q.CreateIfNot == true {
                        //
                        // let's create this slice if it still doesn't exist 
                        //
                        entry_map[sliceNameStr]          =  valueToAppend
                        newEntryByte , errNewEntryEncode := json.Marshal(entry_map)
                        //
                        if errNewEntryEncode == nil {
                            // now we have to overwrite existing entry_map . now it should contains updated map
                            return table.Put(entryIdByte, newEntryByte)
                        } else {
                            return errNewEntryEncode
                        }
                        //
                    } else {
                        return cross.SliceDoesntExist
                    }
                    //
                }
                //
                return nil
            } else  {
                return cross.DecodeError
            }
            //
            // --
            return cross.EntryDoesntExist
            //
            //
        } else {
            //
            // when bucket identified by key exists
            //
            // decimal          := 10
            //
            sliceNameStr   := fmt.Sprintf( "%v", sliceName)
            sliceNameByte  := []byte(sliceNameStr)
            // //sliceBucket    := bucket.Bucket(sliceNameByte)
            sliceByte      := bucket.Get(sliceNameByte)
            if sliceByte == nil  {
                if q.CreateIfNot == true {
                    // var errCreate error
                    newTargetSlice              := make([]interface{}, 0)
                    // sliceBucket, errCreate = bucket.CreateBucketIfNotExists(sliceNameByte)
                    newTargetSlice, errOnAppend := flexi.Append( newTargetSlice, valueToAppend )
                    if errOnAppend != nil { return errOnAppend }
                    sliceValueByte , errSliceEncode := json.Marshal(newTargetSlice)
                    if errSliceEncode == nil {
                        // now we have to overwrite existing entry_map . now it should contains updated map
                        return bucket.Put(sliceNameByte, sliceValueByte)
                    } else {
                        return errSliceEncode
                    }
                } else {
                    return cross.SliceDoesntExist
                }
            }
            //
            //
            // //stats     := sliceBucket.Stats()
            // //size      := stats.KeyN
            // //new_index := size
            //
            //
            // //key_map   := make(map[string]interface{}, 0)
            // //key_word  := "index"
            //
            //
            // //if newKeyOk {
                //
                // //key_word          := "key"
                // //newKeyStr         := fmt.Sprintf("%v", newKey)
                // newKeyByte     := []byte(newKeyStr)
                // //key_map[key_word] =  newKeyStr
                //
            // //} else {
                //
                // //key_map[key_word] =  new_index
                //
            // //}
            //
            // bucket_size_str  := strconv.FormatInt(int64(bucket_size), decimal)
            //
            // key_map[key_word] =  new_index
            //
            // new_key_byte     := []byte(key_map)
            //
            // //new_key_byte,err_tx := json.Marshal(key_map)
            // //if err_tx != nil {
                // //return err_tx
            // //}
            //
            // //err_tx = table.Put(new_key_byte, query_byte) // what ??? ... seems code is wrong
            // //return err_tx
            //
        }
        return nil // !!!
    });
    return

}

func (d *Database)RemoveFromSlice(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){

    key_exist, value_exist, err := q.Validate()
    if !key_exist   { return nil, cross.KeyIsEmpty   }
    if !value_exist { return nil, cross.ValueIsEmpty }
    if err!=nil     { return nil, err                }
    // key_byte,err_key     := json.Marshal(q.KeyBody)
    // if err_key!=nil {
    //    return nil, cross.EncodeError
    //}
    //
    entryId,entryIdOk              := q.KeyBody["entry_id"]
    sliceName,sliceNameOk          := q.KeyBody["slice_name"]
    removeIndex, removeIndexExists := q.QueryBody["index"] // index or list of indexes to remove //
    //
    //
    if !entryIdOk           { return nil, cross.EntryIdIsEmpty     }
    if !sliceNameOk         { return nil, cross.SliceNameIsEmpty   }
    if !removeIndexExists   { return nil, cross.RemoveIndexIsEmpty }
    //
    //
    err=d.db.Update(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table==nil{ return cross.TableDoesntExist }
        //
        entryIdByte,errMarshal := json.Marshal(entryId)
        if errMarshal != nil { return errMarshal }
        //
        bucket:=table.Bucket(entryIdByte)
        if bucket==nil {
            entry_byte:=table.Get(entryIdByte)
            if entry_byte == nil { return cross.EntryDoesntExist } else {
                entry_map  := make(map[string]interface{}, 0)
                err_entry  := json.Unmarshal(entry_byte, &entry_map)
                if err_entry == nil {
                    sliceNameStr             := fmt.Sprintf( "%v", sliceName)
                    targetSlice, sliceExists := entry_map[sliceNameStr]
                    if sliceExists == true {
                        // a = append(a[:i], a[i+1:]...)
                        newTargetSlice,errOnRemove := flexi.Remove(targetSlice, removeIndex)
                        //
                        if errOnRemove != nil { return errOnRemove }
                        //
                        entry_map[sliceNameStr]          =  newTargetSlice
                        newEntryByte , errNewEntryEncode := json.Marshal(entry_map)
                        if errNewEntryEncode == nil {
                            // now we have to overwrite existing entry_map . now it should contains updated map
                            return table.Put(entryIdByte, newEntryByte)
                        } else {
                            return errNewEntryEncode
                        }
                        //

                    } else {
                        return cross.SliceDoesntExist
                    }

                } else {
                    return cross.DecodeError
                }
            }
        } else {
            sliceNameStr  := fmt.Sprintf( "%v", sliceName)
            sliceNameByte := []byte(sliceNameStr)
            sliceByte     := bucket.Get(sliceNameByte)
            if sliceByte != nil {

            } else {
                return cross.SliceDoesntExist
            }

        }
        return nil // !!!
    });
    return

}

func (d *Database)GetSliceElem(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){

    key_exist, value_exist, err := q.Validate()
    if !key_exist   { return nil, cross.KeyIsEmpty   }
    if !value_exist { return nil, cross.ValueIsEmpty }
    if err!=nil     { return nil, err                }
    // key_byte,err_key     := json.Marshal(q.KeyBody)
    // if err_key!=nil {
    //    return nil, cross.EncodeError
    //}
    //
    entryId,entryIdOk              := q.KeyBody["entry_id"]
    sliceName,sliceNameOk          := q.KeyBody["slice_name"]
    removeIndex, removeIndexExists := q.QueryBody["index"] // index or list of indexes to remove //
    sliceNameStr                   := fmt.Sprintf( "%v", sliceName)
    //
    //
    if !entryIdOk           { return nil, cross.EntryIdIsEmpty     }
    if !sliceNameOk         { return nil, cross.SliceNameIsEmpty   }
    if !removeIndexExists   { return nil, cross.RemoveIndexIsEmpty }
    //
    //
    err=d.db.View(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table==nil{ return cross.TableDoesntExist }
        //
        entryIdByte,errMarshal := json.Marshal(entryId)
        if errMarshal != nil { return errMarshal }
        //
        bucket:=table.Bucket(entryIdByte)
        if bucket==nil {
            entry_byte:=table.Get(entryIdByte)
            if entry_byte == nil { return cross.EntryDoesntExist } else {
                entry_map  := make(map[string]interface{}, 0)
                err_entry  := json.Unmarshal(entry_byte, &entry_map)
                if err_entry == nil {
                    targetSlice, sliceExists := entry_map[sliceNameStr]
                    if sliceExists == true {
                        // a = append(a[:i], a[i+1:]...)
                        newTargetSlice,errOnRemove := flexi.Remove(targetSlice, removeIndex)
                        //
                        if errOnRemove != nil { return errOnRemove }
                        //
                        entry_map[sliceNameStr]          =  newTargetSlice
                        newEntryByte , errNewEntryEncode := json.Marshal(entry_map)
                        if errNewEntryEncode == nil {
                            // now we have to overwrite existing entry_map . now it should contains updated map
                            return table.Put(entryIdByte, newEntryByte)
                        } else {
                            return errNewEntryEncode
                        }
                        //

                    } else {
                        return cross.SliceDoesntExist
                    }

                } else {
                    return cross.DecodeError
                }
            }
        } else {
            sliceNameByte := []byte(sliceNameStr)
            sliceByte     := bucket.Get(sliceNameByte)
            if sliceByte != nil {

            } else {
                return cross.SliceDoesntExist
            }

        }
        return nil // !!!
    });
    return

}

func(d *Database)GetSlice(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    //
    result_slice := make([]map[string]interface{},0)
    //
    key_exist,_,err := q.Validate()
    if !key_exist { return nil, cross.KeyIsEmpty }
    // get entry key and slice name
    entryId, entryIdOk     := q.KeyBody["entry_id"]
    sliceName, sliceNameOk := q.KeyBody["slice_name"]
    //
    if !entryIdOk   { return nil, cross.EntryIdIsEmpty   }
    if !sliceNameOk { return nil, cross.SliceNameIsEmpty }
    //
    err = d.db.View( func(tx *bolt.Tx) error {
        //
        table    := tx.Bucket( []byte(q.Table) )
        if table == nil { return cross.TableDoesntExist }
        //
        // entryIdStr  := fmt.Sprintf( "%v", entryId )
        //
        // entryIdByte := []byte(entryIdStr)
        //
        entryIdByte,errMarshal := json.Marshal( entryId )
        if errMarshal != nil { return errMarshal }
        //   fmt.Printf("\n GetSlice\tEntryIdByte: %v\n", entryIdByte)
        //
        bucket := table.Bucket(entryIdByte)
        if bucket == nil {
            //
            // try to get just entry instead of bucket
            //
            entry_byte := table.Get(entryIdByte)
            if entry_byte != nil {
                //
                entry_map  := make(map[string]interface{}, 0)
                err_entry  := json.Unmarshal(entry_byte, &entry_map)
                if err_entry == nil {
                    fmt.Printf("\n::>> decoded map\n%v\n<<::\n", entry_map )
                    //
                    sliceNameStr   := fmt.Sprintf( "%v", sliceName)
                    //
                    targetSlice, sliceExists     := entry_map[sliceNameStr]
                    if sliceExists {
                        // //
                        // //
                        // //
                        search_result_slice          := make(map[string]interface{}, 0)
                        search_result_slice["value"] =  targetSlice
                        result_slice                 =  append(result_slice, search_result_slice)
                        return nil
                        // //
                        // //
                        // //
                    } else {
                        // //
                        // //
                        return cross.SliceDoesntExist
                        // //
                        // //
                    }
                    //
                    return nil
                } else  {
                    return cross.DecodeError
                }
                //
            } else {
                return cross.EntryDoesntExist
            }
            //
            //
            //
            // return cross.EntryDoesntExist
            //
            //
            //
        } else {
            //
            sliceNameStr   := fmt.Sprintf( "%v", sliceName)
            sliceNameByte  := []byte(sliceNameStr)
            sliceBucket    := bucket.Bucket(sliceNameByte)
            if sliceBucket == nil { return cross.SliceDoesntExist }
            //
            sliceBucket.ForEach(func(k, v []byte) error {
		fmt.Printf("key=%s, value=%s\n", k, v)
		return nil
	    })
            return nil
            //
        }
        //
    });
    //
    return &result_slice, err
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
