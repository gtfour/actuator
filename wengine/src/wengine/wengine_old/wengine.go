//server side
package main

import (
    "fmt"
    "net"
    "net/http"
    "io/ioutil"
//    "reflect"
//    "errors"
//    "fmt"
    "encoding/json"
    "os"
//    "wengine/settings"
)

type TestJson struct {

    //struct to test json reciever
    Data string `json:"data"`

}



func main() {

    port, filepath, err := get_args()
    if err!=nil{

        fmt.Println(err)
        os.Exit(1)

    }
    port = ":"+port

    messages := make(chan string)

    printLogMessage(messages)

    web_page,_:=makeHandlerFunc(filepath,messages,"html_get")
    client_info,_:=makeHandlerFunc(filepath,messages,"client_info")

    http.HandleFunc("/", web_page)
    http.HandleFunc("/clinfo", client_info)

    panic(http.ListenAndServe(port, nil))
}


func makeHandlerFunc(filepath string,messages chan string, reqtype string) (handle_func func(w http.ResponseWriter, r *http.Request),err error) {

    if reqtype == "html_get" {
 
    handle_func=func(w http.ResponseWriter, r *http.Request) {

    text , _ := readFile(filepath)

    fmt.Fprintf(w, text)

    messages <-(r.Method+" | "+r.Proto+" | "+r.URL.Path)

}
}

    if reqtype ==  "client_info" {

    handle_func=func (rw http.ResponseWriter, r *http.Request) {

    body, err := ioutil.ReadAll(r.Body)
    fmt.Println(r.Body)

    if err != nil {
        panic("can't read request body")
    }

    var t TestJson

    err = json.Unmarshal(body, &t)

    if err != nil {
        panic("can't parse json post request")
    }

    messages <-(r.Method+" | "+r.Proto+" | "+r.URL.Path+"|")
    messages <-(t.Data)


    }
}

    return handle_func,nil

}

func readFile(filename string) (text string ,err error ) {

    rawBytes, err := ioutil.ReadFile(filename)

    text = string(rawBytes)

    return text, nil
}

func render_test_json(json string){



}

func get_args()(port string,filepath string , err error) {


    if len(os.Args) == 3 {

        port = os.Args[1]
        filepath = os.Args[2]
        if inFile, err := os.Open(filepath); err != nil{

            defer inFile.Close()
            return "","",fmt.Errorf("error: Unable to open file %s",filepath)
        }

        } else {
              return "","",fmt.Errorf("usage: <port_number> <path_to_html_file>") 
        }


        ln, err := net.Listen("tcp", ":"+port);
        if err!=nil {

            return "","",fmt.Errorf("error: Unable to open port %s",port)
        }

        defer ln.Close()

        return port,filepath,err



}




func printLogMessage(message_channel chan string) {

    go func() {
    for {
    message:= <-message_channel
    fmt.Printf("| %s |\n",message)

    }
   }()

}


