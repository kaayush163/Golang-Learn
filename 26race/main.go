package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Race Condition ")
	wg := &sync.WaitGroup{} //weight group is ususally a pointer
	mut := &sync.RWMutex{}  //whenever you access a memory you should lock it

	// mut.RLock() //not use here
	var score = []int{0}
	// mut.Unlock()

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut) //goroutines fucnton having multiple threads

	// wg.Add(1)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

		wg.Done()
	}(wg, mut) //goroutines fucnton having multiple threads
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three routine")
		mut.RLock()
		score = append(score, 3)
		mut.RUnlock()

		wg.Done()
	}(wg, mut) //goroutines fucnton having multiple threads

	wg.Wait()

	fmt.Println((score))
}
