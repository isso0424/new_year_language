package server

import (
	"net/http"
	"new_year_language/server/handler"
)

func Serve() {
	server := &http.Server{ Addr: ":8080" }
	http.HandleFunc("/encode", handler.EncodeHandler)
	http.HandleFunc("/decode", handler.DecodeHandler)

	server.ListenAndServe()
}
