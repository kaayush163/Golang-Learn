package main

import "fmt"

func main() {

	fmt.Println("Let's learn struct alternative of class")

	//Noi inheritence in golang or no super /parent conecpt is there
	aayush := User{"Aayush", "aayush@gmail.com", true, 23}

	fmt.Println(aayush)                            //{Aayush aayush@gmai.com true 23}
	fmt.Printf("aayush details are %+v\n", aayush) //+v gives you more deatils structure
	// {Name:Aayush Email:aayush@gmail.com Status:true Age:23}
	fmt.Printf("aayush anme are %vand email is%v\n", aayush.Name, aayush.Email) //+v gives you more deatils structure

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
