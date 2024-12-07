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
	mux.HandleFunc("GET /ws", cfg.establishConnectionHandler)

}

func (cfg *apiConfig) startHandler(w http.ResponseWriter, r *http.Request) {
	// auth check

}

// Function that establishes the websocket upgrade
func (cfg *apiConfig) establishConnectionHandler(w http.ResponseWriter, r *http.Request) {
	upper := websocket.Upgrader{HandshakeTimeout: time.Minute * 10, ReadBufferSize: 0, WriteBufferSize: 0, CheckOrigin: nil}
	connection, err := upper.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	// if connection is properly upgraded send it into a goroutine.
	go func(connection *websocket.Conn) {
		for {
			msgType, msg, err := connection.ReadMessage()
			if err != nil {
				log.Println("Connection closed", err)
				break
			}
			log.Printf("Received: %s", msg)
			connection.WriteMessage(msgType, msg)
		}
	}(connection)
	defer connection.Close()
}
