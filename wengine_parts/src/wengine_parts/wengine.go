//server side
package main

import (
    "fmt"
    "net"
    "net/http"
    "io/ioutil"
//    "errors"
//    "fmt"
//    "encoding/json"
    "os"
//    "wengine_parts/settings"
)



func main() {

    port, filepath, err := get_args()
    if err!=nil{

        fmt.Println(err)
        os.Exit(1)

    }
    port = ":"+port

    messages := make(chan string)

    printLogMessage(messages)

    handler_func,_:=makeHandlerFunc(filepath,messages,"html_get")

    http.HandleFunc("/", handler_func)

    panic(http.ListenAndServe(port, nil))
}


func makeHandlerFunc(filepath string,messages chan string, reqtype string) (func(w http.ResponseWriter, r *http.Request), error) {

    if reqtype == "html_get" {
   
    return func(w http.ResponseWriter, r *http.Request) {

    text , _ := readFile(filepath)

    fmt.Fprintf(w, text)

    messages <-(r.Method+" | "+r.Proto+" | "+r.URL.Path)

} , nil
}

    if reqtype ==  "client_info" {


    }
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


