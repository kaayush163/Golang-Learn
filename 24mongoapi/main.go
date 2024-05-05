package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Mongo databse ")

	r := router.Router()
	fmt.Println("Serve si geting started")

	// http.ListenAndServe((":4000",))

	//
	log.Fatal(http.ListenAndServe(":4000", r))
}
