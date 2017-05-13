package main

import "fmt"
import "net/http"
import "io/ioutil"
import "bytes"

func main() {
    // curl --data 'userid=AF35CEFC-1AEA-A399-7448-C2EF4B80E77F' "http://127.0.0.1:9000/dusk/get-user-by-id"
    //  curl --cookie "USER_TOKEN=B6F515B1-B549-BA1C-2497-67C1A8B10D4B;USER_ID=AF35CEFC-1AEA-A399-7448-C2EF4B80E77F"  --data 'userid=AF35CEFC-1AEA-A399-7448-C2"http://127.0.0.1:9000/rest/user/getuserbyid" 

    // curl --cookie "USER_TOKEN=763ED913-600B-C833-BC5D-78FDC5541F00;USER_ID=C5952D91-9AA5-4EEB-A21A-F138445103D5" --data "userid=C5952D91-9AA5-4EEB-A21A-F138445103D5" http://127.0.0.1:9000/rest/user/getuserbyid
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


    //
    // sending json data with curl
    //
    curl --data 'dashboardName="Users";sourceType="TARGET_FILE";sourcePath="/etc/passwd";clientId="my-ubuntu-host"' "http://127.0.0.1:9000/rest/dashboard/add-dashboard/"

    curl --data 'dashboardName="Users";sourceType="TARGET_FILE";sourcePath="/etc/passwd";clientName=1' "http://127.0.0.1:9000/rest/dashboard/add-dashboard/"

    //
    //
    //


}
