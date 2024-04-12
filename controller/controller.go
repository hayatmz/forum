package controller

import (
	"fmt"
	"net/http"
)

var port string = ":15040"

func Controller() {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server {
		Addr: port,
		Handler: mux,
	}
	fmt.Println("http://localhost"+port)
	server.ListenAndServe()
}
