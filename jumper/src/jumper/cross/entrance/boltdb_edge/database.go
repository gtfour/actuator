package boltdb_edge

import "jumper/cross"
import "github.com/boltdb/bolt"

type Database struct {
    db     *bolt.DB
    err    error
    path   string
}

func (d *Database) Close () {
    d.db.Close()
}


func (d *Database) Connect ()(err error) {
    d.db,err=bolt.Open(d.path, 0600, nil)
    return err
}



func GetDatabase(g *cross.Garreth)(d *Database,err error){
    path:=g.GetPath()
    d.path=path
    return d, nil
}


