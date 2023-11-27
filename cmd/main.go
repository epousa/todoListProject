package main

import (
	"fmt"

	dbOp "github.com/epousa/todoListProject/internal/db"
	taskOp "github.com/epousa/todoListProject/internal/task"
	userOp "github.com/epousa/todoListProject/internal/user"
)

// users info
type user struct {
	username string
	password string
}

// tasks info
type task struct {
	id           int    //task 1,2,3,4,5 ...
	todo         string //what to do
	state        bool   //completed or not
	timestampReg string //timestamp of registration
	timestampCom string //timestamp of completion
}

// Array of users
var users []user

// Array of tasks
var todoList []task

// Each user corresponds an array of tasks
// key -> user password hash / value -> todoList
var database map[string][]task

func main() {
	fmt.Println("My todoListProject !")
	dbOp.HelloDb()
	taskOp.HelloTask()
	userOp.HelloUser()
}
