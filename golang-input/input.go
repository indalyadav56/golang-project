package main

import (
	"fmt"
	"math"
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

	// fmt.Println(countDigitsMath(100000000000))

	// fmt.Println(SquareRoot(16))
	// fmt.Println(Power(2, 3))

	// fmt.Println(IsPrime(2))
	// fmt.Println(IsPrime(3))
	// fmt.Println(IsPrime(5))
	// fmt.Println(IsPrime(7))
	// fmt.Println(IsPrime(8))
	// fmt.Println(IsPrime(10))

}

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	sqrt := math.Sqrt(float64(num))

	for i := 3; i < int(sqrt); i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func SquareRoot(num int) int {
	return int(math.Sqrt(float64(num)))
}

func Power(x, y float64) int {
	return int(math.Pow(x, y))
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

// countDigitsMath counts digits using division, handling negative numbers
func countDigitsMath(num int) int {
	if num == 0 {
		return 1
	}

	absNum := int(math.Abs(float64(num)))
	count := 0
	for absNum > 0 {
		absNum /= 10
		count++
	}

	return count
}
