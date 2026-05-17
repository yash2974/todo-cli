package services

import (
    "fmt"
    "time"
    "github.com/xuri/excelize/v2"
)

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

func AddTaskExcel(taskTitle string) bool{
    f, err := excelize.OpenFile("Book1.xlsx")
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
    createdOn := time.Now().Format("02-01-2006 15:04:05")
    f.SetCellValue("Sheet1", fmt.Sprintf("A%d", nextRow), nextRow-1)
    f.SetCellValue("Sheet1", fmt.Sprintf("B%d", nextRow), taskTitle)
    f.SetCellValue("Sheet1", fmt.Sprintf("C%d", nextRow), false)
    f.SetCellValue("Sheet1", fmt.Sprintf("D%d", nextRow), createdOn)
    err = f.Save()
	if err != nil {
		fmt.Println("Error saving file")
		return false
	}

	return true
}

func ReadExcel() {
	f, err := excelize.OpenFile("Book1.xlsx")
    if err != nil {
        fmt.Println(err)
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
    fmt.Printf("%-5s %-25s %-10s %-25s\n", "ID", "Title", "Done", "Created_On")

    for i, row := range rows {
        if i == 0 {
            continue
        }

        fmt.Printf(
            "%-5s %-25s %-10s %-25s\n",
            row[0],
            row[1],
            row[2],
            row[3],
        )
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
    // index, err := f.NewSheet("Sheet2")
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // Set value of a cell.
    f.SetCellValue("Sheet1", "A1", "ID")
    f.SetCellValue("Sheet1", "B1", "Title")
    f.SetCellValue("Sheet1", "C1", "Done")
    f.SetCellValue("Sheet1", "D1", "Created_On")
    // Set active sheet of the workbook.
    // f.SetActiveSheet(index)
    // Save spreadsheet by thex given path.
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        fmt.Println(err)
    }
}

func ExcelTaskDetail(taskId string) []string {
    f, err := excelize.OpenFile("Book1.xlsx")
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