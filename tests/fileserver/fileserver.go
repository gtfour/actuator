package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/home/venom/gopro/repo_data/rpm/"))))
        log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/tmp/static_test1"))))
        //log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/tmp/static_test2"))))
}
