package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	// Describe(10)
	// Describe("hello")
	// Describe([]int{1, 2, 3})

	// Describe(map[string]int{"a": 1, "b": 2})

	// Describe(struct {
	// 	A, B int
	// }{
	// 	A: 1,
	// 	B: 2,
	// })

	// StructToMap(struct {
	// 	A, B       int `json:"a_data", json:"b_data" yaml:"a_data" yaml:"b_data"`
	// 	Data       string
	// 	Test       rune
	// 	CustomData map[string]int
	// }{
	// 	A: 1,
	// 	B: 2,
	// })

	// StructToMap(10)

}

func Describe(v interface{}) {
	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Ptr:
		fmt.Println("Pointer to:", rv.Elem().Kind())
	case reflect.Slice:
		fmt.Println("Slice with length:", rv.Len())
	case reflect.Map:
		fmt.Println("Map with keys:", rv.MapKeys())
	case reflect.Struct:
		fmt.Println("THis is struct data types:", rv.Kind())
	default:
		fmt.Println("Kind:", rv.Kind())
	}
}

func StructToMap(structObj interface{}, shouldJsonTag ...bool) (error, map[string]interface{}) {
	result := make(map[string]interface{})
	t := reflect.TypeOf(structObj)
	v := reflect.ValueOf(structObj)

	if t.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("given object is not a struct, provided object is of type: %s", t.Kind())), nil
	}

	fmt.Println("values of struct:-->", v)

	for i := range v.NumField() {
		fmt.Println("value of struct:-->", t.Field(i).Name, v.Field(i).Type(), t.Field(i).Tag.Get("json"))
		key := t.Field(i).Tag.Get("json")
		value := v.Field(i).Interface()
		fmt.Println("key:-->", key, "value:-->", value)
		result[key] = value
	}

	// val := reflect.ValueOf(structObj)
	// typ := reflect.TypeOf(structObj)

	// fmt.Println("Type:", typ)
	// fmt.Println("Kind:", typ.Kind())
	// fmt.Println("Value:", val)

	return nil, result
}

func MapToStruct(structObj interface{}, shouldJsonTag ...bool) (error, map[string]interface{}) {
	result := make(map[string]interface{})
	t := reflect.TypeOf(structObj)
	v := reflect.ValueOf(structObj)

	if t.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("given object is not a struct, provided object is of type: %s", t.Kind())), nil
	}

	fmt.Println("values of struct:-->", v)

	for i := range v.NumField() {
		fmt.Println("value of struct:-->", t.Field(i).Name, v.Field(i).Type(), t.Field(i).Tag.Get("json"))
		key := t.Field(i).Tag.Get("json")
		value := v.Field(i).Interface()
		fmt.Println("key:-->", key, "value:-->", value)
		result[key] = value
	}

	// val := reflect.ValueOf(structObj)
	// typ := reflect.TypeOf(structObj)

	// fmt.Println("Type:", typ)
	// fmt.Println("Kind:", typ.Kind())
	// fmt.Println("Value:", val)

	return nil, result
}
