package main

import (
	"fmt"
	"sync"
)


func main() {
	myChannel := make(chan int)

	wg := &sync.WaitGroup{}
	
	wg.Add(10)

	go func (ch chan int, wg *sync.WaitGroup)  {
		fmt.Println("Channel",<-myChannel)
		wg.Done()
	}(myChannel, wg)


	go func (ch chan int, wg *sync.WaitGroup)  {
		myChannel<- 12
		wg.Done()
	}(myChannel, wg)
	

	fmt.Println("myChannel:", myChannel)
	fmt.Println("myChannel:", <-myChannel)
}


