package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// struct to keep DB, ENV info etc
type apiConfig struct {
	db string
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
	// initialize database queries(SQL)
	dbQueries := database.New(db)
	cfg := &apiconfig{
		db: dbQueries,
	}

	PORT := os.Getenv("PORT")
	IP := os.Getenv("IP")

	serverIP := "0.0.0.0:8080"

	mux := http.NewServeMux()
	servr := http.Server{Addr: serverIP, Handler: mux}

	handlerRegistry(mux)

	err = servr.ListenAndServe()
	if err != nil {
		log.Fatal("Couldnt start server", err)
	}
}
