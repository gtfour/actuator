package userstorage
import "sync"
import "wapour/api/webclient"


type MutexUsers struct {
    sync.RWMutex
    wrappers         []webclient.WengineWrapper
}

type MutexSessions struct {
    sync.RWMutex
    sessions         []webclient.Session
}

func (u *MutexUsers) Set(value []webclient.WengineWrapper) {
    u.Lock()
    defer u.Unlock()
    u.wrappers = value
}

func (u *MutexUsers) Get()([]webclient.WengineWrapper) {
    return u.wrappers
}

func (u *MutexUsers) GetItem(user_id string, token_id string)(value *webclient.WengineWrapper) {
    u.Lock()
    defer u.Unlock()
    for i := range u.wrappers {
        wrapper := &u.wrappers[i]
        if wrapper.UserId == user_id && wrapper.TokenId == token_id {
            return wrapper
        }
    }
    return nil
}

func (u *MutexUsers) AddItem(value webclient.WengineWrapper) {
    u.Lock()
    defer u.Unlock()
    u.wrappers = append(u.wrappers, value)
}

func (u *MutexUsers) RemoveItem(user_id string, token_id string) {
    u.Lock()
    defer u.Unlock()
    for i:= range u.wrappers {
        wrapper := &u.wrappers[i]
        if wrapper.UserId == user_id && wrapper.TokenId == token_id {
            u.wrappers = append(u.wrappers[:0], u.wrappers[0+i:]...)
            break
        }
    }
}

//
//
//

func (s *MutexSessions) Set(value []webclient.Session) {
    s.sessions = value
}

func (s *MutexSessions) Get()([]webclient.Session) {
    return s.sessions
}

func (s *MutexSessions) GetItem(session_id string)(value *webclient.Session) {
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

func (s *MutexSessions) AddItem(value webclient.Session) {
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
            s.sessions = append(s.sessions[:0], s.sessions[0+i:]...)
            break
        }
    }
}
