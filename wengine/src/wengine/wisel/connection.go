package wisel

type Session struct {
    //
    connections []*Connection
    //
}

type Connection struct {
    //
    id                  string   //
    clientId            string   // websocket client id
    subscriptionsALL    []string // this is a list with dashboard source id's (dynima id's)
    subscriptionsVIEWED []string // dashboards that's open now on client side
    active              bool     // connection is active
    //
}

type Dynima struct {
    //
}


func( s *Session )NewConnection( clientId string )(*Connection){
    //
    var c Connection
    //
    return &c
    //
}



//func( c *Connection )
