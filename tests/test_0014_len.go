package main

import "fmt"

type Test1 struct {


    Hello []*chan bool


}

type Test2 struct {


    Hello []chan bool


}

func main(){

test1:=Test2{}
//test2:=&Test1{}

fmt.Printf("%d",len(test1.Hello))
if (len(test1.Hello)>0){


    fmt.Println("wtf")

}





}
