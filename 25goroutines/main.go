package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // isntead of using time.Sleep () this is the best method of using sync.WaitGroup
// v ery steroid advanced version of time.Sleep()

var mut sync.Mutex //pointer

func main() {

	//example aprallelism by using go keyword
	// go greeter("heelo")  //go keywork launch multilple threads or goroutines and it's going to just do the job
	// greeter("world")

	websitelist := []string{
		"https://github.com",
		"https://go.dev",
		"https://youtube.com",
	}

	for _, web := range websitelist {
		// getStatusCode(web)
		go getStatusCode(web)
		wg.Add(1) // 1. add

	}

	wg.Wait() //2. not allow main method to finish
}

// func greeter(s string) {

// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second) //3 seconds wait for three second s and print hello
// 		//by millisecond it can be done much more fatser
// 		fmt.Println((s))
// 	}
// }

// world
// hello
// hello
// world
// world
// hello
// world

func getStatusCode(endpoint string) {

	defer wg.Done() //3. done of waig group
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")

	} else {
		mut.Lock()
		signals = append(signals, endpoint)

		mut.Unlock()
		fmt.Printf("%d status code for %s", res.StatusCode, endpoint)
	}
}
