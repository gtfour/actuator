package webclient

import "net/http"

type Proxy interface {

    Connect(error)
    Request()
    UsersList()
    UserLogin()
    UserLogout()


}

type WengineProxy struct {

    proxyUserId    string
    proxyUserToken string
    client         *http.Client
    url            string

}

func CreateProxy ( ptype string, phost string, username string, password string ) ( p *Proxy ) {

    return p


}

func Init( ptype, url,  username, password ) ( p Proxy ) {

    switch {
        case ptype == "wengine":
            d=&WengineProxy{username:username,
                      password:password,
                      host:host,
                      dbname:dbname}
            err:=d.Connect()
            if err == nil {
                return d
            }
        case dbtype == "postgres":
    }
    return nil



}
