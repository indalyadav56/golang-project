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
	"time"
)

func process(ch chan string) {
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
		case <-time.After(time.Second * 10):
			return
		}
	}

}
