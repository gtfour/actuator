package boltdb_edge

import "errors"
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
        result_slice_single:=make([]map[string]interface{},1)
        //fmt.Printf("\nGet result slice %v\n",result_slice)
        result_slice_single=append(result_slice_single, result_slice[0])
        return &result_slice_single,err
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
