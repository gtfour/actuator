package aristo

import "jumper/cross"

var SESSION_IS_VALID  int = 6000
var SESSION_NOT_VALID int = 6001

var session, session_open_error = OpenSession()

type Session struct {
    db cross.Database
}

func OpenSession()(s *Session, err error){
    return
}
