package controller

import (
	"fmt"
	"net/http"
	"forum/model"
)

var port string = ":15040"

func init() {
	model.InitDB()
}

func Controller() {
	var mux *http.ServeMux = http.NewServeMux()
	handlers(mux)
	server := http.Server {
		Addr: port,
		Handler: mux,
	}
	fmt.Println("http://localhost"+port)
	server.ListenAndServe()
}
