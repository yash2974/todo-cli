package services

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
	"strconv"
	"path/filepath"
	"os"
)

func getExcelFilePath() string {
	// err := godotenv.Load()
	// fmt.Println("Load err:", err)
	filePath := "C:\\Users\\yashb\\Documents\\Book1.xlsx"
	return filePath
}

func getBackupExcelFilePath() string {
	filePath := "C:\\Users\\yashb\\Documents\\todocli-backups\\Backup.xlsx"
	return filePath
}

func RowsGetter(f *excelize.File) [][]string {
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("Error")
		return nil
	}

	return rows
}

func IsExcelAvailable() bool {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println("File not availble creating new one.")
		return false
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return true
}

func AddTaskExcel(taskTitle string, tag string) bool {
	filePath := getExcelFilePath()
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("Some error occured, contact dev!")
		return false
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows := RowsGetter(f)
	nextRow := len(rows) + 1
	value, err := f.GetCellValue("Sheet2", "B1")
	if err != nil {
		fmt.Println("Error reading cell:", err)
		return false
	}
	num, _ := strconv.Atoi(value)
	newId := num + 1
	tagValue := "nil"
	if tag != "" {
		tagValue = tag
	}
	createdOn := time.Now().Format("02-01-2006 15:04:05")
	f.SetCellValue("Sheet1", fmt.Sprintf("A%d", nextRow), newId)
	f.SetCellValue("Sheet1", fmt.Sprintf("B%d", nextRow), taskTitle)
	f.SetCellValue("Sheet1", fmt.Sprintf("C%d", nextRow), "false")
	f.SetCellValue("Sheet1", fmt.Sprintf("D%d", nextRow), createdOn)
	f.SetCellValue("Sheet1", fmt.Sprintf("E%d", nextRow), "nil")
	f.SetCellValue("Sheet1", fmt.Sprintf("F%d", nextRow), tagValue)
	f.SetCellValue("Sheet2", "B1", newId)
	err = f.Save()
	if err != nil {
		fmt.Println("Error saving file")
		return false
	}
	fmt.Println("Task saved successfully with task ID -",newId )
	return true
}

func ReadExcel(count int, done string) {
	filePath := getExcelFilePath()
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("create a new one?")
		var choice string
		fmt.Scanln(&choice)
		if choice == "yes" {
			CreateExcel()
			return 
		}
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	printed := 0
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if row[2] != done {
			continue
		}
		fmt.Printf("Task ID: %s\n", row[0])
		fmt.Printf("Task: %s\n", row[1])
		fmt.Printf("Completed: %s\n", row[2])
		fmt.Printf("Created At: %s\n", row[3])
		fmt.Printf("Completed At: %s\n", row[4])
		fmt.Printf("Tag: %s\n\n", row[5])
		printed++
		if count > 0 && printed >= count {
			break
		}
	}
}

func CreateExcel() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// // Create a new sheet.
	f.NewSheet("Sheet2")
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Title")
	f.SetCellValue("Sheet1", "C1", "Done")
	f.SetCellValue("Sheet1", "D1", "Created_On")
	f.SetCellValue("Sheet2", "A1", "nextId")
	f.SetCellValue("Sheet2", "B1", 0)
	f.SetCellValue("Sheet2", "A2", "faps")
	f.SetCellValue("Sheet2", "B2", 0)
	// Set active sheet of the workbook.
	// f.SetActiveSheet(index)
	// Save spreadsheet by thex given path.
	filePath := getExcelFilePath()
	
	// Create the data directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	
	if err := f.SaveAs(filePath); err != nil {
        fmt.Println(err)
    }
}

func ExcelTaskDetail(taskId string) []string {
	filePath := getExcelFilePath()
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}

		if row[0] == taskId {
			return row
		}
	}

	return nil
}


//TODO: understand - YASH
func DeleteTask(flag string) {
	
    if flag == "0" {
        fmt.Printf("Error!")
    }
	filePath := getExcelFilePath()
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	if flag == "all" {
		fmt.Println("This will delete all tasks - yes or no")
		var option string
		fmt.Scanln(&option)
		count := 0
		if option == "yes" {
			for rowIndex := len(rows) - 1; rowIndex >= 1; rowIndex-- {

				if len(rows[rowIndex]) == 0 {
					continue
				}

				err = f.RemoveRow("Sheet1", rowIndex+1)
				if err != nil {
					fmt.Println(err)
					return
				}

				count++
			}

			if err := f.Save(); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%d tasks deleted successfully\n", count)
			return
		} else {
			return
		}
	}
	//iterate through rows and then find the row number of the task id
	rowNumber := -1
	for rowIndex, row := range rows {

		// skip empty rows
		if len(row) == 0 {
			continue
		}

		if row[0] == flag {
			// Excel rows start from 1
			rowNumber = rowIndex + 1
			break
		}
	}
	if rowNumber == -1 {
		fmt.Println("Task not found")
		return
	}
	err = f.RemoveRow("Sheet1", rowNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Save changes
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Task deleted successfully")
}

func EditTaskExcel(task_id string, done string, taskTitle string, taskTag string) {
	filePath := getExcelFilePath()
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err) 
		return 
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return 
	}
	taskRowIndex := -1
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if row[0] == task_id {
			taskRowIndex = i
		}
	}
	if taskRowIndex == -1 {
		fmt.Println("Task not found!")
		return
	}
	
	if done != "" {
		finishedOn := time.Now().Format("02-01-2006 15:04:05")
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", taskRowIndex+1), done)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", taskRowIndex+1), finishedOn)
	}

	if taskTitle != "" {
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", taskRowIndex+1), taskTitle)
	}

	if taskTag != "" {
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", taskRowIndex+1), taskTag)
	}

	err = f.Save()
	if err != nil {
		fmt.Println("Error saving file")
		return 
	}

	fmt.Println("Task updated Successfully!")
	
}