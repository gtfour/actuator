package fila

import "fmt"
import "time"
import "math/rand"

func GenData(path string)(){

    rand.Seed( time.Now().UTC().UnixNano())
    digit:=rand.Intn(256)
    fmt.Printf("\n%v\n",digit)
}
