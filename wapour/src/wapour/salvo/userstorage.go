package salvo

import "errors"
import "wapour/settings"
//import "wapour/api/webclient"
import "github.com/boltdb/bolt"

// dont forget to use SessionId

var UserStorageInstance   =  GetUserStorage()

//var UsersStorage webclient.WengineWrapperStorage
//var UserStorageInstance UserStorage
//var not_exist             =  errors.New("users_doesnot_exist")

var wrapper_exists        =  errors.New("wrapper_exists")
var session_dsnt_exist    =  errors.New("session_dsnt_exist")

type Session struct {
    SessionId    string
    DashboardId  string
    UserId       string
    TokenId      string
}

type WengineWrapper struct {
    username       string
    password       string
    //Url            string
    UserId         string
    TokenId        string
}


type UserStorage interface {
    FindWrapper(user_id string, token_id string)(err error)
    AddWrapper(w WengineWrapper)(err error)
    RemoveWrapper(user_id string, token_id string)
    //
    //
    AddSession(s Session)(err error)
    SetDashboard( session_id string, dashboard_id string )(err error)
    RemoveSession(session_id string)(err error)
    GetSession( session_id string )(s *Session, err error)
}


type RamUserStorage struct {
    StorageType   string
    UserWrappers  MutexUsers
    Sessions      MutexSessions
}

type BoltdbUserStorage   struct {
    StorageType        string
    StorageFileName    string
    DbName             string
    db                 *bolt.DB
    usersTableName     string
    sessionsTableName  string
}

func CreateWengineWrapper(username string, password string )( w WengineWrapper ) {
    w.username = username
    w.password = password
    return w
}


func GetUserStorage()(UserStorage) {
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
            return &user_storage
        }
    } else if settings.ONLINE_USERS_STORAGE_TYPE == "ram" || db_open_failed == true {
        user_storage:=RamUserStorage{StorageType:"ram"}
        user_storage_wrappers := make([]WengineWrapper,0)
        sessions              := make([]Session,0)
        user_storage.UserWrappers.Set(user_storage_wrappers)
        user_storage.Sessions.Set(sessions)
        //UserStorageInstance=user_storage
        return &user_storage
    }
    return nil
}


func (boltdb *BoltdbUserStorage) FindWrapper (user_id string, token_id string) (err error) {
    err = boltdb.db.View(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.usersTableName))
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

func ( ram *RamUserStorage) FindWrapper (user_id string, token_id string) (err error) {

    // have to fix panic: runtime error: index out of range

    /*for w := range (*ram.Wrappers) {
        wrapper:=(*ram.Wrappers)[w]
        if wrapper.UserId == user_id && wrapper.TokenId == token_id {
            return wrapper_exists
        }
    }*/
    wrapper:=ram.UserWrappers.GetItem(user_id, token_id)
    //fmt.Printf("\nram::FindWrapper:%v\n",ram.UserWrappers.Get())
    if wrapper != nil { return wrapper_exists } else { return nil }
}


func ( boltdb *BoltdbUserStorage )AddWrapper( w WengineWrapper )( err error ){
    err=boltdb.db.Update(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.usersTableName))
        //fmt.Printf("\nwrapper token_id: %v session_id: %v\n",w.TokenId,w.SessionId)
        if b==nil{ return nil }
        user,err:=b.CreateBucket([]byte(w.UserId))
        if err==nil{
            err=user.Put([]byte("token_id"),[]byte(w.TokenId))
            if err!=nil{return err}
        } else { return err }
        return nil
    });
    return err
}


func ( ram *RamUserStorage )AddWrapper(w WengineWrapper)(err error) {
    // have to fix panic: runtime error: index out of range
    //(*ram.UserWrappers)=append((*ram.Wrappers),w)
    ram.UserWrappers.AddItem(w)
    //fmt.Printf("Existing Wrappers : %v",ram.UserWrappers.Get())
    return nil
}

func (boltdb *BoltdbUserStorage)AddSession(s Session)(err error) {
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

func ( ram *RamUserStorage )AddSession(s Session)(err error) {

    ram.Sessions.AddItem(s)
    return nil


}

func  (boltdb *BoltdbUserStorage)GetSession( session_id string )(session *Session, err error) {
    err = boltdb.db.View(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.sessionsTableName))
        if b==nil{ return session_dsnt_exist }
        s:=b.Bucket([]byte(session_id))
        if s==nil{ return session_dsnt_exist }
        dashboard_id := s.Get([]byte("dashboard_id"))
        if dashboard_id == nil { dashboard_id=[]byte("")  }
        user_id      := s.Get([]byte("user_id"))
        if user_id == nil { user_id=[]byte("")  }
        token_id     := s.Get([]byte("token_id"))
        if token_id == nil { token_id=[]byte("")  }
        session.SessionId   = session_id
        session.DashboardId = string(dashboard_id)
        session.UserId      = string(user_id)
        session.TokenId     = string(token_id)
        return nil
        });
    if err == nil { return session,nil } else { return nil,err }
}

func ( ram *RamUserStorage )GetSession( session_id string )(*Session,error) {
    session:=ram.Sessions.GetItem(session_id)
    if session == nil { return nil,session_dsnt_exist } else { return session,nil }
}


func ( boltdb *BoltdbUserStorage ) RemoveSession (session_id string)(err error) {
    err = boltdb.db.Update(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.sessionsTableName))
        if b==nil{ return session_dsnt_exist }
        err=b.Delete([]byte(session_id))
        return err
    });
    return err
}

func ( ram *RamUserStorage ) RemoveSession (session_id string)(err error) {
    ram.Sessions.RemoveItem(session_id)
    return nil
}

func ( boltdb *BoltdbUserStorage ) SetDashboard( session_id string, dashboard_id string )(err error) {

    err=boltdb.db.Update(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.sessionsTableName))
        if b==nil{ return session_dsnt_exist }
        session:=b.Bucket([]byte(session_id))
        if session!=nil  {
            err=session.Put([]byte("dashboard_id"),[]byte(dashboard_id))
            if err!=nil{ return err }
        } else { return session_dsnt_exist }
        return nil
    });

    return err


}

func ( ram *RamUserStorage ) SetDashboard ( session_id string, dashboard_id string )(err error) {
    err=ram.Sessions.SetDashboard(session_id,dashboard_id)
    return err
}

func ( boltdb *BoltdbUserStorage )RemoveWrapper(user_id string, token_id string) {

    /*err = boltdb.db.Update(func(tx *bolt.Tx) error {
        b:=tx.Bucket([]byte(boltdb.usersTableName))
        if b==nil{ return session_dsnt_exist }
        err=b.Delete([]byte(session_id))
        return err
    });
    return err
    */

}

func ( ram *RamUserStorage )RemoveWrapper(user_id string, token_id string) {
     ram.UserWrappers.RemoveItem(user_id, token_id)
}
