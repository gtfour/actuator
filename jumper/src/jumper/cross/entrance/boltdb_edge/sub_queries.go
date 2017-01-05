package boltdb_edge

import "encoding/json"
import "github.com/boltdb/bolt"
import "jumper/cross"
import "jumper/common/maps"

func(d *Database)RunQueryCreateNew(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    _,_,err=q.ValidateBodies()
    if err!=nil{ return }
    err=q.CheckTableName()
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
    return
}

func(d *Database)RunQueryUpdate(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    match_by_key,match_by_value,err:=q.ValidateBodies()
    if err!=nil{return}
    return
}

func(d *Database)RunQueryInsert(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    match_by_key,match_by_value,err:=q.ValidateBodies()
    if err!=nil{return}
    return
}

func(d *Database)RunQueryGet(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){

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
    match_by_key,match_by_value,err=q.ValidateBodies()
    if err!=nil{return}

    err = d.db.View(func(tx *bolt.Tx) error {

        table := tx.Bucket([]byte(q.Table))
        if table==nil { return cross.TableDoesntExist }
        err=table.ForEach(func(key, value []byte)(error){

            var key_satisfied    bool = false
            var value_satisfied  bool = false

            search_result_slice       := make(map[string] interface{}, 0)
            key_map                   := make(map[string]interface{},  0)
            query_map                 := make(map[string]interface{},  0)

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

func(d *Database)RunQueryGetAll(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    return
}

func(d *Database)RunQueryRemove(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    return
}



