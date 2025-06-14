package main

import "fmt"

func main() {
	// fmt.Println(ReverseString("indal"))
	// fmt.Println(ReverseArray([6]int{2, 4, 6, 8, 10, 12}))
	// fmt.Println(isPelindrome(""))

	// for i := range 10 {
	// 	fmt.Println(fibonnachi(i))
	// }

	fmt.Println(EvenOdd(123))
}

func ReverseString(str string) string {
	r := []rune(str)
	left, right := 0, len(r)-1

	for left < right {
		r[left], r[right] = r[right], r[left]

		left++
		right--
	}

	return string(r)
}

func ReverseArray(arr [6]int) [6]int {
	left, right := 0, len(arr)-1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]

		left++
		right--
	}

	return arr
}

func isPelindrome(str string) bool {
	if len(str) == 0 {
		return false
	}

	left, right := 0, len(str)-1

	for left < right {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func fibonnachi(num int) int {
	if num <= 1 {
		return num
	}
	return fibonnachi(num-1) + fibonnachi(num-2)
}

func EvenOdd(num int) string {
	if num%2 == 0 {
		return "this is even"
	}
	return "this is odd"
}
