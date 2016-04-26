package salvo
import "sync"
import "fmt"
//import "wapour/api/webclient"


type MutexUsers struct {
    sync.RWMutex
    wrappers         []WengineWrapper
}

type MutexSessions struct {
    sync.RWMutex
    sessions         []Session
}

func (u *MutexUsers) Set(value []WengineWrapper) {
    u.Lock()
    defer u.Unlock()
    u.wrappers = value
}

func (u *MutexUsers) Get()([]WengineWrapper) {
    return u.wrappers
}

func (u *MutexUsers) GetItem(user_id string, token_id string)(value *WengineWrapper) {
    u.Lock()
    defer u.Unlock()
    fmt.Printf("\nGetItem::u.wrappers %v\n", u.wrappers)
    for i := range u.wrappers {
        wrapper := &u.wrappers[i]
        if wrapper.UserId == user_id && wrapper.TokenId == token_id {
            return wrapper
        }
    }
    return nil
}

func (u *MutexUsers) AddItem(value WengineWrapper) {
    u.Lock()
    defer u.Unlock()
    u.wrappers = append(u.wrappers, value)
}

func (u *MutexUsers) UpdateItem(item WengineWrapper)() {
    u.Lock()
    defer u.Unlock()
    for i:= range u.wrappers {
        wrapper:=u.wrappers[i]
        if wrapper.UserId == item.UserId && wrapper.TokenId == item.TokenId {
            u.wrappers[i] = item
            break
        }
    }
}


func (u *MutexUsers) RemoveItem(user_id string, token_id string) {
    u.Lock()
    defer u.Unlock()
    for i:= range u.wrappers {
        wrapper := &u.wrappers[i]
        if wrapper.UserId == user_id && wrapper.TokenId == token_id {
            u.wrappers = append(u.wrappers[:i], u.wrappers[1+i:]...)
            fmt.Printf("\nRemoveItem user_id %v token_id %v  i %v   :: u.wrappers  %v\n",user_id ,token_id ,i,u.wrappers)
            break
        }
    }
}

//
//
//

func (s *MutexSessions) Set(value []Session) {
    s.sessions = value
}

func (s *MutexSessions) Get()([]Session) {
    return s.sessions
}

func (s *MutexSessions) GetItem(session_id string)(value *Session) {
    s.Lock()
    defer s.Unlock()
    for i := range s.sessions {
        session := &s.sessions[i]
        if session.SessionId == session_id  {
            return session
        }
    }
    return nil
}

func (s *MutexSessions) SetDashboard(session_id string, dashboard_id string)(error) {
    s.Lock()
    defer s.Unlock()
    for i := range s.sessions {
        session := s.sessions[i]
        if session.SessionId == session_id  {
            session.DashboardId = dashboard_id
            return nil
        }
    }
    return session_dsnt_exist
}


func (s *MutexSessions) AddItem(value Session) {
    s.Lock()
    defer s.Unlock()
    s.sessions = append(s.sessions, value)
}

func (s *MutexSessions) RemoveItem(session_id string) {
    s.Lock()
    defer s.Unlock()
    for i := range s.sessions {
        session := &s.sessions[i]
        if session.SessionId == session_id  {
            s.sessions = append(s.sessions[:i], s.sessions[i+1:]...)
            break
        }
    }
}


