package main

import "fmt"
import "net/http"
import "io/ioutil"
import "bytes"

type TestJson struct {

    Data string `json:"data"`

}
func main() {
    url := "http://127.0.0.1:8081/clinfo"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"data":"Buy cheese and bread for breakfast."}`)
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
