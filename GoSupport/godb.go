package godb

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/GoRockypt/FunToLearn/goAPI"
	_ "github.com/go-sql-driver/mysql" //Driver
)

//UserInfo struct to connect user info DB
type UserInfo struct {
	UserID, Hash string
}

//Actors tem db
type Actors struct {
	actorID                         int
	firstName, lastName, lastUpdate string
}

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

//DBQueryUserName Check if userName checks out
func DBQueryUserName(UserName string) string {

	db, err := sql.Open("mysql", "admin:A1s1D2f2@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var data UserInfo
	err = db.QueryRow("SELECT * FROM user_data.userinfo WHERE UserID = ?", UserName).Scan(&data.UserID, &data.Hash)

	log.Print(data.Hash)

	return data.Hash
}

//DBGetActorReport Check if userName checks out
func DBGetActorReport() string {

	db, err := sql.Open("mysql", "admin:A1s1D2f2@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	rows, err := db.Query("SELECT * FROM sakila.actor")

	var data Actors
	var JSONTable string

	for rows.Next() {
		if err := rows.Scan(&data.actorID, &data.firstName, &data.lastName, &data.lastUpdate); err != nil {
			log.Fatal(err)
			return "Failed"
		}

		var JSONRow string
		JSONRow += goapi.AddJSONFieldAndData("Field1", strconv.Itoa(data.actorID)) + ","
		JSONRow += goapi.AddJSONFieldAndData("Field2", data.firstName) + ","
		JSONRow += goapi.AddJSONFieldAndData("Field3", data.lastName) + ","
		JSONRow += goapi.AddJSONFieldAndData("Field4", data.lastUpdate)

		JSONTable += goapi.SurroundWithChars("{", "}", JSONRow) + ","
	}

	JSONTable = strings.TrimRight(JSONTable, ",")
	JSONTable = goapi.SurroundWithQuote("Rows") + ":" + goapi.SurroundWithChars("[", "]", JSONTable)
	JSONTable = goapi.SurroundWithChars("{", "}", JSONTable)

	return JSONTable
}
