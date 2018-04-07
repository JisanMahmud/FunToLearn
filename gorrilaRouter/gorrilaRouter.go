package gorrilaRouter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Fprint(w, UserName)
	fmt.Fprintln(w, PassWord)

	//hasing
	bytes, err := bcrypt.GenerateFromPassword([]byte(PassWord), 14)
	if err == nil {
		log.Println("hash not done")
	}

	fmt.Fprintln(w, string(bytes))
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

//StartServer  This starts the server
func StartServer() {
	print("Server started")
	mux := mux.NewRouter()
	mux.Host("127.0.0.1")
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/user/{username}/pass/{password}", HandleUserPass).Methods("POST")
	mux.HandleFunc("/{Any}", AnyGetMethod).Methods("Get")
	mux.HandleFunc("/{Any}", AnyPostMethod).Methods("POST")
	http.Handle("/", mux)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
