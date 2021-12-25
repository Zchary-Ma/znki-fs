package main

import (
	"github.com/znkisoft/znki-fs/server"
)

func main() {
	s := server.NewServer()
	s.ListenAndServe("", "")
}
