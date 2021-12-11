package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/znkisoft/znki-fs/handler"
	"log"
	"net/http"
)

var (
	host string = "127.0.0.1"
	port string = "8080"
)

func main() {
	// new router with mux
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/ping", handler.PingHandler)
	r.HandleFunc("/user", handler.UserHandler)

	fmt.Printf("server is running on %s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r))
}
