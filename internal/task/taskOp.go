package taskOp

import (
	"fmt"
	"time"
)

// tasks info
type Task struct {
	id           int    //task 1,2,3,4,5 ...
	todo         string //what to do
	state        bool   //completed or not
	timestampReg string //timestamp of registration
	timestampCom string //timestamp of completion
}

func AddTask(todoList *[]Task, todo string) {
	newTask := Task{}
	newTask.id = len(*todoList) + 1
	newTask.todo = todo
	newTask.state = false
	newTask.timestampReg = time.Now().String()
	newTask.timestampCom = ""

	*todoList = append(*todoList, newTask)
}

func DoneTask(todoList *[]Task, taskID int) {
	(*todoList)[taskID].state = true
}

func PrintTodoList(todoList *[]Task) {
	for _, task := range *todoList {
		fmt.Printf("Task id: %dTask todo: %sTask state: %tTask timestampReg: %s Task timestampCom: %s", task.id, task.todo, task.state, task.timestampReg, task.timestampCom)
	}

}
