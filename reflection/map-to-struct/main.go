package main

import "reflect"

func MapToStrcut(m map[string]interface{}) {
	// // Define fields for the new struct
	// fields := []reflect.StructField{
	// 	{
	// 		Name: "Age",
	// 		Type: reflect.TypeOf(0), // int
	// 		Tag:  `json:"age"`,
	// 	},
	// }

	// // Create a new struct type dynamically
	// dynamicType := reflect.StructOf(fields)

	// t := reflect.TypeOf(m)
	// v := reflect.ValueOf(m)

	// for i := range t.NumField() {

	// }

}

func MapToStruct(data map[string]interface{}, result interface{}) {
	val := reflect.ValueOf(result).Elem()

	for key, value := range data {
		field := val.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			valValue := reflect.ValueOf(value)
			if valValue.Type().AssignableTo(field.Type()) {
				field.Set(valValue)
			}
		}
	}
}
