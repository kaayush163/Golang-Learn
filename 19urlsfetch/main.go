package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are %T\n", qparams)

	fmt.Println(qparams["list"])

	for _, value := range qparams {
		fmt.Println("Param is ", value) // all query params value will be printed
	}

	partsOfUrl := &url.URL{ //this is to create a url
		Scheme:  "https",
		Host:    "youtube.com",
		Path:    "/watch",
		RawPath: "v=yNjPL_LJrWs",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
