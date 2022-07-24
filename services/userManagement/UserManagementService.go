package userManagement

import (
	"database/sql"
	"fmt"
)

var channel chan Request
var db *sql.DB

func Init(inputDb *sql.DB) chan Request {
	channel = make(chan Request)
	db = inputDb

	go run()

	return channel
}

func run() {
	for in := range channel {
		fmt.Println("[UserManagementService Channel] Message received with command: ", in.Command)

		switch in.Command {
		case "SIGNUP":
			handleSignup(in)
		case "LOGIN":
			handleLogin(in)
		}

	}
	fmt.Println("[UserManagementService Channel] Closed")
}
