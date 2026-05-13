package main

import (
	"fmt"
	"TODOCLI/services"
)

func main() {
	services.ReadExcel()

	services.DisplayMenu()
	for {
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			services.ListTask()
		case 2:
			services.Addtask()
		case 3:
			services.TaskDetail()
		}
	}
}




