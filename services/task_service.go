package services

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
	"TODOCLI/models"
)

var tasks []models.Task
var reader = bufio.NewReader(os.Stdin)

func DisplayMenu() {
	fmt.Println("TodoCLI")
	fmt.Println("LIST TASK - 1")
	fmt.Println("ADD TASK - 2")
	fmt.Println("TASK DETAILS - 3")
}

func Addtask() {
	fmt.Println("Enter your task here")
	var taskTitle string
	// fmt.Scanln(&taskTitle)
	taskTitle, _ = reader.ReadString('\n')
	taskTitle = strings.TrimSpace(taskTitle)

	createdOn := time.Now()

	task := models.Task{
		Title:      taskTitle,
		Done:       false,
		Created_On: createdOn.String(),
	}

	tasks = append(tasks, task)
}

func ListTask() {
	
	if len(tasks) > 0 {
		fmt.Println("CURRENT TASKS")
		for i := 0; i < len(tasks); i++ {
			fmt.Printf("%d. %s\n", i+1, tasks[i].Title)
		}
	} else {
		fmt.Println("No tasks found!")
	}
	DisplayMenu()
	
}

func TaskDetail() {
	fmt.Println("Enter task ID")
	var taskId int
	fmt.Scanln(&taskId)
	if (len(tasks)<taskId){
		fmt.Println("No task found")
		DisplayMenu()
	} else {
		fmt.Printf("%s\n", tasks[taskId-1].Title)
		fmt.Printf("%s: %s\n", "Created On", tasks[taskId-1].Created_On)
		fmt.Printf("%s: %t\n", "Completed", tasks[taskId-1].Done)
	}
}