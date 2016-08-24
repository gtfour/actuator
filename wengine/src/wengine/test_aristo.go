package main

import "wengine/dusk"

var database dusk.Database = dusk.DATABASE_INSTANCE

func main() {

    database.CheckAccess("11","12","41","42")

}
