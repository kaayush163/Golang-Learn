package main

import "fmt"

func main() {
	fmt.Println("srray learn")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Tomato"
	fruits[2] = "Peach"

	fmt.Println("Fruit list is", fruits)
	fmt.Println("Fruit list is", len(fruits))

	var vegList = [4]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegetable", vegList)
	fmt.Println("Vegetable", len(vegList))

}
