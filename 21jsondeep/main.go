package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //want to create alias aname like we do in sql by AS
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`

	Tags []string `json:"tags,omitempty"` //space should not be there between , and omitemptyis imprtant a fter , and before omitempty

	//using omiempty will not show any firls if its' value is nil
}

func main() {
	fmt.Println("Json learn")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"web-dev", "js"}},
		{"Android Bootcamp", 299, "LearnCodeOnline.in", "def123", nil},
	}

	// finalJson, err := json.Marshal(lcoCourses) //marshal is actually the way how you implement or pass an interface
	//interface is a word being borrowed  alternative version of struct

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //"" blank is every it tab it starts with empty space first can heck by put some string instead of empty""

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	{
		"coursename":"ReactJS Bootcamp",
		"price" :209,
		"Website" : "LearnCodeOnline.in",
		"tags": [
			"web-dev",
			"js"
		]
	} 
	
	`)
	//If put ,after [] or after anything without writing anything after another key value in objects mistkenly then it will got o else condition so take care of this small things in golang

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was avlid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("Json was not valid")
	}

	//spme case where you want to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
