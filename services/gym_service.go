package services

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
	"slices"
	"strconv"
)

func TagGym(log int, day string) {
	filePath := getExcelFilePath()
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	validation := validateSheet(f)
	if !validation {
		// create a new sheet - Sheet3 and add date and what day of gym was it
		sheetCreated := createGymSheet(f)
		if !sheetCreated {
			fmt.Printf("error - closing")
			return
		}
		fmt.Printf("Gym Sheet Created Successfully!")
	}
	currentLog, err := f.GetCellValue("Sheet2", "B2")
	if err != nil {
		fmt.Printf("error contact dev!")
		return
	}
	logIdx, err := strconv.Atoi(currentLog)
	if err != nil {
		fmt.Printf("error parsing gym index!")
		return
	}
	method := "log"
	if day != "" {
		method = "day"
	}
	switch method {
	case "log":
		gymlogs(f, log)
	case "day":
		gymlogger(f, day, logIdx)
	default: 
		gymlogs(f, log)
	}
	return
}

//function for gym logs
func gymlogs(f *excelize.File, log int) {
	logs, err := f.GetRows("Sheet3")
	if err != nil {
		fmt.Printf("Error in reading!")
		return
	}
	for i, log := range slices.Backward(logs) {
		if i == 0 {
			continue
		}
		fmt.Printf("%s %s \n", log[0], log[1])
	}
}

//function for gym logger
func gymlogger(f *excelize.File, day string, currentLog int) {
	today := time.Now().Format("02-01-2006")
	nextLog := currentLog + 1
	f.SetCellValue("Sheet3", fmt.Sprintf("A%d", nextLog), today)
	f.SetCellValue("Sheet3", fmt.Sprintf("B%d", nextLog), day)
	f.SetCellValue("Sheet2", "B2", nextLog)
	err := f.Save()
	if err != nil {
		fmt.Printf("error saving file")
	}
	return
}

// returns boolean based on gym validation is correct or not
func validateSheet(f *excelize.File) bool{
	sheetName := "Sheet3"
	index, _ := f.GetSheetIndex(sheetName)
	if index != -1 {
		date, _ := f.GetCellValue("Sheet3", "A1")
		day, _ := f.GetCellValue("Sheet3", "B1")
		if date != "Date" && day != "Description" {
			f.SetCellValue("Sheet3", "A1", "Date")
			f.SetCellValue("Sheet3", "B1", "Description")
			if err := f.Save(); err != nil {
				fmt.Println(err)
				return false
			}
		}
	}
	return index != -1
}
//creates gym sheet
func createGymSheet(f *excelize.File) bool{
	f.NewSheet("Sheet3")
	f.SetCellValue("Sheet3", "A1", "Date")
	f.SetCellValue("Sheet3", "B1", "Description")
	f.SetCellValue("Sheet2", "A2", "gym_idx")
	f.SetCellValue("Sheet2", "B2", 1)
	if err := f.Save(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}