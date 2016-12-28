package boltdb_edge

import "jumper/cross"
import "github.com/boltdb/bolt"

type Storage struct {
    Db                 *bolt.DB
    Error              error
}

func (s *Storage) Close () {
    s.Db.Close()
}


func (s *Storage) Connect ()() {
    s.Db.Close()
}



func GetDatabase(g *cross.Garreth)(s *Storage,err error){

    path:=g.GetPath()
    db, err := bolt.Open(path, 0600, nil)
    if err!= nil { return  nil,cross.CantOpenDatabase }
    s.Db=db
    return s,nil



}


