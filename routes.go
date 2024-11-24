// HTTP and WebSocket endpoint handlers
package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func (cfg *apiConfig) handlerRegistry(mux *http.ServeMux) {

	mux.HandleFunc("GET /", cfg.startHandler)

}

func (cfg *apiConfig) startHandler(w http.ResponseWriter, r *http.Request) {
	websocket.Upgrader{HandshakeTimeout: time.Minute * 10, ReadBufferSize: 10, WriteBufferSize: 10}

}
