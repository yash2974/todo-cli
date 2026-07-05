package services

import (
	"fmt"
	"bufio"
	"os"
	"TODOCLI/models"
)

var tasks []models.Task
var reader = bufio.NewReader(os.Stdin)

func HelpMenu() {
	fmt.Println(`
████████╗ ██████╗ ██████╗  ██████╗      ██████╗██╗     ██╗
╚══██╔══╝██╔═══██╗██╔══██╗██╔═══██╗    ██╔════╝██║     ██║
   ██║   ██║   ██║██║  ██║██║   ██║    ██║     ██║     ██║
   ██║   ██║   ██║██║  ██║██║   ██║    ██║     ██║     ██║
   ██║   ╚██████╔╝██████╔╝╚██████╔╝    ╚██████╗███████╗██║
   ╚═╝    ╚═════╝ ╚═════╝  ╚═════╝      ╚═════╝╚══════╝╚═╝

            
`)
	fmt.Println(`
================ TODO CLI ================

Commands:

list
    Show all tasks

add
    Create a new task

detail <task_id>
    Show details of a specific task

help
    Show available commands

exit
    Close the application

==========================================
`)
}

// func Addtask() {
// 	fmt.Println("Enter your task here")
// 	var taskTitle string
// 	// fmt.Scanln(&taskTitle)
// 	taskTitle, _ = reader.ReadString('\n')
// 	taskTitle = strings.TrimSpace(taskTitle)


// 	AddTaskExcel(taskTitle)
// }

// func ListTask() {
// 	ReadExcel()
// 	fmt.Printf("\n")
// }

func TaskDetail(param string) {
	var taskId string
	if param == "" {
	fmt.Println("Enter task ID")
		fmt.Scanln(&taskId)
	} else {
		taskId = param
	}
	taskDetailData := ExcelTaskDetail(taskId)

	if taskDetailData == nil {
		fmt.Println("Task not found")
		return
	}

	fmt.Printf("ID: %s\n", taskDetailData[0])
	fmt.Printf("Title: %s\n", taskDetailData[1])
	fmt.Printf("Completed: %s\n", taskDetailData[2])
	fmt.Printf("Created On: %s\n", taskDetailData[3])
}