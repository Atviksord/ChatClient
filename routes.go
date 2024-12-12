// HTTP and WebSocket endpoint handlers
package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Atviksord/ChatClient/internal/database"
	"github.com/gorilla/websocket"
)

func (cfg *apiConfig) handlerRegistry(mux *http.ServeMux) {

	mux.HandleFunc("GET /", cfg.startHandler)
	mux.HandleFunc("GET /ws", cfg.establishConnectionHandler)

}

func (cfg *apiConfig) startHandler(w http.ResponseWriter, r *http.Request) {
	// auth check

	// if not authed, infinite loop awaiting server commands
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Password please:")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		_, err := cfg.db.CreateUser(r.Context(),
			database.CreateUserParams{Username: username,
				Password:  password,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC()})
		if err != nil {
			fmt.Println("User Creation failed ")
		}
		break

	}
	// Login sequence
	for {
		fmt.Print("Enter username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Password please:")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		user, err := cfg.db.LoginUser(context.Background(),
			database.LoginUserParams{Username: username,
				Password: password})
		if err != nil {
			fmt.Println("Login failed")
			return
		}
		apiKey, err := generateAPIKey()
		if err != nil {
			fmt.Println("Error generating API key:", err)
			return
		}
		fmt.Println("Generated API Key:", apiKey)
		cfg.db.AddApikey(context.Background(),
			database.AddApikeyParams{ApiKey: sql.NullString{String: apiKey, Valid: true},
				Username: username})

	}
	// Login initiated right after creation

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
