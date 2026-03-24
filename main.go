package main

import (
	"log"
	"net/http"

	"github.com/sunesimonsen/microbe/server"
)

func main() {
	port := "8080"

	srv, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, srv); err != nil {
		log.Fatal(err)
	}
}
