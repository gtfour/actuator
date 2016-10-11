package petrovich

type huyamba struct {
    initial_config map[string]string
    self_config    map[string]string
}

func CreateHuyamba(initial_config ...map[string]string)(parser huyamba){
    if len(initial_config)>0{
        for key := range initial_config {
        }

    }
    return parser
}
