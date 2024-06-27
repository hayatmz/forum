package controller

import (
	"fmt"
	"net/http"
	"log"
	model "forum/model"
)

const port string = ":15040"

// init the database and open it
func init() {
	if model.InitDB() != nil {
		log.Fatal("database doesn't exist")
	}
}

// main function, listen and serve the server, load the assets, and listen the handlers
func Controller() {
	var mux *http.ServeMux = http.NewServeMux()
	handlers(mux)
	server := http.Server {
		Addr: port,
		Handler: mux,
	}

	mux.Handle("/view/static/", http.StripPrefix("/view/static/", http.FileServer(http.Dir("view/static"))))

	fmt.Println("http://localhost"+port)
	server.ListenAndServe()
}
