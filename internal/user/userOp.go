package userOp

import (
	"fmt"
	"log"

	taskOp "github.com/epousa/todoListProject/internal/task"
	"golang.org/x/crypto/bcrypt"
)

// users info
type User struct {
	Id       int
	Username string
	Password []byte
	TodoList []taskOp.Task
}

func AddUser(users *[]User, username string, password []byte) {
	newUser := User{}

	hash, err := bcrypt.GenerateFromPassword(password, 15)
	if err != nil {
		log.Fatal(err)
	}

	newUser.Id = len(*users)
	newUser.Username = username
	newUser.Password = hash

	*users = append(*users, newUser)
}

func FindUser(users []User, id int) int {
	//find user in the list of structs and return its index
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			return i
		}
	}
	return -1
}

func FindUserByUserNamePassword(users []User, username string, password []byte) (bool, int) {
	//find user in the list of structs and return its index
	for i := 0; i < len(users); i++ {
		if users[i].Username == username && bcrypt.CompareHashAndPassword(users[i].Password, password) == nil {
			return true, i
		}
	}
	return false, -1
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
		fmt.Println("User Info:\n", user)
	}
}
