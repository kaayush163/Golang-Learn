package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("")

	var fruitList = []string{"Apple", "Tomato", "Paech"}
	fmt.Printf("types of fruits is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banna")
	fmt.Println(fruitList)

	//except 1 elemtn(index 0) all other get dispalyed
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	//start from index 1 and go upto 3 index not inclusive of 4
	fruitList = append(fruitList[1:5])
	fmt.Println(fruitList)

	highScore := make([]int, 4) //this is how to create array by := and upper one was with var
	highScore[0] = 234
	highScore[1] = 234
	highScore[2] = 234
	highScore[3] = 234
	// highScore[4] = 278 // crash out of boundary
	highScore = append(highScore, 555, 999, 000) //this added succefully even after having 4 size
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore)) //give true if all value are sorted

	var course = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(course)

	index := 2
	course = append(course[:index], course[index+1:]...) // react,javascript,python,ruby,
	fmt.Println(course)

}
