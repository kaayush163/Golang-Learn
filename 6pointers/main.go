package main

import "fmt"

func main() {
	fmt.Println("Wlecome to a class on pointer")

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr * 3

	fmt.Println("New value is ", myNumber) //the changes seen on the original number
}
