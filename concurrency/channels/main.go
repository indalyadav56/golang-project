package main

import "fmt"

func main() {
	done := make(chan int, 10)
	// done <- 2
	fmt.Println(<-done)
}
