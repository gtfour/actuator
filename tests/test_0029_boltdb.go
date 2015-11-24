package main

import "log"
import "github.com/boltdb/bolt"

func main() {

    db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

}
