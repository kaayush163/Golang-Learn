package main

import "fmt"

func main() {
	defer fmt.Println("aayush")
	defer fmt.Println("shanu")
	defer fmt.Println("RDJ")
	fmt.Println("heelo")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
