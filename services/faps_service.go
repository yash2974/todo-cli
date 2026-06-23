package services

import (
	"fmt"
	"strconv"
	"github.com/xuri/excelize/v2"
)

func Faps(reset bool, increment bool) {
	filePath := getExcelFilePath()
	f, err := excelize.OpenFile(filePath)
	currentStreak, err := f.GetCellValue("Sheet2", "B2")

	if err != nil {
		fmt.Println("Error reading cell:", err)
		return 
	}



	num, _ := strconv.Atoi(currentStreak)
	if !increment && !reset{
		fmt.Println("Current Streak:", num )
		return
	}
	newStreakCount := num + 1
	if reset == true {
		newStreakCount = 0
	}
	


	f.SetCellValue("Sheet2", "B2", strconv.Itoa(newStreakCount))
	err = f.Save()
	if err != nil {
		fmt.Println("Error saving file")
		return 
	}

	fmt.Println("Streak updated:", newStreakCount)
}