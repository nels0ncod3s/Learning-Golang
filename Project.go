package main 

import (
	"fmt"
)

var userTask string
var userAns int

Taskss := map[int] string{
	1 : "Do dishes",
}

func addTask(task string) {
	fmt.Print("ENTER A TASK:")
	fmt.Scanln(&userTask)
}

func deleteTask(task string) {
	//
}

func listTasks() {
	fmt.Println(Taskss)
}

func main (){

	for i := 1; i > 0; i++{
		fmt.Println("1- Add Task, 2 - End")
		fmt.Scanln(&userAns)
		if userAns == 1{
			addTask()
		} else if userAns == 2{
			break
		} else {
			fmt.Println("Invalid Command!")
		}
	}
}