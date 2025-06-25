package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string
	Email string
}

func main() {
	// sliceType := reflect.SliceOf(reflect.TypeOf(0))
	// slice := reflect.MakeSlice(sliceType, 0, 0)

	// slice = reflect.Append(slice, reflect.ValueOf(100), reflect.ValueOf(200))

	// sl := slice.Interface().([]int)

	// sl = append(sl, 300, 400)

	// fmt.Println(slice.Interface())
	// fmt.Println(sl)

	mapData := make(map[string]interface{}, 0)

	user := User{
		Name:  "indal",
		Email: "indal",
	}

	utype := reflect.TypeOf(user)
	// value := reflect.ValueOf(user)

	for i := range utype.NumField() {
		fmt.Println(utype.Field(i).Name)
		mapData[utype.Field(i).Name] = utype.Field(i).
	}

	fmt.Println(mapData)

}
