package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcom := "This is Aaush"
	fmt.Println(welcom)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for corenetto")
	// comma ok \\ err err
	// input , err:=
	// or
	//input, err if you care about error and what iinput is then type tthis
	//_,err if not care about input then write like this

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T", input) //-> this will come as string instead of int
}
