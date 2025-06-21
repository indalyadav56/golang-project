package main

import (
	"fmt"
	"sync"
)

func main() {
	myChannel := make(chan int, 10)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Channel", <-myChannel)
		wg.Done()
	}(myChannel, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myChannel <- 12
		wg.Done()
		close(myChannel)
	}(myChannel, wg)

	fmt.Println("myChannel:", <-myChannel)

	wg.Wait()
}
