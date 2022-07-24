package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gobank.com/services/userManagement"
)

func main() {
	fmt.Println("Starting GoBank Server...")
	db, err := setUpAndTestDBConnection()
	if err != nil {
		fmt.Printf("There was a problem connecting to the database: %s", err.Error())
		fmt.Println("Shutting down...")
		return
	}

	userManagementServiceChannel := userManagement.Init(db)
	//TODO initialize ClaimManagementService

	runServer(userManagementServiceChannel)
}

func setUpAndTestDBConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		//User:   os.Getenv("DBUSER"),
		//Passwd: os.Getenv("DBPASS"),
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "gobank",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	fmt.Println("Connected to database successfully")
	return db, nil
}

// delete later
//func testDB() {
//	connectToDB()
//	users, err := loadAllUsers()
//	if err == nil {
//		fmt.Println("USERS:")
//		for _, user := range users {
//			fmt.Printf("Username: %s\n", user.username)
//			fmt.Printf("Password: %s\n", user.password)
//			fmt.Printf("Email: %s\n", user.email)
//			fmt.Println("===")
//		}
//	}
//}
