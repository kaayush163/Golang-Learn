package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int, 1) //kind of buffer channel want to add 1 more value to it like 6

	wg := &sync.WaitGroup{} //whenver creating goroutines we always create weight group

	// myCh <- 5 //put value in channel

	// fmt.Println(<-myCh)
	//channel work dont like this they work with go routines

	wg.Add(2) //adding two go routines

	//RECEIVE ONLY
	//<-befor chan means value going outside of the box
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-myCh //this thing we are doing to if get val 0 how do we know if it actuall a val or it a message 0 that the connection close

		fmt.Println(isChanelOpen) //give true on paasing value or false if not passing any value and closed
		fmt.Println(val)

		//fmt.Println(<-myCh)
		// fmt.Println(<-myCh) //for print ing 6 make println another copy but this not good way we use buffer channel

		wg.Done()
	}(myCh, wg) //first channel will responsibel for receiving the value

	//SEND ONLY
	//<- after chan means something value going inside of channel box
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// close(myCh) //if close here then it goes in panic statement still trying to accessing the values
		myCh <- 5 //print 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg) //putting value to the channel

	wg.Wait()
}
