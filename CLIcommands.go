package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Atviksord/ChatClient/internal/database"
)

func (cfg *apiConfig) getCommands() map[string]commandBlock {

	return map[string]commandBlock{
		"help": {
			name:        "help",
			description: "Displays commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},

		"logout": {
			name:        "logout",
			description: "logs you out",
			callback:    commandLogout,
		},
		"login": {
			name:        "login",
			description: "logs you in",
			callback:    commandLogin,
		},
		"signup": {
			name:        "signup",
			description: "signup",
			callback:    commandSignup,
		},
	}
}

func commandHelp(cfg *apiConfig) error {
	for _, v := range cfg.commands {
		fmt.Println("---------")
		fmt.Printf("%s: %s", v.name, v.description)
	}
	return nil
}
func commandExit(cfg *apiConfig) error {
	os.Exit(0)

	return nil
}
func commandSignup(cfg *apiConfig) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Password please:")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		_, err := cfg.db.CreateUser(context.Background(),
			database.CreateUserParams{Username: username,
				Password:  password,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC()})
		if err != nil {
			fmt.Println("User Creation failed ")
		}
		break

	}

	return nil
}

func commandLogin(cfg *apiConfig) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Password please:")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)
		// Will return user data for auth endpoints later
		_, err := cfg.db.LoginUser(context.Background(),
			database.LoginUserParams{Username: username,
				Password: password})
		if err != nil {
			fmt.Println("Login failed, Either password or Username wrong", err)
			return err
		}
		apiKey, err := generateAPIKey()
		if err != nil {
			fmt.Println("Error generating API key:", err)
			return err
		}
		fmt.Println("Generated API Key:", apiKey)
		_, err = cfg.db.AddApikey(context.Background(),
			database.AddApikeyParams{ApiKey: sql.NullString{String: apiKey, Valid: true},
				Username: username})
		if err != nil {
			fmt.Println("Error adding API key to DB")
			return err
		}

	}

}
func commandLogout(cfg *apiConfig) error {
	fmt.Println("LogoutHit test")
	return nil
}
