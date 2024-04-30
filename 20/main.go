package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web ver videos - LCO")
	// PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status OCde", response.StatusCode)

	fmt.Println("content length is", response.ContentLength)

	var responseString strings.Builder //this is also methoid to create strings
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is :", byteCount)
	fmt.Println(responseString.String()) //holding the raw data all th time
	//best way to define strings is upper one becauseit allows more options as well
	// fmt.Println(content)
	// fmt.Println(string(content))

}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	  {
		"coursename":"Let's go with golang",
		"price" : 0,
		"platform" : "LearnCodeOnline"
	  }
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "aayush")
	data.Add("lastname", "man")
	data.Add("email", "aayush@gmail.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
