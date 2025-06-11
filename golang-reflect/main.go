package main

import "reflect"

func main() {
	//dynamic data types creations using reflect package

	intData := reflect.TypeOf(0)
	intValue := reflect.New(intData).Elem()
	intValue.SetInt(42) // Set the value to 42
	println("int value:", intValue.Int())

	stringData := reflect.TypeOf("")
	stringType := reflect.New(stringData).Elem()

	stringType.SetString("Hello, World!") // Set the value to "Hello, World!"
	println("string value:", stringType.String(), stringData)

	floatData := reflect.TypeOf(0.0)
	boolData := reflect.TypeOf(true)
	sliceData := reflect.TypeOf([]int{})
	mapData := reflect.TypeOf(map[string]int{})
	structData := reflect.TypeOf(struct {
		Name string
		Age  int
	}{})
	// Print the types
	println("int type:", intData.String())
	println("string type:", stringData.String())
	println("float type:", floatData.String())
	println("bool type:", boolData.String())
	println("slice type:", sliceData.String())
	println("map type:", mapData.String())
	println("struct type:", structData.String())
	// You can also use reflect.ValueOf to get the value of a variable

}
