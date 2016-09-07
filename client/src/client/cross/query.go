package cross

// i also have to implement simple query where type of  Key or/and Query is string 

//import "fmt"
import "encoding/json"
import "github.com/boltdb/bolt"
import "client/common/types"

var TRUE_SLICE  int =  9001
var FALSE_SLICE int =  9000
var BUNT_SLICE  int =  9002
var EMPTY_SLICE int =  9004

var SATISFIED   int =  9006
var UNSATISFIED int =  9008

type Query struct {
    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}
}

func (s *Storage)RunQuery(q Query)(result_slice_addr *[]map[string]interface{}, err error){


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
    } else if q.Type == types.CREATE_NEW {
        if q.QueryBody != nil && q.KeyBody != nil {
            //err         = c.Insert(q.QueryBody)
            key_byte,err_key     := json.Marshal(q.KeyBody)
            query_byte,err_query := json.Marshal(q.QueryBody)
            if err_key!=nil || err_query!=nil {
                return nil, encode_error
            }
            //
            err=s.Db.Update(func(tx *bolt.Tx) error {
                table := tx.Bucket([]byte(q.Table))
                if table==nil { return table_doesnt_exist }

                entry := table.Get(key_byte)
                if entry == nil {
                    err:=table.Put(key_byte, query_byte)
                    return err
                } else {
                    return entry_already_exist
                }

                err:=table.Put(key_byte, query_byte)
                return err
            })
            return nil, err
        } else {
            return nil, empty_query
        }

    } else if q.Type == types.CREATE_NEW_IFNOT {


    } else if q.Type == types.REPLACE {
        if q.QueryBody != nil && q.KeyBody != nil {
            key_byte,err_key     := json.Marshal(q.KeyBody)
            query_byte,err_query := json.Marshal(q.QueryBody)
            if err_key!=nil || err_query!=nil {
                return nil, encode_error
            }
            err=s.Db.Update(func(tx *bolt.Tx) error {
                table := tx.Bucket([]byte(q.Table))
                if table==nil { return table_doesnt_exist }
                entry := table.Get(key_byte)
                if entry == nil {
                    return entry_doesnt_exist
                } else {
                    err:=table.Put(key_byte, query_byte)
                    return err
                }
            })
        }
    } else if q.Type == types.UPDATE || q.Type == types.EDIT   {
        if q.KeyBody   == nil { return nil, empty_key   }
        if q.QueryBody == nil { return nil, empty_query }
        //
        //err = c.Update(bson.M(q.KeyBody), bson.M{"$set": bson.M(q.QueryBody)})
        return nil, err
        //
    } else if q.Type == types.GET || q.Type == types.GET_ALL  || q.Type == types.CHECK_EXIST {
        result_slice_addr,err := s.RunQueryGet(q)
        return result_slice_addr, err
    } else  if q.Type == types.REMOVE {
        if q.KeyBody != nil {
            //err    =  c.Remove(bson.M(q.KeyBody))
            return nil, err
        } else {
            return nil, empty_key
        }
    } else if q.Type == types.INSERT_ITEM || q.Type == types.REMOVE_ITEM {
        if q.KeyBody   == nil { return nil,empty_key   }
        if q.QueryBody == nil { return nil,empty_query }

        if q.Type      == types.INSERT_ITEM {
            //err            =  c.Update(bson.M(q.KeyBody), bson.M{"$push":bson.M(q.QueryBody)})
            return nil,err
        } else {
            //err            =  c.Update(bson.M(q.KeyBody), bson.M{"$pull":bson.M(q.QueryBody)})
            return nil,err
        }
    } else {
        return nil, incorrect_query_type
    }
    // err    =  c.Find(bson.M(q.KeyBody)).One(&result)
    // if err != nil {
    //     return result,err
    // }
    return &result_slice, err
}

func(s *Storage)RunQueryGet(q Query)(result_slice_addr *[]map[string]interface{}, err error) {

    // run queries with map types of KeyBody or/and QueryBody

    result_slice := make([]map[string]interface{},0)

    var match_by_key    bool
    var match_by_value  bool

    if q.KeyBody != nil || q.QueryBody == nil {
        match_by_key   = true
    }
    if q.KeyBody == nil || q.QueryBody != nil {
        match_by_value = true
    }
    if match_by_key == false && match_by_value == false {
        return nil, key_and_value_empty
    }

    err = s.Db.View(func(tx *bolt.Tx) error {

        table := tx.Bucket([]byte(q.Table))
        if table==nil { return table_doesnt_exist }
        err=table.ForEach(func(key, value []byte)(error){

            var key_satisfied    bool = false
            var value_satisfied  bool = false

            search_result_slice       := make(map[string] interface{}, 0)
            key_map                   := make(map[string]interface{},  0)
            query_map                 := make(map[string]interface{},  0)

            err_key                   := json.Unmarshal(key,   &key_map   )
            err_value                 := json.Unmarshal(value, &query_map )

            if err_key != nil || err_value != nil {
                return encode_error
            }
            if match_by_key {

                key_satisfied   = CompareMap(q.KeyBody, key_map)

            }
            if match_by_value {

                value_satisfied = CompareMap(q.QueryBody, query_map)

            }
            if (match_by_key && match_by_value == false && key_satisfied) || (match_by_value && match_by_key == false && value_satisfied) || (match_by_key && match_by_value && key_satisfied && value_satisfied ) {
                search_result_slice["key"]   = key_map
                search_result_slice["value"] = query_map
                result_slice                 = append(result_slice, search_result_slice)
            }

            return nil
         })
         return err
    })
    return &result_slice, err
}

func CheckBoolSlice(slice []bool)(slice_type int){
    true_values  := make([]bool,0)
    false_values := make([]bool,0)
    for i:= range slice {
        value := slice[i]
        if value {
            true_values=append(true_values,value)
        } else {
            false_values=append(false_values,value)
        }
    }
    if len(true_values)>0 && len(false_values)>0 {
        return BUNT_SLICE
    } else if len(true_values)>0 && len(false_values)==0 {
        return TRUE_SLICE
    } else if len(true_values)==0 && len(false_values)>0 {
        return FALSE_SLICE
    } else {
        return EMPTY_SLICE
    }
}

func CompareMap(query map[string]interface{}, dest map[string]interface{})(bool) {
    matching := make([]bool,0)
    for key,value := range query {
        if dest_value,ok := dest[key]; ok == true {
            if dest_value == value {
                matching = append(matching, true)
            } else {
                matching = append(matching, false)
            }
        }
    }
    if CheckBoolSlice(matching) == TRUE_SLICE {
        return true
    } else {
        return false
    }
}
