package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	jobs := make(chan int)
	result := make(chan int)

	numObJob := 10
	numOfWorker := 5

	wg := sync.WaitGroup{}

	for i := range numObJob {
		jobs <- i + 1
	}
	close(jobs)

	for range numOfWorker {
		wg.Add(1)
		go workerPool(jobs, &wg, result)
	}

	wg.Wait()

	for d := range result {
		fmt.Println(d)
	}

	// ch2 := make(chan string)

	// wg := sync.WaitGroup{}

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for range 5 {
	// 		ch1 <- "From ch1"
	// 		time.Sleep(time.Second)
	// 	}
	// 	close(ch2)
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for range 5 {
	// 		ch2 <- "From ch2"
	// 		time.Sleep(time.Second)
	// 	}
	// 	close(ch2)
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for {
	// 		select {
	// 		case msg1 := <-ch1:
	// 			fmt.Println(msg1)
	// 		case msg2 := <-ch2:
	// 			fmt.Println(msg2)
	// 		}
	// 	}
	// }()

	// wg.Wait()

}

func workerPool(jobs <-chan int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	for d := range jobs {
		fmt.Println("channel data:= ", d)
	}

	result <- 1
}
