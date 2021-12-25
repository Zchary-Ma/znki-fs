package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/znkisoft/znki-fs/server/handler"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	return &Server{router: newRouter()}
}

func newRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/").Subrouter()
	// TODO add []middleware
	api.Use()

	routerMap := map[string]map[string]handler.APIHandler{
		"/health": {
			http.MethodGet: handler.Health,
		},
		"/user": {
			http.MethodGet:  handler.GetUser,
			http.MethodPost: handler.CreateUser,
		},
	}

	for path, handlers := range routerMap {
		for m, h := range handlers {
			api.Path(path).Methods(m).Handler(h)
		}
	}

	return router
}

func (s *Server) ListenAndServe(host, port string) {
	if host == "" {
		host = "127.0.0.1"
	}

	if port == "" {
		port = "8080"
	}

	fmt.Printf("server is running on %s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.router))
}
