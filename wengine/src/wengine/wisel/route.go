package wisel

import "wengine/wsserver"

type Router struct {
   webclient *wsserver.Server
   srvclient *wsserver.Server
}
