package main

import (
	"fmt"
)

var numChan chan int32
var resChan chan int32
var sum int32

func main() {
	numChan = make(chan int32, 2000)
	resChan = make(chan int32, 1)
	for i := 1; i <= 2000; i++ {
		numChan <- int32(i)
	}
	close(numChan)
	for i := 0; i < 2; i++ {
		go getSum()
	}
	fmt.Println(<-resChan)

}
func getSum() {
	for v := range numChan {
		fmt.Println(v)
		sum += v
		//atomic.AddInt32(&sum, v)
	}

	resChan <- sum
}
