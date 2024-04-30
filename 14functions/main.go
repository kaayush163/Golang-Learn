package main

import "fmt"

func main() {

	result := adder(3, 5)

	fmt.Println("Result is:", result)

	proRes, myMsg := proAdder(2, 5, 8, 7, 6, 8, 99, 67, 65)
	fmt.Println("Total sum", proRes)
	fmt.Println("Total sum", myMsg)

}

func adder(valOne int, valTwo int) int { //these int have to specify also what value is return back these are fuctions signature
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val

	}

	return total, "Hi I am aayush"
}
