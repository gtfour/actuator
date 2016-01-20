package common

import "crypto/rand"
import "fmt"

func GenId()(uuid string,err error) {
    b := make([]byte, 16)
    _,err= rand.Read(b)
    if err!= nil {
        return "",err
    }
    uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4],b[4:6],b[6:8],b[8:10],b[10:])
    return uuid, nil
}


func CombineErrors(errors ...error) ( map[string][]string ) {
    errors_map :=make(map[string][]string)
    errors_map["errors"] = []string {}
    for e := range errors {
        if (errors[e]!=nil) {
            errors_map["errors"] = append(errors_map["errors"], fmt.Sprintf("%v",errors[e]))
        }
    }
    return errors_map
}
