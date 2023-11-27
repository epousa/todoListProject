package main

import (
	"fmt"

	userOp "github.com/epousa/todoListProject/internal/user"
)

// tasks info
type Task struct {
	id           int    //task 1,2,3,4,5 ...
	todo         string //what to do
	state        bool   //completed or not
	timestampReg string //timestamp of registration
	timestampCom string //timestamp of completion
}

// Array of users
var users []userOp.User

// Array of tasks
var todoList []Task

// Each user corresponds an array of tasks
// key -> user password hash / value -> todoList
var database map[string][]Task

func main() {
	fmt.Println("My todoListProject !")
	users = userOp.AddUser(users, "eduardo", "somethin")
	userOp.ListUsers(users)
}
