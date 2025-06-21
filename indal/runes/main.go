package main

import (
	"fmt"
	"strings"
)

func main() {
	toLower := strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}
		return r
	}, "INDAL")

	fmt.Println(toLower)

	fmt.Println(ToLower("KUMAR"))
}

func ToLower(str string) string {
	newRunes := []rune{}
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			newRunes = append(newRunes, char+32)
		}
	}
	return string(newRunes)
}
