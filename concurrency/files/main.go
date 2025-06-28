package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
		close(ch)
	}()

	for d := range ch {
		fmt.Println(d)
	}

}
