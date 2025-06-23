package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// myChannel := make(chan int, 10)

	// wg := &sync.WaitGroup{}

	// wg.Add(2)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	fmt.Println("Channel", <-myChannel)
	// 	wg.Done()
	// }(myChannel, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	myChannel <- 12
	// 	wg.Done()
	// 	close(myChannel)
	// }(myChannel, wg)

	// fmt.Println("myChannel:", <-myChannel)

	// wg.Wait()

	startTime := time.Now()

	jobs := make(chan int)
	workerNums := 3

	go func() {
		for i := range 10 {
			jobs <- i + 1
		}

		close(jobs)
	}()

	wg := sync.WaitGroup{}

	for i := range workerNums {
		wg.Add(1)
		go workerPool(i+1, jobs, &wg)
	}

	wg.Wait()

	fmt.Println("time it took to complete:=>", time.Now().Sub(startTime))

}

func workerPool(id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range jobs {
		fmt.Printf("WORKER_ID:= %d AND JOB_ID:==> %d \n", id, i)
		time.Sleep(time.Second)
	}

}
