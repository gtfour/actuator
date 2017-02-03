package db

import "jumper/cross"

func CordQuery()(*cross.Query){
    var q cross.Query
    q.Table = CORDS_T
    return &q
}

func MembershipQuery()(*cross.Query){
    var q cross.Query
    q.Table = MEMBERSHIPS_T
    return &q
}
