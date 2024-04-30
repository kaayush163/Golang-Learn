package main

import "fmt"

func main() {

	fmt.Println("lets learn map")

	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("JS Shorts for:", language["JS"])

	delete(language, "RB")
	fmt.Println("ag=fter deleting:", language)

	//lops work with map
	for _, value := range language { //not need key so _, used with :=
		fmt.Printf("For key value is %v\n", value)
	}
}
