package main

import (
	taskOp "github.com/epousa/todoListProject/internal/task"
	userOp "github.com/epousa/todoListProject/internal/user"
	"github.com/epousa/todoListProject/internal/view"
)

// Array of users
var users []userOp.User

func main() {

	if len(users) == 0 {
		username, password := view.InitView()
		userOp.AddUser(&users, username, password)
		view.CleanStdout()
	}

	taskOp.AddTask(&users[0].TodoList, "workout")
	taskOp.DoneTask(&users[0].TodoList, 0)
	// taskOp.AddTask(&users[1].TodoList, "workout")
	// taskOp.DoneTask(&users[1].TodoList, 0)
	// fmt.Println(userOp.FindUser(users, 0))
	// fmt.Println(userOp.FindUser(users, 1))

	userOp.PrintUsers(users)
	// userOp.RemoveUser(&users, 1)
	// fmt.Println(userOp.FindUser(users, 1))
	// userOp.PrintUsers(users)

}
