// HTTP and WebSocket endpoint handlers
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func (cfg *apiConfig) handlerRegistry(mux *http.ServeMux) {

	mux.HandleFunc("GET /", cfg.startHandler)

}

func (cfg *apiConfig) startHandler(w http.ResponseWriter, r *http.Request) {
	upper := websocket.Upgrader{HandshakeTimeout: time.Minute * 10, ReadBufferSize: 0, WriteBufferSize: 0, CheckOrigin: nil}
	connection, err := upper.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer connection.Close()
}
