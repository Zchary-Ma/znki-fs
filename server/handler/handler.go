package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"

	"google.golang.org/protobuf/proto"
)

type APIHandler func(w http.ResponseWriter, r *http.Request) (proto.Message, error)

func (fn APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	message, err := fn(w, r)
	if err != nil {
		// TODO expand error type
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// marshal message
	if message == nil {
		return
	}

	accept := r.Header.Get("Accept")
	switch accept {
	case "application/json":
		marshalJSON(ctx, w, message)
	case "application/x-protobuf":
		marshalProtobuf(ctx, w, message)
	default:
		marshalProtobuf(ctx, w, message)
	}

}

func UnmarshalRequest(r *http.Request, msg proto.Message) error {
	ct := r.Header.Get("Content-Type")
	switch ct {
	case "application/json":
		return unmarshalJSON(r, msg)
	case "application/protobuf":
		return unmarshalProtobuf(r, msg)
	}
	return unmarshalProtobuf(r, msg)
}

func unmarshalJSON(r *http.Request, msg proto.Message) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(body, msg)
}

func unmarshalProtobuf(r *http.Request, msg proto.Message) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return proto.Unmarshal(body, msg)
}

func marshalJSON(ctx context.Context, w http.ResponseWriter, m proto.Message) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(m)
	if err != nil {
		httpError(ctx, w, err, "failed to marshal message")
	}
	if _, err := w.Write(bytes); err != nil {
		httpError(ctx, w, err, "failed to write response")
	}
}

func marshalProtobuf(ctx context.Context, w http.ResponseWriter, m proto.Message) {
	w.Header().Set("Content-Type", "application/x-protobuf")
	bytes, err := proto.Marshal(m)
	if err != nil {
		httpError(ctx, w, err, "failed to marshal message")
	}

	if _, err := w.Write(bytes); err != nil {
		httpError(ctx, w, err, "failed to write response")
	}
}

func httpError(ctx context.Context, w http.ResponseWriter, err error, msg string) {
	http.Error(w, fmt.Sprintf("cause: %s, msg: %s", err, msg), http.StatusInternalServerError)
}
