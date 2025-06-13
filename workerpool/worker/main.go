package main

import (
	"fmt"
	"sync"
	"time"
)

type EmailData struct {
	Header      string
	Subject     string
	Description string
}

func main() {
	numOfJobs := 10
	numOfWorkers := 2

	jobs := make(chan EmailData, numOfJobs)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// for i := range numOfJobs {
		// 	n := i + 1
		// 	emailData := EmailData{
		// 		Header:      fmt.Sprintf("Header_%d", n),
		// 		Subject:     fmt.Sprintf("Subject_%d", n),
		// 		Description: fmt.Sprintf("Desc_%d", n),
		// 	}
		// 	jobs <- emailData
		// }
		// close(jobs)

		// send every second now
		jobID := 1
		for {
			email := EmailData{
				Header:      fmt.Sprintf("Header_%d", jobID),
				Subject:     fmt.Sprintf("Subject_%d", jobID),
				Description: fmt.Sprintf("Desc_%d", jobID),
			}
			fmt.Printf("Queued job: %s\n", email.Header)
			jobs <- email
			jobID++
			time.Sleep(1 * time.Second)
		}
	}()

	for i := range numOfWorkers {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	// // Wait for the producer to finish, then close the channel
	// go func() {
	// 	wg.Wait()
	// 	close(jobs)
	// }()

	// // Wait for all workers to finish
	// wg.Wait()
}

func worker(workerID int, wg *sync.WaitGroup, jobs <-chan EmailData) {
	defer wg.Done()
	for emailData := range jobs {
		sendEmail(emailData)
	}

}

func sendEmail(emailData EmailData) {
	fmt.Println("sending email", emailData.Header)
	time.Sleep(time.Second)
	fmt.Println("email sending done....", emailData.Subject)
}
