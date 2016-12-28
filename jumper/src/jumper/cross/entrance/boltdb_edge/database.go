package boltdb_edge

import "github.com/boltdb/bolt"

type Storage struct {
    Db                 *bolt.DB
    Error              error
}

func (s *Storage) Close () {
    s.Db.Close()
}


