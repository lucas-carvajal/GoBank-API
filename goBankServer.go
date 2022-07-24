package main

import (
	"fmt"
	"gobank.com/services/userManagement"
	"log"
	"net/http"
)

const port int = 8080

func runServer(userMgmChan chan userManagement.Request) {
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handleSignup(w, r, userMgmChan)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handleLogin(w, r, userMgmChan)
	})
	http.HandleFunc("/claim", func(w http.ResponseWriter, r *http.Request) {
		handleClaims(w, r, userMgmChan)
	})
	http.HandleFunc("/claim/", func(w http.ResponseWriter, r *http.Request) {
		handleClaims(w, r, userMgmChan)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO
		fmt.Printf("route \"/\": %s", r.URL)
	})

	fmt.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
