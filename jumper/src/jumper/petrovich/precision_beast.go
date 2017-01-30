package petrovich

type Huyamba struct {
    initial_config    map[string]string
    default_config    map[string]string
    cmdline_config    map[string]string
    self_config       map[string]string
    ontherun_config   map[string]string
    Config            map[string]string


    config_file_path  string
    config_dir_path   string
    config_type       string
    AppName           string
    Description       string
}

//
//
//
// -- -- -- -- -- -- -- -- -- -- -- --
// search value strategy
// initial_config -> ontherun_config 
// -- -- -- -- -- -- -- -- -- -- -- --
//
//
//

func(h *Huyamba)GetValue(key string)(value string,err int){
    return value,err
}

func(h *Huyamba)ProceedLine(line string)(){
}

func CreateHuyamba(initial_config ...map[string]string)(parser Huyamba){
    parser.initial_config = make(map[string]string,0)
    if len(initial_config)>0{
        first_initial_config:=initial_config[0]
        for key,value := range first_initial_config {
            parser.initial_config[key] = value
        }
    }
    return parser
}

func GetMapValue(key string, mymap map[string]string)(value string,err error){
    return
}

func GetHuyamba()(h *Huyamba,e error){
    var myHuyamba Huyamba
    myHuyamba.Description = "Empty config"
    return &myHuyamba, nil
}
