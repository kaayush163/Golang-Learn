package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23

	var result string

	if loginCount < 10 { //Not allowed to take this { onto next line this structure of Golang}
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	}

	if num := 3; num < 10 { //assign at the same time and see cocndition at same time
		fmt.Println("Num is less than 10")
	} else {

		fmt.Println(result)
	}
}
