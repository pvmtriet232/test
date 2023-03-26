package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establish database connection
	db, err := sql.Open("mysql", "web:pass@/snippetbox?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Execute SELECT query
	rows, err := db.Query("SELECT * FROM snippets")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate over the result set and display the rows
	for rows.Next() {
		var column1 string
		var column2 string
		var column3 string
		var column4 string
		var column5 string
		err := rows.Scan(&column1, &column2, &column3, &column4, &column5)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(column1, column2, column3, column4, column5)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
}
