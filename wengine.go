//server side
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
//    "errors"
//    "fmt"
    "os"
)



func main() {

    port , _, _ := get_args()

    port = ":"+port

    http.HandleFunc("/", homeHandler)
    panic(http.ListenAndServe(port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    _,filepath,_:=get_args()
    text , _ := readFile(filepath)
    fmt.Fprintf(w, text)

}

func readFile(filename string) (text string ,err error ) {
    rawBytes, err := ioutil.ReadFile(filename)
    text = string(rawBytes)
    return text, nil
}

func get_args()(port string,filepath string , err error) {


    if len(os.Args) == 3 {

        port = os.Args[1]
        filepath = os.Args[2]
        if inFile, err := os.Open(filepath); err != nil{
            defer inFile.Close()
            return "","",err
        }
        }
        return port,filepath,nil



}
