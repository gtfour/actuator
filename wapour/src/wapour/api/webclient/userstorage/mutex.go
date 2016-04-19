package userstorage
import "sync"
import "wapour/api/webclient"

type MutexSessions struct {
    sync.Mutex
    Sessions         []webclient.Session
}

type MutexUsers struct {
    sync.Mutex
    Wrappers         []webclient.WengineWrapper
}
