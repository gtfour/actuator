package userstorage

import "errors"
import "wapour/settings"
import "wapour/api/webclient"
import "github.com/boltdb/bolt"

//var UsersStorage webclient.WengineWrapperStorage
var UserStorageInstance UserStorage
var not_exist =  errors.New("users_doesnot_exist")

type UserStorage interface {
    FindWrapper(user_id string, token_id string, session_id string)(err error)
    AddWrapper(w *webclient.WengineWrapper)(err error)
}


type FileUserStorage struct {
    StorageType     string
    Wrappers        *[]*webclient.WengineWrapper
}

type BoltdbUserStorage   struct {
   StorageType      string
   StorageFileName  string
   DbName           string
   db               *bolt.DB
   tableName        string 
}

func init() {
    var db_open_failed bool
    if settings.ONLINE_USERS_STORAGE_TYPE        == "db" {
        //UsersStorage.StorageType = "db"
        db_path                  := settings.ONLINE_USERS_DB_FILE
        db, err                  := bolt.Open(db_path, 0600, nil)
        if err!= nil {
            defer db.Close()
            db_open_failed = true
            return
        } else {
            user_storage:=BoltdbUserStorage{StorageType:"db", StorageFileName:db_path, tableName:"users"}
            user_storage.db = db
            UserStorageInstance=user_storage
            // Remove old bucket
            db.Update(func(tx *bolt.Tx) error {
                _=tx.DeleteBucket([]byte(user_storage.tableName))
                return nil
            });
            // Create bucket
            db.Update(func(tx *bolt.Tx) error {
                _,err:=tx.CreateBucket([]byte(user_storage.tableName))
                if err != nil {
                    db_open_failed = true
                    defer db.Close()
                    return err
                }
                return nil
            });
            return
        }
    } else if settings.ONLINE_USERS_STORAGE_TYPE == "ram" || db_open_failed == true {
        user_storage:=FileUserStorage{StorageType:"ram"}
        UserStorageInstance=user_storage
        return
    }
}

func main() {

}

func (boltdb BoltdbUserStorage) FindWrapper (user_id string, token_id string, session_id string) (err error) {
    err = boltdb.db.View(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.tableName))
        if b!=nil{ return not_exist }
        

        return nil
    });


    return err
}

func (boltdb BoltdbUserStorage)AddWrapper(w *webclient.WengineWrapper)(err error) {
    db.Update(func(tx *bolt.Tx) error {



    });
    return nil
}

func (boltdb FileUserStorage) FindWrapper (user_id string, token_id string, session_id string) (err error) {


    return err
}

func (boltdb FileUserStorage)AddWrapper(w *webclient.WengineWrapper)(err error) {
    return nil
}

