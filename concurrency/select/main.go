// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func server1(ch chan string) {
// 	for i := range 10 {
// 		ch <- "from server1" + strconv.Itoa(i)
// 	}
// }

// func server2(ch chan string) {
// 	for i := range 10 {
// 		ch <- "from server2" + strconv.Itoa(i)
// 	}
// }

// func main() {
// 	output1 := make(chan string)
// 	output2 := make(chan string)

// 	go server1(output1)

// 	go server2(output2)

// 	for {
// 		select {
// 		case s1 := <-output1:
// 			fmt.Println(s1)
// 		case s2 := <-output2:
// 			fmt.Println(s2)
// 		default:
// 			fmt.Println("exiting... ")
// 			return
// 		}
// 	}
// }

package main

import (
	"fmt"
	"runtime"
)

func process(ch chan string) {
	for range 10 {
		ch <- "process successful"
	}
}

func process2(ch chan string) {
	for range 5 {
		ch <- "process successful"
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	ch1 := make(chan string)
	ch2 := make(chan string)

	go process(ch1)
	go process2(ch1)

	for range 15 {
		select {
		case v := <-ch1:
			fmt.Println("received value: ", v)
		case v := <-ch2:
			fmt.Println("received value: ", v)
		}
	}

}
