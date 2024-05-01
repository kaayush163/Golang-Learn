package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello mod in golang")
	greeter()
	r := mux.NewRouter() //this adds github.com/gorilla in mport
	// r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/", serveHome).Methods("GET")
	// http.ListenAndServe(":4000",r) //can throw some error
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hey There mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series</h1>"))
}
