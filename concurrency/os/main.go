package main

import (
	"fmt"
	"os"
)

func main() {
	defer recover()

	myFunc()

	for range 10 {
		fmt.Println("INDAL")
	}

}

func myFunc() {
	// panic("My Custom Panic")
	anotherFunc()
}

func anotherFunc() {
	os.Exit(1)
}
