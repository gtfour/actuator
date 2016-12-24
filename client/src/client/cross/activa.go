package cross

import "encoding/json"
import "github.com/boltdb/bolt"
import "client/activa"

func WriteMotion(m *activa.Motion)(err error) {

    if Database.Error == nil {
        db:=Database.Db
        err=db.Update(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(Database.motionsTableName))
            if b==nil{ return collection_open_error }
            encoded, err := json.Marshal(m)
            if err!=nil{ return err }
            return b.Put([]byte(m.Id), encoded) //CreateBucket has been replaced to CreateBucketIfNotExists because when err==bolt.ErrBucketExists - dynima is nil
        });
        return err

    } else {
        return unable_to_open_db
    }
}




func GetAllMotions() ( motions  []activa.Motion , err  error ) {

    if Database.Error == nil {
        db     := Database.Db
        motions = make([]activa.Motion,0)
        err = db.View(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(Database.motionsTableName))
            if b==nil{ return collection_open_error }
            err=b.ForEach(func(key, value []byte)(error){
                motion := activa.Motion{}
                err    = json.Unmarshal(value, &motion)
                if err == nil {
                        motions = append(motions, motion)
                }
                return nil
            })
            if len(motions) == 0 {
                return collection_entry_list_is_empty
            }
            return err
        });
        if err == nil {  return motions, err } else {  return nil,err }
    } else {
        return nil, unable_to_open_db
    }
    return nil, err
}
