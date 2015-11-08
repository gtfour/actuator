package main

import "math/rand"
import "fmt"
import "time"


func main() {

    rand.Seed( time.Now().UTC().UnixNano())
    rand_int := rand.Int31n(1)
    fmt.Printf("\nrand digit %d\n",rand_int)



}
