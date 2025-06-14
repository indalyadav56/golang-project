package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter numbers separated by space: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("input:-->", input)

	// numbers := strings.Fields(input)
	// fmt.Println("numbers:-->", numbers)
	// sum := 0

	// for _, num := range numbers {
	// 	n, err := strconv.Atoi(num)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		continue
	// 	}
	// 	sum += n
	// }

	// fmt.Println("Sum:", sum)

	// LeapYear(2004)

	// var x interface{} = -1

	// fmt.Println(reflect.TypeOf(x)) // Output: int

	// num := uint(-10)
	// fmt.Println(num)

}

// if year%4 == 0 {
//     if year%100 == 0 {
//         if year%400 == 0 {
//             // Leap year
//         } else {
//             // Not a leap year
//         }
//     } else {
//         // Leap year
//     }
// } else {
//     // Not a leap year
// }

// A leap year is:
// - divisible by 4
// - not divisible by 100, unless also divisible by 400
func LeapYear(year int) {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("leap year")
	} else {
		fmt.Println("not leap year")
	}
}
