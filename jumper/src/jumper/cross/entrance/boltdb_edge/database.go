package boltdb_edge

import "jumper/cross"
import "github.com/boltdb/bolt"

type Storage struct {
    db     *bolt.DB
    err    error
    path   string
}

func (s *Storage) Close () {
    s.db.Close()
}


func (s *Storage) Connect ()(err error) {
    s.db,err=bolt.Open(s.path, 0600, nil)
    return err
}



func GetDatabase(g *cross.Garreth)(s *Storage,err error){
    path:=g.GetPath()
    s.path=path
    return s, nil
}


