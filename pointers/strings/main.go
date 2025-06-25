package main

import (
	"fmt"
	"strings"
)

func main() {

	// intToStr := 100

	// newStrValue := strconv.Itoa(intToStr)
	// // fmt.Println("newStrValue", newStrValue)

	// newIntValue, _ := strconv.Atoi(newStrValue)
	// fmt.Println(newIntValue)

	// newBool := true
	// strData := strconv.FormatBool(newBool)

	// fmt.Println(strData)

	// str to float

	// //float to string
	// floatVlaue, _ := strconv.ParseFloat("1.00", 64)
	// fmt.Println(floatVlaue)

	// str := "   in d al    "
	// fmt.Println(str)

	// newValue := strings.ToUpper(str)
	// strLower := strings.ToLower(str)

	// strData := strings.TrimSpace(str)
	// fmt.Println(strData)

	// fmt.Println(newValue, strLower)

	// str := "indal"

	// exists := strings.HasPrefix(str, "i")
	// exists := strings.HasSuffix(str, "a")
	// fmt.Println(exists)

	// substrings

	// str := "indal"

	// // i, in, ind, inda, indal, nd, nda, ndal, dal, .l

	// isExists := strings.Contains(str, "lad")
	// fmt.Println(isExists)

	// str := []string{"test", "indal", "kumar", "another", "data"}

	// //  "test-indal-kumar-another-data"
	// fmt.Println(strings.Join(str, "?"))

	// str := "test-indal-kumar-another-data"

	// fmt.Println(strings.Split(str, ""))

	// strings replace
	str := "indalaa"

	// fmt.Print(strings.Replace(str, "a", "new", 3))

	fmt.Println(strings.ReplaceAll(str, "a", "newvalue"))

}
