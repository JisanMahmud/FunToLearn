package gorrilaRouter

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/GoRockypt/FunToLearn/GoSupport"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//HomeHandler H
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Home")
	fmt.Fprint(w, "Hello")
}

//HandleUserPass C
func HandleUserPass(w http.ResponseWriter, r *http.Request) {
	log.Print("Other")
	vars := mux.Vars(r)
	UserName := vars["username"]
	PassWord := vars["password"]

	//import "crypto/sha256"
	h := sha256.New()
	io.WriteString(h, PassWord)

	//Encode the password to hex to save in the DB
	encodedStr := hex.EncodeToString(h.Sum(nil))
	log.Print(encodedStr)

	//Let's see if the password matches
	HashedPass := godb.DBQueryUserName(UserName)

	if strings.Compare(encodedStr, HashedPass) == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

//AnyPostMethod C
func AnyPostMethod(w http.ResponseWriter, r *http.Request) {
	log.Print("Other")
	vars := mux.Vars(r)
	Any := vars["any"]
	fmt.Fprint(w, Any+" is not recognised")
}

//AnyGetMethod C
func AnyGetMethod(w http.ResponseWriter, r *http.Request) {
	log.Print("Other")
	vars := mux.Vars(r)
	Any := vars["any"]
	fmt.Fprint(w, Any+" is not recognised")
}

//ReportHandler C
func ReportHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Report")
	vars := mux.Vars(r)
	ReportID := vars["ReportID"]
	if ReportID != "" {
		log.Print("Reports " + ReportID)
	}

	table := godb.DBGetActorReport()

	fmt.Fprint(w,table)
}

//StartServer  This starts the server
func StartServer() {
	print("Server started")
	mux := mux.NewRouter()
	mux.Host("127.0.0.1")
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/GetReport/{ReportID}", ReportHandler).Methods("Get")
	mux.HandleFunc("/user/{username}/pass/{password}", HandleUserPass).Methods("POST")
	mux.HandleFunc("/{Any}", AnyGetMethod).Methods("Get")
	mux.HandleFunc("/{Any}", AnyPostMethod).Methods("POST")
	http.Handle("/", mux)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
