package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Will build a real time chat client with sockets here.")

	d := http.NewServeMux()
	servr := http.Server{Addr: "0.0.0.0.0", Handler: d}

	http.ListenAndServe("0.0.0.0", d)
}
