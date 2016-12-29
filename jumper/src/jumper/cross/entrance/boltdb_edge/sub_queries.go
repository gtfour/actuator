package boltdb_edge

import "fmt"
import "jumper/cross"

func(d *Database)RunQueryCreateNew(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    return
}

func(d *Database)RunQueryUpdate(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    return
}

func(d *Database)RunQueryInsert(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    return
}

func(d *Database)RunQueryGet(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    return
}

func(d *Database)RunQueryGetAll(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    fmt.Printf("\nRun query: Get all :\n")
    return
}

func(d *Database)RunQueryRemove(q *cross.Query)(result_slice_addr *[]map[string]interface{}, err error){
    return
}



