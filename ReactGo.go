package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" //Driver
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Databasetest afdsa
func create() {

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

//HomeHandler H
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Home")
	fmt.Fprint(w, "Hello")
}

//CopyHandle C
func CopyHandle(w http.ResponseWriter, r *http.Request) {
	log.Print("Other")
	vars := mux.Vars(r)
	title := vars["Copy"]
	fmt.Fprint(w, title)
}

func main() {
	create()

	print("Server started")
	mux := mux.NewRouter()
	mux.Host("127.0.0.1")
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/{Copy}", CopyHandle)
	http.Handle("/", mux)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
