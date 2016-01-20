package main

import "fmt"
import "net/http"
import "io/ioutil"
import "bytes"

func main() {
    // curl --data 'userid=AF35CEFC-1AEA-A399-7448-C2EF4B80E77F' "http://127.0.0.1:9000/dusk/get-user-by-id"
    url := "http://10.10.111.143:9000/dusk/get-user-by-id"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"user_id":"AF35CEFC-1AEA-A399-7448-C2EF4B80E77F"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
