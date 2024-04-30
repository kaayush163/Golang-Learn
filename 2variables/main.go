package main

import "fmt"

// token :=40000 //globally cant do like this
func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	// if put 266- then it will show error as unint8 only read value from 0-255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4555453444444
	//this will show 5 decimals only after the point
	//if do float64 it will take more integer in output
	fmt.Println(smallFloat)
	fmt.Printf("Varibale is of type: %T \n", smallFloat)

	var anyValue int
	fmt.Println(anyValue)
	//It will not print any Garbage value but it prints 0
	fmt.Printf("Variable is of type: %T \n", anyValue)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	//website = 3 not able toc hange as it's already declared string

	//no VAR style anymore
	//but this style only be allowed inside any method, function not globally we can do like this
	numberOfUser := 3000
	fmt.Println(numberOfUser)
}
