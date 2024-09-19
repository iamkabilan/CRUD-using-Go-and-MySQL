package controllers

import (
	"fmt"
	"net/http"

	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/database"
	"github.com/xuri/excelize/v2"
)

func ExportUsers(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	db := database.GetDB()

	rows, queryErr := db.Query("SELECT id, name, email FROM users")
	if queryErr != nil {
		fmt.Println("ERROR (Query Error): ", queryErr)
		return
	}
	defer rows.Close()

	var file = excelize.NewFile();
	var index, sheetErr = file.NewSheet("users");
	file.DeleteSheet("Sheet1");
	if sheetErr != nil {
		fmt.Println("ERROR (Sheet Error): ", sheetErr)
	}

	var sheetName = "Users"
	file.SetCellValue(sheetName, "A1", "Id")
	file.SetCellValue(sheetName, "B1", "Name")
	file.SetCellValue(sheetName, "C1", "Email")

	sheetRowNumber := 2

	for rows.Next() {
		var user User
		scanErr := rows.Scan(&user.Id, &user.Name, &user.Email)
		if scanErr != nil {
			fmt.Println("ERROR: ", scanErr)
		}

		file.SetCellValue(sheetName, fmt.Sprintf("A%d", sheetRowNumber), user.Id)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", sheetRowNumber), user.Name)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", sheetRowNumber), user.Email)
		sheetRowNumber++
	}
	file.SetActiveSheet(index)

	responseWriter.Header().Set("Content-Disposition", "attachment;filename=users.xlsx")
	if err := file.Write(responseWriter); err != nil {
		fmt.Println("ERROR (Write Error): ", err)
		return
	}
}
