package wisel
// handle ws-messages

import "fmt"
import "wengine/wsserver"


type Router struct {
    web RouteWeb
    srv RouteSrv
}

type RouteWeb struct {
    server *wsserver.Server
    client_map *[]WebClient
}

type RouteSrv struct {
    server *wsserver.Server
    client_map *[]SrvClient
}

type WebClient struct {
    websocket_id string
    client_id    string
    connected_at string
    FirstName    string
}


type SrvClient struct {
    websocket_id string
    client_id    string
    connected_at string
}

func (r *Router)GetAllWebClients()(error) {
    clients:=r.web.server.Clients
    for i:= range clients {
        client:=clients[i]
        fmt.Printf("\nclient id %d\n",client.Id)
    }
    return nil
}

func (r *Router)GetAllHostsClients()(error) {
    clients:=r.srv.server.Clients
    for i:= range clients {
        client:=clients[i]
        fmt.Printf("\nclient id %d\n",client.Id)
    }
    return nil
}

