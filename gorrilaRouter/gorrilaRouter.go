package gorrilaRouter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

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

func StartServer() {
	print("Server started")
	mux := mux.NewRouter()
	mux.Host("127.0.0.1")
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/{Copy}", CopyHandle)
	http.Handle("/", mux)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
