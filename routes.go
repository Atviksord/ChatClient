// HTTP and WebSocket endpoint handlers
package main

import "net/http"

func (cfg *apiConfig) handlerRegistry(mux *http.ServeMux) {

	mux.HandleFunc("GET /", cfg.startHandler)

}

func (cfg *apiConfig) startHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type:")

}
