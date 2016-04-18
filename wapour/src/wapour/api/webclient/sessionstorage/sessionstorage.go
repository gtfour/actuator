package sessionstorage

//import "fmt"
import "errors"
import "wapour/settings"
import "wapour/api/webclient"
import "github.com/boltdb/bolt"

// dont forget to use SessionId

//var UsersStorage webclient.WengineWrapperStorage
//var UserStorageInstance UserStorage
var not_exist      =  errors.New("users_doesnot_exist")
var wrapper_exists =  errors.New("wrapper_exists")

type SessionStorage interface {
    AddSession(s *webclient.Session)                            (err error)
    SetDashboard( session_id string, dashboard_id string )      (err error)
    RemoveSession(session_id string)                            (err error)
    GetSession( session_id string )                             (s *webclient.Session, err error)
}


type RamSessionStorage struct {
    StorageType     string
    Wrappers        *[]*webclient.WengineWrapper
}

type BoltdbSessionStorage   struct {
   StorageType      string
   StorageFileName  string
   DbName           string
   db               *bolt.DB
   tableName        string
}

func CreateUserStorage()(SessionStorage) {
    var db_open_failed bool
    if settings.ONLINE_USERS_STORAGE_TYPE        == "db" {
        //UsersStorage.StorageType = "db"
        db_path                  := settings.ONLINE_USERS_DB_FILE
        db, err                  := bolt.Open(db_path, 0600, nil)
        if err!= nil {
            defer db.Close()
            db_open_failed = true
        } else {
            user_storage:=BoltdbSessionStorage{StorageType:"db", StorageFileName:db_path, tableName:"user_sessions"}
            user_storage.db = db
            //UserStorageInstance=user_storage
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
            return user_storage
        }
    } else if settings.ONLINE_USERS_STORAGE_TYPE == "ram" || db_open_failed == true {
        user_storage:=RamUserStorage{StorageType:"ram"}
        user_storage_wrappers := make([]*webclient.WengineWrapper,0)
        user_storage.Wrappers = &user_storage_wrappers
        //UserStorageInstance=user_storage
        return user_storage
    }
    return nil
}


func (boltdb BoltdbUserStorage) FindWrapper (user_id string, token_id string) (err error) {
    err = boltdb.db.View(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.tableName))
        if b==nil{ return nil }
        user:=b.Bucket([]byte(user_id))
        if user==nil{ return nil }
        ex_token_id   := user.Get([]byte("token_id"))
        if string(ex_token_id) == token_id {
            return wrapper_exists
        }
        //fmt.Printf("token_id:%v  session_id:%v",string(ex_token_id),string(ex_session_id))
        return nil
        });

    return err
}

func (boltdb BoltdbUserStorage)AddSession(s *webclient.Session)(err error) {
    err=boltdb.db.Update(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.tableName))
        if b==nil{ return nil }
        session,err:=b.CreateBucket([]byte(s.SessionId))
        if err==nil || err==bolt.ErrBucketExists { // If the key exist then its previous value will be overwritten
            err=session.Put([]byte("dashboard_id"),[]byte(s.DashboardId))
            if err!=nil{ return err }
            err=session.Put([]byte("user_id"),[]byte(s.UserId))
            if err!=nil{ return err }
            err=session.Put([]byte("token_id"),[]byte(s.TokenId))
            if err!=nil{ return err }
        } else { return err }
        return nil
    });
    return err
}

func ( ram RamUserStorage) FindWrapper (user_id string, token_id string) (err error) {

   for w := range (*ram.Wrappers) {
       wrapper:=(*ram.Wrappers)[w]
       if wrapper.UserId == user_id && wrapper.TokenId == token_id {
            return wrapper_exists
       }
    }
    return nil
}

func ( ram RamUserStorage )AddWrapper(w *webclient.WengineWrapper)(err error) {
    (*ram.Wrappers)=append((*ram.Wrappers),w)
    return nil
}

