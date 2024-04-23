package controller

import (
	"fmt"
	"net/http"
	"log"
	model "forum/model"
)

var port string = ":15040"

func init() {
	if model.InitDB() != nil {
		log.Fatal("database doesn't exist")
	}
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
