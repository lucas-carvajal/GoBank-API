package userManagement

import "fmt"

type user struct {
	username string
	password string
	email    string
}

func getUsersByEmail(email string) []user {
	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email)
	return queryUsers(query)
}

func getUsersByUsername(username string) []user {
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", username)
	return queryUsers(query)
}

func queryUsers(query string) []user {
	var users []user
	rows, err := db.Query(query)
	if err != nil {
		fmt.Errorf("error while loading all users for query: %s", query)
		return []user{}
	}
	defer rows.Close()
	for rows.Next() {
		var usr user
		if err := rows.Scan(&usr.username, &usr.password, &usr.email); err != nil {
			fmt.Errorf("error while parsing rows for all users for query: %s", query)
			return []user{}
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("error after parsing rows for all users for query: %s", query)
		return []user{}
	}
	return users
}

func insertUser(user user) bool {
	_, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)",
		user.username, user.password, user.email)
	if err != nil {
		return false
	}
	return true
}
