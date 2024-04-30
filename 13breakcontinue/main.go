package main

import "fmt"

func main() {
	fmt.Println("Welcome to llops")

	days := []string{"sunday", "Tuedays", "Wednesday", "Friday", "saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and Value is %v\n", index, day)
	}

	for _, day := range days { //when not needed index anymore
		fmt.Printf("index is Value is %v\n", day)
	}

	rougueVal := 1

	for rougueVal < 10 {
		if rougueVal == 2 {
			goto lco
		}
		if rougueVal == 5 {
			break
		}
		fmt.Println("value", rougueVal)
		rougueVal++
	}

lco:
	fmt.Println("Jumpjing at learning CodeONline.in")

}
