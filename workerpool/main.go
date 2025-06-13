package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 10)
	numOfWorkers := 5
	jobResult := make(chan map[string]any, 10)

	// Send jobs
	for i := range 10 {
		jobs <- i
	}

	wg := sync.WaitGroup{}

	for i := range numOfWorkers {
		wg.Add(1)
		go worker(i, jobs, &wg, jobResult)
	}

	close(jobs)
	wg.Wait()

	close(jobResult)

	for key := range jobResult {
		fmt.Println(key)
	}

}

func worker(workerID int, jobs <-chan int, wg *sync.WaitGroup, result chan<- map[string]any) {
	defer wg.Done()
	for i := range jobs {
		mapData := map[string]any{
			fmt.Sprintf("job_%d", i): "this job result store in string value",
		}
		result <- mapData
	}
}
