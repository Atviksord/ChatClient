package main

import (
	"fmt"
	"os"
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
			name:        "logout",
			description: "logs you out",
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
func commandLogout(cfg *apiConfig) error {
	// System design: remove API key from user DB
	return nil
}
