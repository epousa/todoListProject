package userOp

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// users info
type User struct {
	username string
	password []byte
}

func AddUser(users *[]User, username string, password []byte) {
	newUser := User{}
	hash, err := bcrypt.GenerateFromPassword(password, 15)
	if err != nil {
		log.Fatal(err)
	}
	newUser.username = username
	newUser.password = hash

	*users = append(*users, newUser)
}

func FindUser(users []User, username string, password string) int {
	//find user in the list of structs and return its index
	return -1
}

func PrintUsers(users *[]User) {
	numberUsers := len(*users)
	fmt.Printf("Number of registered users %d;\n", numberUsers)
	for _, user := range *users {
		fmt.Printf("\nUsername: %sPassword: %d\n", user.username, user.password)
	}
}
