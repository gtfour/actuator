package webclient

import "net/http"
import "encoding/json"
import "bytes"
//import "io/ioutil"
import "fmt"

type Proxy interface {

    Connect()(error)
    /*Request()
    UsersList()
    UserLogin()
    UserLogout()*/

}

type WengineProxy struct {
    username       string
    password       string
    url            string
    user_id        string
    token_id       string
    //proxyUserId    string
    //proxyUserToken string
    client         *http.Client
}

type WengineUser struct {

    user_id  string
    token_id string


}

type Credentials struct {

    Username string `json:"username"`
    Password string `json:"password"`

}

var WENGINE_AUTH_LOGIN_URL  string = "/auth/login"
var WENGINE_AUTH_LOGOUT_URL string = "/auth/logout"

func (wp *WengineProxy )Connect()(error) {
    wp.client=&http.Client{}
    credentials:=Credentials{Username:wp.username,Password:wp.password}
    credentials_json,err:=json.Marshal(credentials)
    fmt.Printf("\nMarshal: %v\n",bytes.NewBuffer(credentials_json))
    if err!=nil { return err }
    auth_url:=wp.url+WENGINE_AUTH_LOGIN_URL
    req, err := http.NewRequest("POST", auth_url, bytes.NewReader(credentials_json))
    fmt.Printf("Error ::: %v",err)
    req.Header.Set("Content-Type", "application/json")
    //req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    resp,err := wp.client.Do(req)
    //
    if err!=nil { return err }
    defer resp.Body.Close()
    wp.user_id,wp.token_id=GetResponseCookies(resp)

    //fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    //body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))
    return nil

}

func Init( ptype,url,username,password string) ( p Proxy,err error ) {

    switch {
        case ptype == "wengine":
            p=&WengineProxy{username:username,
                      password:password,
                      url:url}
            err=p.Connect()
            return p,err
        case ptype == "something_else":
    }
    return p,nil

}

func SetAuthCookie( r *http.Request , user_id string, token_id string )(error) {


    // func (r *Request) AddCookie(c *Cookie) {
    return nil

}

func GetResponseCookies(response *http.Response)(user_id string,token_id string) {
    var TOKEN_COOKIE_FIELD_NAME  string = "USER_TOKEN"
    var USERID_COOKIE_FIELD_NAME string = "USER_ID"
    cookies:=response.Cookies()
    for i := range cookies {
        cookie:=cookies[i]
        if cookie.Name == TOKEN_COOKIE_FIELD_NAME {
            token_id = cookie.Value
        }
        if cookie.Name == USERID_COOKIE_FIELD_NAME  {
            user_id = cookie.Value
        }
    }
    return
}
