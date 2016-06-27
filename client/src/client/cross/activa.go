package cross

import "client/activa"
import "encoding/json"
import "github.com/boltdb/bolt"


func WriteMotion(m *activa.Motion)(err error) {

    if STORAGE_INSTANCE.Error == false {
        db:=STORAGE_INSTANCE.Db
        err=db.Update(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.motionsTableName))
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
