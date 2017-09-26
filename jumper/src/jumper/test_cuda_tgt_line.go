package main

import "fmt"
import "encoding/json"
import "jumper/cuda/targets"
import "jumper/cuda/handling"
import "jumper/cuda/filtering"

type Line struct {
    Data []string `json:"data"`
}


func main(){
    lines := []string{`"a" "b" "c" "d" 'e'`, `"1" "2" "3333" [13/Sep/2017:12:25:57 +0300] "GET /static/wapour/fonts/fontawesome-webfont.woff2?v=4.4.0 HTTP/1.1"`}
    /*
    lines:=[]string{`127.0.0.1 - - [13/Sep/2017:12:25:57 +0300] "GET /static/wapour/fonts/fontawesome-webfont.woff2?v=4.4.0 HTTP/1.1" 200 66624 "http://127.0.0.1/static/wapour/css/font-awesome.min.css" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/127.0.0.1 Chrome/127.0.0.1 Safari/537.36"`,
                    `127.0.0.1 - - [13/Sep/2017:16:09:25 +0000] "GET /static/wapour/css/font-awesome.min.css HTTP/1.1" 200 26711 "http://127.0.0.1/wapour/auth/login?redirect_to=/userspace" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/127.0.0.1 Chrome/127.0.0.1 Safari/537.36"`}
    */
    handler :=  handling.NewHandler(nil)
    fl      :=  filtering.CreateDefaultFilterList()
    handler.AddFilters(fl)
    target_config          := make(map[string]string, 0)
    target_config["type"]  =  "SINGLE_LINE"
    tgt,_                  := targets.NewTarget(target_config)
    handler.AddTargetPtr(tgt)
    for i := range lines {
        line:=lines[i]
        tgt.SetLine(line)
        result,err :=  handler.Handle()
        if err == nil {
            var line Line
            result_js,err := result.GetJson()
            if err == nil {
                err_unmarshal:=json.Unmarshal(result_js,&line)
                if err_unmarshal == nil {
                    fmt.Printf("result_line: %v",line)
                }
            }
        }
    }
}
