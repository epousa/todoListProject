package main

import (
	"fmt"

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

	fmt.Println(userOp.FindUser(users, 1))
	userOp.PrintUsers(users)
	userOp.RemoveUser(&users, 1)
	fmt.Println(userOp.FindUser(users, 1))
	userOp.PrintUsers(users)

	taskOp.AddTask(&todoList, "Do a workout")
	fmt.Println(taskOp.FindTask(todoList, 1))
	taskOp.PrintTodoList(&todoList)
	taskOp.RemoveTask(&todoList, 1)
	fmt.Println(taskOp.FindTask(todoList, 1))
	taskOp.PrintTodoList(&todoList)
}
