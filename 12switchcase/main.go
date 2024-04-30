package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano()) //everytime it changes random value

	diceNumber := rand.Intn(6) + 1 //this Intn will generate  umber from 0 to 5 non inclusive 6 that's why we do +1
	fmt.Println("valie", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice is 1")

	case 2:
		fmt.Println("You can mov eto 2")

	case 3:
		fmt.Println("You can 3")
		fallthrough //this allow to exceute case 4 also and only that one only

	case 4:
		fmt.Println("You can 3")
		fallthrough //this allow to exceute case 5 also and only that one only

	case 5:
		fmt.Println("You can 3")

	case 6:
		fmt.Println("You can 3")

	default:
		fmt.Println("What si the value given")

	}
}
