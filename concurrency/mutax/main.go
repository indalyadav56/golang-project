package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	counter := 0
	mutax := sync.Mutex{}

	go func() {
		mutax.Lock()
		defer mutax.Unlock()

		for range 100 {
			counter += 1
		}
	}()

	go func() {
		counter += 1
	}()

	fmt.Println(counter)
	<-time.After(time.Second * 5)
	fmt.Println(counter)

}
