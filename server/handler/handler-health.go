package handler

import (
	"net/http"

	"google.golang.org/protobuf/proto"

	"github.com/znkisoft/znki-fs/api"
)

func Health(w http.ResponseWriter, r *http.Request) (proto.Message, error) {
	return &api.ServerStatus{Status: "OK"}, nil
}
