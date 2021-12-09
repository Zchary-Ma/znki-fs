package main

import (
	"github.com/gorilla/mux"
	"github.com/znkisoft/znki-fs/handler"
	"log"
	"net/http"
)

func main() {
	// new router with mux
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/ping", handler.PingHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
