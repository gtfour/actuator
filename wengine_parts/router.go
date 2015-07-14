package router

import ("net/http";"fmt")


func route(path string,method chan string) (func(w http.ResponseWriter, r *http.Request), error) {

    apps:=map[string]string {"receive_client_update":"handle_client_update","request_repo_update":"handle_request_repo_update"}

    funcs:= map[string]interface{} {"handle_client_update":test1}

    

    fmt.Println(apps)
    fmt.Println(funcs)

    return func(w http.ResponseWriter, r *http.Request) {

    text , _ := "test","hello"

    fmt.Fprintf(w, text)


} , nil
}

func test1(path string ){

}

func test2(path string ){

}
func test3(path string ){

}
