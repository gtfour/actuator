package router

import ("net/http";"fmt")



func route_request_to_func(path string,method chan string) (func(w http.ResponseWriter, r *http.Request), error) {

    funcs:= map[string]interface{} {"receive_client_update":handle_client_update,
                                    "request_repo_update":handle_request_repo_update,
                                    "dashboard":render_dashboard}

    fmt.Println(funcs)

    return func(w http.ResponseWriter, r *http.Request) {

    text , _ := "test","hello"

    fmt.Fprintf(w, text)


} , nil
}

func handle_client_update(w http.ResponseWriter, r *http.Request){

}

func handle_request_repo_update(w http.ResponseWriter, r *http.Request){

}

func render_dashboard(w http.ResponseWriter, r *http.Request){

}
