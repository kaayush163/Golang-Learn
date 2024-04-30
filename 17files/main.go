package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golan")

	content := "This needs to go in file"

	file, err := os.Create("./mylocalsave.txt")

	checkNilErr(err)

	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Length", length)
	defer file.Close() //this recommentd to put defer because since user need to write some code afetr this so after exceuting thise code then this should run f.close
	readFile("./mylocalsave.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Text data inside the fiel \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
