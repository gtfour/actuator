package main

import "fmt"
import "encoding/json"
import "jumper/cuda/targets"
import "jumper/cuda/handling"
import "jumper/cuda/filtering"
import "jumper/cuda/analyze"

type Line struct {
    Data []string `json:"data"`
}


func main(){
    myString1:=`127.0.0.1 - - [13/Sep/2017:12:25:57 +0300] "GET /static/wapour/fonts/fontawesome-webfont.woff2?v=4.4.0 HTTP/1.1" 200 66624 "http://127.0.0.1/static/wapour/css/font-awesome.min.css" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/127.0.0.1 Chrome/127.0.0.1 Safari/537.36"`
    myString2:=`78.231.89.119 - - [13/Sep/2017:16:09:25 +0000] "GET /static/wapour/css/font-awesome.min.css HTTP/1.1" 200 26711 "http://127.0.0.1/wapour/auth/login?redirect_to=/userspace" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/127.0.0.1 Chrome/127.0.0.1 Safari/537.36"`
    //myString1:="0.1.2.3 0.1.2.3"
    //myString2:="gg.dsc.2 2 a    c ddd"
    lines:=[]string{myString1, myString2}

    fmt.Printf("\n<<<  String 1:  >>>\n")
    analyze.DebugPrintCharCounter(myString1)
    fmt.Printf("\n<<<  String 2:  >>>\n")
    analyze.DebugPrintCharCounter(myString2)
    fmt.Printf(": === :\n")
    handler             :=  handling.NewHandler(nil)
    //
    // initial filter list to load square_filter and  prevent loading colon_filter from default set
    //
    var fl filtering.FilterList
    var sq_filter       = filtering.Filter{ Name:"square_brackets_filter", Call:filtering.SquareBracketsFilter, Enabled:true }
    var url_filter      = filtering.Filter{ Name:"url_filter",             Call:filtering.UrlFilter,            Enabled:true }
    var path_filter     = filtering.Filter{ Name:"path_filter",            Call:filtering.PathFilter,           Enabled:true }
    var quotes_filter   = filtering.Filter{ Name:"quotes_filter",          Call:filtering.QuotesFilter,         Enabled:true }
    var dot_filter     =  filtering.Filter{ Name:"dot_filter",             Call:filtering.DotFilter,            Enabled:true }
    fl.Append(sq_filter)
    fl.Append(url_filter)
    fl.Append(path_filter)
    fl.Append(quotes_filter)
    fl.Append(dot_filter)
    handler.AddFilters(fl)
    //
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
                    fmt.Printf("result_line: %v\n",line)
                    for i := range line.Data {
                        fmt.Printf("i:%d string:%s\n",i,line.Data[i])
                    }
                }
            }
        }
    }
}
