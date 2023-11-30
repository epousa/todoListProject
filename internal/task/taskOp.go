package taskOp

import (
	"fmt"
	"time"
)

// tasks info
type Task struct {
	Id           int    //task 1,2,3,4,5 ...
	Todo         string //what to do
	State        bool   //completed or not
	TimestampReg string //timestamp of registration
	TimestampCom string //timestamp of completion
}

func AddTask(todoList *[]Task, todo string) {
	newTask := Task{}
	newTask.Id = len(*todoList)
	newTask.Todo = todo
	newTask.State = false
	newTask.TimestampReg = time.Now().String()
	newTask.TimestampCom = ""

	*todoList = append(*todoList, newTask)
	fmt.Println("Task was added successfully !")
}

func FindTask(todoList []Task, id int) int {
	//find user in the list of structs and return its index
	for i := 0; i < len(todoList); i++ {
		if todoList[i].Id == id {
			return i
		}
	}
	return -1
}

func RemoveTask(todoList *[]Task, id int) int {
	taskIndex := FindTask(*todoList, id)
	if taskIndex == -1 {
		return -1
	}
	*todoList = append((*todoList)[:taskIndex], (*todoList)[taskIndex+1:]...)
	return 0
}

func DoneTask(todoList *[]Task, taskID int) {
	(*todoList)[taskID].State = true
	(*todoList)[taskID].TimestampCom = time.Now().String()
}

func PrintTodoList(todoList *[]Task) {
	for _, task := range *todoList {
		fmt.Printf("Task id: %dTask todo: %sTask state: %tTask timestampReg: %s Task timestampCom: %s", task.Id, task.Todo, task.State, task.TimestampReg, task.TimestampCom)
	}

}
