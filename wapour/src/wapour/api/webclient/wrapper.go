package webclient

import "net/http"
import "encoding/json"
import "bytes"
import "errors"
import "wapour/settings"

//import "io/ioutil"


type WengineWrapper struct {
    username       string
    password       string
    url            string
    UserId         string
    TokenId        string
}


type Credentials struct {

    Username string `json:"username"`
    Password string `json:"password"`

}

var WENGINE_AUTH_LOGIN_URL  string = "/auth/login"
var WENGINE_AUTH_LOGOUT_URL string = "/auth/logout"

func (wp *WengineWrapper )Connect()(error) {
    client  :=  &http.Client{}
    credentials:=Credentials{Username:wp.username,Password:wp.password}
    credentials_json,err:=json.Marshal(credentials)
    if err!=nil { return err }
    auth_url  :=wp.url + WENGINE_AUTH_LOGIN_URL
    req, err := http.NewRequest("POST", auth_url, bytes.NewReader(credentials_json))
    req.Header.Set("Content-Type", "application/json")
    resp,err                  :=  client.Do(req)
    if err!=nil {  return err }
    defer                       resp.Body.Close()
    wp.UserId,wp.TokenId    = GetResponseCookies(resp)
    if (wp.UserId=="" || wp.TokenId == "")  { return errors.New("token_id or user_id was not found in cookie") }
    return nil

}

func Init( username, password string) ( w WengineWrapper,err error ) {


    w=WengineWrapper{username:username, password:password, url:settings.RESTAPI_URL}
    err=w.Connect()
    return w,err
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

func FindWrapper(user_id string,token_id string ,wrappers []*WengineWrapper )(w *WengineWrapper) {

    for w := range wrappers {
        wrapper:=wrappers[w]
        if wrapper.UserId == user_id && wrapper.TokenId == token_id {
            return wrapper
        }
    }
    return nil
}

func   AppendWrapper ( wrappers []*WengineWrapper,w *WengineWrapper ) {
    wrappers=append(wrappers,w)
}
