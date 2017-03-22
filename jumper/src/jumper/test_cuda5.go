package main

import "fmt"
import "jumper/cuda"

func main(){

    d1:=cuda.Dynima{}
    d2:=cuda.NewDynima()

    fmt.Printf("\nDynima1:%v\nDynima2:%v\n",d1,d2)


}
