package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET") // .Methods("GET") indicates that it only accepts GET request

	// http.ListenAndServe can through error and if that happens then it will be displayed by log.fatal
	log.Fatal(http.ListenAndServe(":4000",r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

// we will take the request in r and write the response using the w
func serveHome(w http.ResponseWriter, r *http.Request) {
	// we get the data as byte from the server hence we are writing in byte only
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}
