package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Initializing....")

	db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/details")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("DataBase is Connected..")

	insert, err := db.Query("Insert into employeedetails values(1,'Ritwik','CSE')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	fmt.Println("Successfully inserted")

}
