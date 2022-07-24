package main

import (
	"fmt"
	"gobank.com/services/userManagement"
	"gobank.com/util"
	"math/rand"
	"net/http"
	"strings"
)

func handleSignup(
	w http.ResponseWriter,
	r *http.Request,
	userMgmChan chan userManagement.Request,
) {
	username, email, password := r.Header.Get("username"), r.Header.Get("email"), r.Header.Get("password")
	if username == "" {
		sendErrorResponse(w, "Username was empty", http.StatusBadRequest)
		return
	}
	if email == "" {
		sendErrorResponse(w, "Email was empty", http.StatusBadRequest)
		return
	}
	if password == "" {
		sendErrorResponse(w, "Password was empty", http.StatusBadRequest)
		return
	}
	responseChannel := make(chan userManagement.Response)
	userMgmChan <- userManagement.Request{
		Id:              rand.Intn(1000000),
		Command:         "SIGNUP",
		Username:        username,
		Email:           email,
		Password:        password,
		ResponseChannel: responseChannel,
	}
	handleSignupResponse(responseChannel, w)
	return
}

func handleLogin(
	w http.ResponseWriter,
	r *http.Request,
	userMgmChan chan userManagement.Request,
) {
	// TODO
}

func handleClaims(
	w http.ResponseWriter,
	r *http.Request,
	userMgmChan chan userManagement.Request,
) {
	urlParts := strings.Split(r.URL.String(), "/")
	urlParts = util.Filter(urlParts, func(s string) bool {
		return s != ""
	})

	//TODO remove
	fmt.Println("===")
	fmt.Printf("route \"/claim*\" - URL parts: %s\n", urlParts)

	switch {
	case urlParts[0] == "claim" && len(urlParts) == 1:
		// claim route "/claim" -> file a claim
		doSomething(w, r)
		fmt.Println("/claim route")
		return
	case urlParts[0] == "claim" && len(urlParts) > 1:
		switch {
		case len(urlParts) == 2:
			// claim id route "/claim/1424" -> retrieve details for id=1424
			doSomething(w, r)
			fmt.Println("/claim/{} route")
			return
		case urlParts[2] == "approve" && len(urlParts) == 3:
			// approve claim route "/claim/1424/approve"
			doSomething(w, r)
			fmt.Println("/claim/{}/approve route")
			return
		case urlParts[2] == "deny" && len(urlParts) == 3:
			// deny claim route "/claim/1424/deny"
			doSomething(w, r)
			fmt.Println("/claim/{}/deny route")
			return
		case urlParts[2] == "settleRequest" && len(urlParts) == 3:
			// settle claim request route "/claim/1424/settleRequest"
			doSomething(w, r)
			fmt.Println("/claim/{}/settleRequest route")
			return
		case urlParts[2] == "settle" && len(urlParts) == 3:
			// settle claim route "/claim/1424/settle"
			doSomething(w, r)
			fmt.Println("/claim/{}/settle route")
			return
		case urlParts[2] == "deadline" && len(urlParts) == 3:
			// set deadline for claim route "/claim/1424/deadline"
			doSomething(w, r)
			fmt.Println("/claim/{}/deadline route")
			return
		case urlParts[2] == "interest" && len(urlParts) == 3:
			// set interest for claim route "/claim/1424/interest"
			doSomething(w, r)
			fmt.Println("/claim/{}/interest route")
			return
		default:
			// TODO return error message
			fmt.Println("ERROR")
			return
		}
	default:
		// TODO return error message
		fmt.Println("ERROR")
		return
	}
}

func doSomething(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Call to %s\n", r.URL.String())
}

func handleSignupResponse(responseChannel chan userManagement.Response, w http.ResponseWriter) {
	response := <-responseChannel
	switch response.Status {
	case "SUCCESS":
		sendSuccessResponse(w, http.StatusOK)
	case "ERROR":
		sendErrorResponse(w, response.Message, http.StatusBadRequest)
	}
}

func sendSuccessResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func sendSuccessResponseWithToken(w http.ResponseWriter, statusCode int, token string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(token))
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
