package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Atviksord/ChatClient/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// struct to keep DB, ENV info etc
type apiConfig struct {
	db *database.Queries
}

func main() {
	// Load Environmental variables/database
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Couldn load env variables")
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in the environment variables")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// Ping the database to confirm the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	// initialize database queries and store in struct
	dbQueries := database.New(db)
	cfg := &apiConfig{
		db: dbQueries,
	}

	PORT := os.Getenv("PORT")
	IP := os.Getenv("IP")

	mux := http.NewServeMux()
	servr := http.Server{Addr: IP + ":" + PORT, Handler: mux}

	go cfg.handlerRegistry(mux)

	err = servr.ListenAndServe()
	if err != nil {
		log.Fatal("Couldnt start server", err)
	}
}
