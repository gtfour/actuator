package main

import "fmt"
import "wengine/activa"
import "wengine/dusk"
import "wengine/core/common"

func main() {

    d         := dusk.OpenDatabase("mongo","wengine","OpenStack123","127.0.0.1","wengine")
    time_now  := common.GetTime()
    my_motion := activa.Motion{Id:time_now, InitTime:time_now}
    err       := d.WriteMotion(&my_motion)
    fmt.Printf("Insert err:%v",err)





}
