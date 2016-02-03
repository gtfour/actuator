package webclient

import "net/http"
import "encoding/json"
import "bytes"
import "errors"
import "wapour/settings"

//import "io/ioutil"

var TOKEN_COOKIE_FIELD_NAME  string = "USER_TOKEN"
var USERID_COOKIE_FIELD_NAME string = "USER_ID"


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

func (wp *WengineWrapper )Request(url string)(interface{},error) {

    return nil,nil
}

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

func SetAuthCookie( request *http.Request , user_id string, token_id string )(error) {

    cookie_userid := &http.Cookie{Name:settings.USERID_COOKIE_FIELD_NAME, Value:user_id,  Path:"/", Domain:settings.RESTAPI_SERVER_ADDR }
    cookie_token  := &http.Cookie{Name:settings.TOKEN_COOKIE_FIELD_NAME,  Value:token_id, Path:"/", Domain:settings.RESTAPI_SERVER_ADDR }
    request.AddCookie(cookie_token)
    request.AddCookie(cookie_userid)
    // func (r *Request) AddCookie(c *Cookie) {
    return nil

}

func GetResponseCookies(response *http.Response)(user_id string,token_id string) {
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

func FindWrapper(user_id string,token_id string ,wrappers *[]*WengineWrapper )(w *WengineWrapper) {

    for w := range (*wrappers) {
        wrapper:=(*wrappers)[w]
        if wrapper.UserId == user_id && wrapper.TokenId == token_id {
            return wrapper
        }
    }
    return nil
}

func   AppendWrapper ( wrappers *[]*WengineWrapper,w *WengineWrapper ) {
    (*wrappers)=append((*wrappers),w)
}
