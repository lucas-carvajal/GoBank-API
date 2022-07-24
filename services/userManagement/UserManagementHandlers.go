package userManagement

import "fmt"

func handleSignup(dto Request) {
	users := getUsersByEmail(dto.Email)
	if len(users) != 0 {
		dto.ResponseChannel <- Response{
			Id:      dto.Id,
			Status:  "ERROR",
			Message: fmt.Sprintf("user with email '%s' already exists", dto.Email),
			Token:   "",
		}
		return
	}
	users = getUsersByUsername(dto.Username)
	if len(users) != 0 {
		dto.ResponseChannel <- Response{
			Id:      dto.Id,
			Status:  "ERROR",
			Message: fmt.Sprintf("user with username '%s' already exists", dto.Username),
			Token:   "",
		}
		return
	}
	success := insertUser(user{
		username: dto.Username,
		email:    dto.Email,
		password: dto.Password,
	})
	if !success {
		dto.ResponseChannel <- Response{
			Id:      dto.Id,
			Status:  "ERROR",
			Message: fmt.Sprintf("user with email '%s' could not be saved", dto.Email),
			Token:   "",
		}
		return
	}
	dto.ResponseChannel <- Response{
		Id:      dto.Id,
		Status:  "SUCCESS",
		Message: fmt.Sprintf("user with email '%s' successfully created", dto.Email),
		Token:   "",
	}
}

func handleLogin(dto Request) {
	//TODO
	// check if user with email exists
	// check if password matches
	// return response with auth token
}
