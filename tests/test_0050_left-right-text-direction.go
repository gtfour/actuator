package main

//import "fmt"

var RIGHT     int = 001
var LEFT      int = 002
var NOT_FOUND int = 999

func GoTillAnyOfSign( since int, direction int, lineAsArray []string, signs []string, breakers ...[]func(string)(bool) ) ( index int, code int ) {


    for i:= range lineAsArray {

        i = since
        if since<0 {i=0}

        if direction==RIGHT {

        } else if direction==LEFT {

        } else {

            return -1, NOT_FOUND

        }

    }

    return index, NOT_FOUND

}

func main(){



}
