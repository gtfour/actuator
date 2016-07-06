package blackout

import "time"
import "strings"

type Bounce struct {
    BinPath  string
    Timeout  time.Duration
    Keys     []Key
    Out      strings.Reader
}

type BounceFlow struct {


}

type Key struct {
    Name  string
    Value string
}
