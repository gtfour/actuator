package main

import (
	"log"
	"net/http"
)

func main() {

        //log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/adminlte/AdminLTE"))))
        //log.Fatal(http.ListenAndServe(":8080", nil))
        //log.Fatal(http.ListenAndServe(":8080/angular", http.FileServer(http.Dir("/actuator/tests/web_tests"))))
        http.Handle("/angular/", http.StripPrefix("/angular/", http.FileServer(http.Dir("/actuator/tests/web_tests"))))
        http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/adminlte"))))
        log.Fatal(http.ListenAndServe(":8080", nil))

}
