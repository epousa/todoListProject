package main

import (
	"github.com/epousa/todoListProject/internal/menu"
	taskOp "github.com/epousa/todoListProject/internal/task"
	userOp "github.com/epousa/todoListProject/internal/user"
)

// Array of users
var users []userOp.User

// Array of tasks
var todoList []taskOp.Task

// Each user corresponds an array of tasks
// key -> User Position in the array of users / value -> todoList
var database map[int][]taskOp.Task

func main() {
	if len(users) == 0 {
		username, password := menu.InitMenu()
		userOp.AddUser(&users, username, password)
	}

	userOp.PrintUsers(&users)

	taskOp.AddTask(&todoList, "Do a workout")
	taskOp.PrintTodoList(&todoList)
}
