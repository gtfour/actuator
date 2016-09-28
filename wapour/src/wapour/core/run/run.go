package run

import "flag"

import "go/ast"
import "go/importer"
import "go/parser"
import "go/token"
import "go/types"


var default_settings_path          = "/etc/wapour/settings.go"
var default_relative_settings_path =  "../../settings/settings.go"
var app_settings_package_name = "wapour/settings"
//var app_settings_directory    = "/actuator/wapour/src/wapour/settings"

var initial = []string {

    "SERVER_ADDR",
    "SERVER_PORT",
    "SERVER_PROTO",
    "SERVER_URL",
    "RESTAPI_SERVER_ADDR",
    "RESTAPI_SERVER_PROTO",
    "RESTAPI_SERVER_PORT",
    "RESTAPI_URL",

}

var Props = GetProps()

func GetProps()(props map[string]string){

    props=make(map[string]string,0)

    // serve
    // flag.Lookup

    ip_version_ptr       := flag.String("ipversion","v4",      "Server ip version"  )
    ip_port_ptr          := flag.String("port",     "80",      "Server port number" )
    ip_addr_ptr          := flag.String("addr",     "0.0.0.0", "Server ip address"  )
    //settings_file_path   := flag.String("SETTINGS", default_settings_path , "Server ip address"  )

    flag.Parse()
    //out_file_ptr       := flag.String("outfile","out.txt","Out file")

    ip_version       := *ip_version_ptr
    ip_port          := *ip_port_ptr
    ip_addr          := *ip_addr_ptr

    props["server_ip_version"]  = ip_version
    props["server_port"]        = ip_port
    props["server_addr"]        = ip_addr

    props["server_addr"] = props["ip_addr"]+":"+props["ip_port"]

    return

}

func GetCurrentAppSettings(settings_path string)(settings map[string]string,err error){

    fset := token.NewFileSet()

    f, err := parser.ParseFile(fset, default_settings_path, nil, 0)
    if err != nil {
        f, err = parser.ParseFile(fset, default_relative_settings_path, nil, 0)
        if err!=nil{
            return nil,err
        }
    }

    conf := types.Config{Importer: importer.Default()}
    pkg, err := conf.Check("wapour/settings", fset, []*ast.File{f}, nil)
    if err != nil {
        return nil,err
    }
    settings=make(map[string]string,0)
    for word_id := range initial {
        word:=initial[word_id]
        existing_set:=pkg.Scope().Lookup(word).(*types.Const).Val().String()
        settings[word]=existing_set
    }
    return settings,err
}



