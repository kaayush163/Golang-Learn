package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for corenetto")
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)
	// numRate, err := strconv.ParseFloat(input, 64) //the alone input will take \n also so we hav to use fucntion of go
	numRate, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //the alone input will take \n also so we hav to use fucntion of go

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRate+1)
	}
}
