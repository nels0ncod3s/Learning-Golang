package main 

import (
	"fmt"
)

var userTask string
var userAns int
var taskID int = 0 + 1
var Tasks = map[int]string{

}

func addTask() {
	fmt.Print("ENTER A TASK:")
	fmt.Scanln(&userTask)
	Tasks[taskID] = userTask
	taskID++
}

func deleteTask() {
	//
}

func listTasks() {
	fmt.Println(Tasks)
}

func main (){

	for i := 1; i > 0; i++{
		fmt.Println("1- Add Task, 2 - List Tasks, 3 - End")
		fmt.Scanln(&userAns)
		if userAns == 1{
			addTask()
		} else if userAns == 2{
			listTasks()
		} else if userAns == 3{
			break
		} else {
			fmt.Println("Invalid Command!")
		}
	}
}