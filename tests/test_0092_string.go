package main

import "fmt"

type Response struct {
    hostname  string
    text      string
    status    int
}

func(r *Response)String()string{
    return "hostname:"+r.hostname+"\n########\n"+r.text+"\n########\nstatus:"+string(r.status)
}

func main(){
    response:=Response{ "localhost","service pcap_log is-enabled",10 }
    fmt.Printf("%v",response.String())
}

