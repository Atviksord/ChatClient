package main

import (
	"fmt"
	"log"
	"net/http"
)

// struct to keep DB, ENV info etc
type apiConfig struct {
	db string
}

func main() {
	fmt.Println("Will build a real time chat client with sockets here.")

	serverIP := "0.0.0.0:8080"

	mux := http.NewServeMux()
	servr := http.Server{Addr: serverIP, Handler: mux}

	cfg.handlerRegistry(mux)
	err := servr.ListenAndServe()
	if err != nil {
		log.Fatal("Couldnt start server", err)
	}
}
