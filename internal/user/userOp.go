package userOp

import (
	"crypto/sha256"
	"fmt"
)

// users info
type User struct {
	username string
	password []byte
}

func AddUser(users []User, username string, password string) []User {
	newUser := User{}
	hash := sha256.New()

	newUser.username = username
	hash.Write([]byte(password))
	newUser.password = hash.Sum(nil)

	users = append(users, newUser)

	return users
}

func FindUser(users []User, username string, password string) int {
	//find user in the list of structs and return its index
	return -1
}

func ListUsers(users []User) {
	numberUsers := len(users)
	fmt.Println("Number of registered users ", numberUsers)
}
