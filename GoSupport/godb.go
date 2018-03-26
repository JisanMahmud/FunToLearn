package godb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Driver
)

//DbOpen function to open a connection
func DbOpen() {

	db, err := sql.Open("mysql", "admin:A1s1D2f2@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO squarenum.new_table VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		fmt.Println(err) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum.new_table WHERE number = ?")
	if err != nil {
		fmt.Println(err) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
		if err != nil {
			fmt.Println(err) // proper error handling instead of panic in your app
		}
	}

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		fmt.Println(err) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		fmt.Println(err) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)
}
