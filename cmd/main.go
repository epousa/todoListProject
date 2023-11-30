package main

import (
	"fmt"

	taskOp "github.com/epousa/todoListProject/internal/task"
	userOp "github.com/epousa/todoListProject/internal/user"
	"github.com/epousa/todoListProject/internal/view"
)

// Array of users
var users []userOp.User

func AdminChoices(choice string, userID int) bool {
	// view.CleanStdout()
	switch choice {
	case "Add Task":
		taskOp.AddTask(&users[userID].TodoList, view.AddTaskView())
		break
	case "Remove Task":
		taskOp.DoneTask(&users[userID].TodoList, view.RemoveTaskView())
		break
	case "Check todo List":
		taskOp.PrintTodoList(&users[userID].TodoList)
		break
	case "Access users list":
		userOp.PrintUsers(users)
		break
	case "Add user":
		username, password := view.InitView()
		userOp.AddUser(&users, username, password)
		break
	case "Remove user":
		userOp.RemoveUser(&users, view.RemoveUserView())
		break
	case "Log out":
		return true
	}

	return false
}

func StandardChoices(choice string, userID int) bool {

	switch choice {
	case "Add Task":
		taskOp.AddTask(&users[userID].TodoList, view.AddTaskView())
		break
	case "Log out":
		return true
	}

	return false
}

func Core() {
	username, password := view.InitView()
	exists, userIndex := userOp.FindUserByUserNamePassword(users, username, password)

	if exists {
		fmt.Printf("\n\n### You logged in successfully! :) ###\n\n")
		if users[userIndex].Id == 0 {
			//admin choices
			for AdminChoices(view.AdminView(users[userIndex].Username), userIndex) != true {
			}
		} else {
			//standard users choices
			for StandardChoices(view.StandardView(users[userIndex].Username), userIndex) != true {
			}
		}
	} else {
		userOp.AddUser(&users, username, password)
	}
}

func CommonChoices(choice string) bool {
	switch choice {
	case "Register/Login":
		Core()
		break
	case "Exit":
		return true
	}

	return false
}

func main() {

	if len(users) == 0 {
		fmt.Println("Welcome! Create an admin account to start.")
		username, password := view.InitView()
		userOp.AddUser(&users, username, password)
		for AdminChoices(view.AdminView(users[0].Username), 0) != true {
		}
	}

	for CommonChoices(view.CommonView()) != true {
	}

	// taskOp.AddTask(&users[0].TodoList, "workout")
	// taskOp.DoneTask(&users[0].TodoList, 0)
	// taskOp.AddTask(&users[1].TodoList, "workout")
	// taskOp.DoneTask(&users[1].TodoList, 0)
	// fmt.Println(userOp.FindUser(users, 0))
	// fmt.Println(userOp.FindUser(users, 1))

	// userOp.PrintUsers(users)
	// userOp.RemoveUser(&users, 1)
	// fmt.Println(userOp.FindUser(users, 1))
	// userOp.PrintUsers(users)

}
