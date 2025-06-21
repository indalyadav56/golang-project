package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 10)

	for i := range 10 {
		c <- i
		time.Sleep(time.Second)
	}
	close(c)

	for data := range c {
		fmt.Println(data)
	}
}
