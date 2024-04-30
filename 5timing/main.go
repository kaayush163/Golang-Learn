package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("welcome to 5th tutorial")

	presentTime := time.Now()
	fmt.Println((presentTime))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // remeber this if you want to change date according any style of your format

	createdDate := time.Date(202, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

//video 9
