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

	aayush.GetStatus()
	aayush.NewMail()

	fmt.Printf("Name is %v and email is %v \n", aayush.Name, aayush.Email) //copy of obecjt is jsut passing here we ned to use pointer

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active", u.Status)
}

func (u User) NewMail() {
	u.Email = "shanu@gmail.com"
	fmt.Println("Is User active", u.Email)
}
