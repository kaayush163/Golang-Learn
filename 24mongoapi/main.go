package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kaayush163/Golang-Learn/blob/main/24mongoapi/router"
)

func main() {

	fmt.Println("Mongo databse ")

	r := router.Router()
	fmt.Println("Serve si geting started")

	// http.ListenAndServe((":4000",))

	//
	log.Fatal(http.ListenAndServe(":4000", r))
}
