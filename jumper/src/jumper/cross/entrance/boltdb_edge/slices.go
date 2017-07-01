package boltdb_edge
/*

import "encoding/json"
import "github.com/boltdb/bolt"
import "jumper/cross"

func (d *Database)AppendToSlice(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
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
    query_byte,err_query := json.Marshal( q.QueryBody )
    if err_query != nil {
        return nil, cross.EncodeError
    }
    //
    // get entry key and slice name
    //
    entryId,entryIdOk     := q.KeyBody["entry_id"]
    newKey, newKeyOk      := q.KeyBody["new_key"]
    sliceName,sliceNameOk := q.KeyBody["slice_name"]
    //
    // 
    //
    if !entryIdOk   { return nil, cross.EntryIdIsEmpty   }
    if !sliceNameOk { return nil, cross.SliceNameIsEmpty }
    //
    //
    //
    err=d.db.Update(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table == nil { return cross.TableDoesntExist  }
        //
        entryIdStr  := fmt.Sprintf( "%v", entryId )
        entryIdByte := []byte(entryIdStr)
        // 
        bucket:=table.Bucket(entryIdByte)
        //
        if bucket == nil {
            //
            //
            entry := table.Get(entryIdByte)
            if entry == nil {
                //
                return cross.EntryDoesntExist
            } else {
                //
            }
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
            sliceBucket    := bucket.Bucket(sliceNameByte)
            if sliceBucket == nil {
                var errCreate error
                sliceBucket,errCreate = bucket.CreateBucketIfNotExists(sliceNameByte)
                if errCreate != nil { return errCreate }
            }
            //
            //
            stats     := sliceBucket.Stats()
            size      := stats.KeyN
            new_index := size
            //
            //
            key_map   := make(map[string]interface{}, 0)
            key_word  := "index"
            //
            //
            if newKeyOk {
                //
                key_word          := "key"
                newKeyStr         := fmt.Sprintf("%v", newKey)
                // newKeyByte     := []byte(newKeyStr)
                key_map[key_word] =  newKeyStr
                //
            } else {
                //
                key_map[key_word] =  new_index
                //
            }
            //
            // bucket_size_str  := strconv.FormatInt(int64(bucket_size), decimal)
            //
            // key_map[key_word] =  new_index
            //
            // new_key_byte     := []byte(key_map)
            //
            new_key_byte,err_tx := json.Marshal(key_map)
            if err_tx != nil {
                return err_tx
            }
            //
            err_tx = table.Put(new_key_byte, query_byte)
            return err_tx
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
    key_byte,err_key     := json.Marshal(q.KeyBody)
    if err_key!=nil {
        return nil, cross.EncodeError
    }
    err=d.db.Update(func(tx *bolt.Tx) error {
        table:=tx.Bucket([]byte(q.Table))
        if table==nil{ return cross.TableDoesntExist }
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

func(d *Database)GetSlice(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
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
        entryIdStr  := fmt.Sprintf( "%v", entryId )
        entryIdByte := []byte(entryIdStr)
        bucket      := table.Bucket(entryIdByte)
        if bucket == nil {
            //
            return cross.EntryDoesntExist
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
*/
