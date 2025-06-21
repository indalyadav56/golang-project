package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]int{"one": 1, "two": 2}

	v := reflect.ValueOf(m)

	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)
		fmt.Printf("Key: %v, Value: %v\n", key.Interface(), val.Interface())
	}

}
