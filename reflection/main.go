package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// slElement := reflect.SliceOf(reflect.TypeOf(0))

	// fmt.Println(slElement)

	// slc := reflect.MakeSlice(slElement, 0, 0)

	// fmt.Println(slc)

	// sl := slc.Interface().([]int)

	// newSlice := append(sl, 100)

	// fmt.Println(newSlice)

	// obj := reflect.typeof()

	type MyStruct struct {
		A int64
		B int32
	}

	offset := unsafe.Offsetof(MyStruct{}.B)
	fmt.Println(offset) // Byte offset of field B within the struct

}
