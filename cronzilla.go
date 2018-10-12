package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router   *mux.Router
	hostname string
	port     string
}

func main() {
	log.Print("This's Cornzilla!")

	srv := &Server{
		hostname: "127.0.0.1",
		port:     "2345",
		router:   mux.NewRouter(),
	}

	// Register routes
	srv.routes()

	n := negroni.Classic()
	n.UseHandler(srv.router)

	// Listen and serve app
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", srv.hostname, srv.port), n))
}
