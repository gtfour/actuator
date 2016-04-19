package userstorage

//import "fmt"
import "errors"
import "wapour/settings"
import "wapour/api/webclient"
import "github.com/boltdb/bolt"

// dont forget to use SessionId

//var UsersStorage webclient.WengineWrapperStorage
//var UserStorageInstance UserStorage
var not_exist          =  errors.New("users_doesnot_exist")
var wrapper_exists     =  errors.New("wrapper_exists")
var session_dsnt_exist =  errors.New("session_dsnt_exist")

type UserStorage interface {
    FindWrapper(user_id string, token_id string)(err error)
    AddWrapper(w *webclient.WengineWrapper)(err error)
    //
    //
    AddSession(s *webclient.Session)(err error)
    SetDashboard( session_id string, dashboard_id string )(err error)
    RemoveSession(session_id string)(err error)
    GetSession( session_id string )(s *webclient.Session, err error)
}


type RamUserStorage struct {
    StorageType   string
    UserWrappers  *[]*webclient.WengineWrapper
    Sessions      *[]*webclient.Session
}

type BoltdbUserStorage   struct {
    StorageType        string
    StorageFileName    string
    DbName             string
    db                 *bolt.DB
    usersTableName     string
    sessionsTableName  string
}

func CreateUserStorage()(UserStorage) {
    var db_open_failed bool
    if settings.ONLINE_USERS_STORAGE_TYPE        == "db" {
        //UsersStorage.StorageType = "db"
        db_path                  := settings.ONLINE_USERS_DB_FILE
        db, err                  := bolt.Open(db_path, 0600, nil)
        if err!= nil {
            defer db.Close()
            db_open_failed = true
        } else {
            user_storage:=BoltdbUserStorage{StorageType:"db", StorageFileName:db_path, usersTableName:"users", sessionsTableName:"sessions"}
            user_storage.db = db
            //UserStorageInstance=user_storage
            // Remove old bucket
            db.Update(func(tx *bolt.Tx) error {
                _=tx.DeleteBucket([]byte(user_storage.usersTableName))
                _=tx.DeleteBucket([]byte(user_storage.sessionsTableName))
                return nil
            });
            // Create bucket
            db.Update(func(tx *bolt.Tx) error {
                _,err_users    := tx.CreateBucket([]byte(user_storage.usersTableName))
                _,err_sessions := tx.CreateBucket([]byte(user_storage.sessionsTableName))
                if err_users != nil || err_sessions != nil {
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
        sessions              := make([]*webclient.Session,0)
        user_storage.UserWrappers = &user_storage_wrappers
        user_storage.Sessions     = &sessions
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

func ( ram RamUserStorage) FindWrapper (user_id string, token_id string) (err error) {

    // have to fix panic: runtime error: index out of range

    for w := range (*ram.Wrappers) {
        wrapper:=(*ram.Wrappers)[w]
        if wrapper.UserId == user_id && wrapper.TokenId == token_id {
            return wrapper_exists
        }
    }
    return nil
}


func ( boltdb BoltdbUserStorage )AddWrapper( w *webclient.WengineWrapper )( err error ){
    err=boltdb.db.Update(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.tableName))
        //fmt.Printf("\nwrapper token_id: %v session_id: %v\n",w.TokenId,w.SessionId)
        if b==nil{ return nil }
        user,err:=b.CreateBucket([]byte(w.UserId))
        if err==nil{
            err=user.Put([]byte("token_id"),[]byte(w.TokenId))
            if err!=nil{return err}
            err=user.Put([]byte("session_id"),[]byte(w.SessionId))
            if err!=nil{ return err }
        } else { return err }
        return nil
    });
    return err
}


func ( ram RamUserStorage )AddWrapper(w *webclient.WengineWrapper)(err error) {
    // have to fix panic: runtime error: index out of range
    (*ram.Wrappers)=append((*ram.Wrappers),w)
    return nil
}

func (boltdb BoltdbUserStorage)AddSession(s *webclient.Session)(err error) {
    err=boltdb.db.Update(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.sessionsTableName))
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

func  GetSession( session_id string )(s *webclient.Session, err error) {

    var session webclient.Session
    err = boltdb.db.View(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.sessionsTableName))
        if b==nil{ return nil }
        session:=b.Bucket([]byte(session_id))
        if session==nil{ return nil }
        dashboard_id := session.Get([]byte("dashboard_id"))
        if dashboard_id == nil { dashboard_id=""  }
        user_id      := session.Get([]byte("user_id"))
        if user_id == nil { user_id=""  }
        token_id     := session.Get([]byte("token_id"))
        if token_id == nil { token_id=""  }
        //fmt.Printf("token_id:%v  session_id:%v",string(ex_token_id),string(ex_session_id))
        return nil
        });

    return err
}
