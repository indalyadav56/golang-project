package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"name" indal:"name_field"`
	Email string `json:"email" indal:"email_field"`
}

func main() {

	// mapData := make(map[string]interface{}, 0)

	// user := User{
	// 	Name:  "indal",
	// 	Email: "indal@gmail.com",
	// }

	// utype := reflect.TypeOf(user)
	// value := reflect.ValueOf(user)

	// for i := range utype.NumField() {
	// 	mapData[utype.Field(i).Tag.Get("indal")] = value.Field(i)
	// }

	// fmt.Println(mapData)

	MakeDynamicStruct()

}

func MakeDynamicStruct() {
	structField := []reflect.StructField{
		{
			Name: "Indal",
			Type: reflect.TypeOf(""),
		},
	}

	dynamicStruct := reflect.StructOf(structField)

	ins := reflect.New(dynamicStruct).Elem()

	ins.FieldByName("Indal").SetString("indal kumar strings goes herer")

	fmt.Printf("%v", dynamicStruct)

}

type dynamicStruct struct {
	indal string
}
