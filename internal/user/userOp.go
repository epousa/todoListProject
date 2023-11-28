package userOp

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// users info
type User struct {
	id       int
	username string
	password []byte
}

func AddUser(users *[]User, username string, password []byte) {
	newUser := User{}
	hash, err := bcrypt.GenerateFromPassword(password, 15)
	if err != nil {
		log.Fatal(err)
	}
	newUser.id = len(*users) + 1
	newUser.username = username
	newUser.password = hash

	*users = append(*users, newUser)
}

func FindUser(users []User, id int) int {
	//find user in the list of structs and return its index
	for i := 0; i < len(users); i++ {
		if users[i].id == id {
			return i
		}
	}
	return -1
}

func RemoveUser(users *[]User, id int) int {
	userIndex := FindUser(*users, id)
	if userIndex == -1 {
		return -1
	}
	*users = append((*users)[:userIndex], (*users)[userIndex+1:]...)
	return 0
}

func PrintUsers(users []User) {
	numberUsers := len(users)
	fmt.Printf("Number of registered users %d;\n", numberUsers)
	for _, user := range users {
		fmt.Printf("\nUserID: %dUsername: %sPassword: %d\n", user.id, user.username, user.password)
	}
}
