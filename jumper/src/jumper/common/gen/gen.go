package gen

import "fmt"
import "time"
import "crypto/rand"

func GenId()(uuid string,err error) {
    b := make([]byte, 16)
    _,err= rand.Read(b)
    if err!= nil {
        return "",err
    }
    uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4],b[4:6],b[6:8],b[8:10],b[10:])
    return uuid, nil
}

func GetTime()(time_now string) {
    t := time.Now()
    return t.Format(time.RFC3339Nano)
}

func GetTimeShort()(time_now string) {
    t := time.Now()
    return t.Format(time.RFC822)
}

