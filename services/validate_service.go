package services

import (
	"errors"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func Validate(check string) error {
	if check == "" {
		return nil
	}
	filePath := getExcelFilePath()
	backupFilePath := getBackupExcelFilePath()

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("Some error occured, contact dev!")
		return err
	}
	err = f.SaveAs(backupFilePath)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	} else {
		fmt.Println("Backup Saved at ", backupFilePath)
	}
	f.Path = filePath
	switch check {
	case "data":
		return validateData(f)
	case "fields":
		return validateFields(f)
	}
	return nil
}

func validateData(f *excelize.File) error {

	//first check for tasks (Sheet1)
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return errors.New("")
	}
	updated := 0
	for i := range rows {
		if i == 0 {
			continue
		}
		
		for ch := 'A'; ch <= 'F'; ch++ {
			cellIndex := fmt.Sprintf("%c%d", ch, i+1)
			val, err := f.GetCellValue("Sheet1", cellIndex)
			if err != nil {
				f.SetCellValue("Sheet1", cellIndex, "nil")
				updated += 1
			}
			if val == "" {
				f.SetCellValue("Sheet1", cellIndex, "nil")
				updated += 1
			}
		}
	}

	// Save changes
	if updated > 0 {
		if err := f.Save(); err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

// if validationErrorCount > 0, then there is a field missing somewhere - repopulate everthing for data integrity
func validateFields(f *excelize.File) error {
	// first check for tasks (Sheet1) - WILL ADD SHEET2 CONTEXT LATER
	intended := []string{"ID", "Title", "Done", "Created_On", "Finished_On", "Tag"}
	var output []string
	validationErrorCount := 0
	i := 0
	for ch := 'A'; ch <= 'F'; ch++ {
		cellIndex := fmt.Sprintf("%c1", ch)
		value, _ := f.GetCellValue("Sheet1", cellIndex)
		if value != intended[i] {
			validationErrorCount += 1
		}
		output = append(output, value)
		i += 1
	}
	if !slices.Equal(output, intended) {
		fmt.Println("Tasks Fields Validation Error! Running Clean up and reinitialization...")
	} else {
		fmt.Println("Tasks Validated Sucessfully!")
	}
	if validationErrorCount > 0 {
		// run function repopulate fields
		// TODO: make this func
		// will send value as a parameter and then return the new value of errorCount
		validationErrorCount = repopulateFieldsTasks(f, validationErrorCount)
		if validationErrorCount > 0 {
			fmt.Println("Error Validation - Tasks Sheets")
			return errors.New("validation failed: missing header row for tasks sheet")
		}
	}

	//second - check faps
	intendedFap := "faps"
	// validationErrorCount should be 0 at this point - assumption
	value, _ := f.GetCellValue("Sheet2", "A3")
	if value != intendedFap {
		validationErrorCount += 1
	}
	if validationErrorCount > 0 {
		// run function to repopulate fields
		fmt.Println("Faps Validation Error! Running Clean up and reinitialization...")
		validationErrorCount = repopulateFieldsFaps(f, validationErrorCount)
		if validationErrorCount > 0 {
			fmt.Println("Error Validation - Maintainer Sheets")
			return errors.New("validation failed: missing header row for maintainer sheet")
		}
	} else {
		fmt.Println("Faps Validated Sucessfully!")
	}

	//third - check gym
	intendedGym := []string{"Date", "Description"}
	var outputGym []string
	i = 0
	for ch := 'A'; ch <= 'B'; ch++ {
		cellIndex := fmt.Sprintf("%c1", ch)
		value, _ := f.GetCellValue("Sheet3", cellIndex)
		if value != intendedGym[i] {
			validationErrorCount += 1
		}
		outputGym = append(outputGym, value)
		i += 1
	}
	if !slices.Equal(outputGym, intendedGym) {
		fmt.Println("Gym Fields Validation Error! Running Clean up and reinitialization...")
	} else {
		fmt.Println("Gym Validated Sucessfully!")
	}
	if validationErrorCount > 0 {
		// run function repopulate fields
		// TODO: make this func
		// will send value as a parameter and then return the new value of errorCount
		validationErrorCount = repopulateFieldsGym(f, validationErrorCount)
		if validationErrorCount > 0 {
			fmt.Println("Error Validation - Gym Sheets")
			return errors.New("validation failed: missing header row for gym sheet")
		}
	}

	return nil
}

func repopulateFieldsFaps(f *excelize.File, errorCount int) int {
	f.SetCellValue("Sheet2", "A3", "faps")
	f.SetCellValue("Sheet2", "B3", 0)
	err := f.Save()
	if err != nil {
		fmt.Println("Error Saving File")
		return 1
	} else {
		fmt.Println("Faps reininitiated successfully")
	}
	return 0
}

func repopulateFieldsTasks(f *excelize.File, errorCount int) int {
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Title")
	f.SetCellValue("Sheet1", "C1", "Done")
	f.SetCellValue("Sheet1", "D1", "Created_On")
	f.SetCellValue("Sheet1", "E1", "Finished_On")
	f.SetCellValue("Sheet1", "F1", "Tag")

	err := f.Save()
	if err != nil {
		fmt.Println("Error Saving File")
		return 6
	} else {
		fmt.Println("Tasks Fields reininitiated successfully")
	}
	return 0

}

func repopulateFieldsGym(f *excelize.File, errorCount int) int {
	f.SetCellValue("Sheet3", "A1", "Date")
	f.SetCellValue("Sheet3", "B1", "Description")
	err := f.Save()
	if err != nil {
		fmt.Println("Error Saving File")
		return 6
	} else {
		fmt.Println("Gym Fields reininitiated successfully")
	}
	return 0
}
