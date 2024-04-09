package controller

import (
	"fmt"
	"net/http"
)

var port string = ":16040"

func Controller() {
	var server *http.ServeMux = http.NewServeMux()
	server.HandleFunc("/", handler)
	fmt.Println("localhost"+port)
	http.ListenAndServe(port, nil)
}
