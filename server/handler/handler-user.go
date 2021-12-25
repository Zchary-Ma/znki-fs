package handler

import (
	"net/http"

	"github.com/znkisoft/znki-fs/api"

	"google.golang.org/protobuf/proto"
)

func GetUser(w http.ResponseWriter, r *http.Request) (proto.Message, error) {
	return &api.User{Id: "1"}, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) (proto.Message, error) {
	return nil, nil
}
