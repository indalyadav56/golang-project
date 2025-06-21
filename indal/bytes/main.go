package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Key string `json:"key"`
}

func main() {
	// var buf bytes.Buffer

	// buf.Read([]byte("indal"))

	// buf.Write([]byte(`{"key":"value","key2":"value2"}`))
	// var mapData User

	// // json.NewDecoder(&buf).Decode(&mapData) // golang datastructure and encoder mean json
	// json.NewEncoder(&buf).Encode(&mapData) // golang datastructure and encoder mean json

	// fmt.Printf("Byte size of buffer: %d bytes\n", buf.Len())

	// fmt.Println(len([]byte("indalkumaryadavhowareyou😀indalkumaryadavhowareyou😀indalkumaryadavhowareyou😀indalkumaryadavhowareyou😀indalkumaryadavhowareyou😀indalkumaryadavhowareyou😀")))

	// http.ResponseWriter.Write()
	// fmt.Println(reflect.TypeOf([]byte('a')).Size())

	// srv := http.NewServeMux()

	// srv.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("successsuccesssuccesssuccesssuccesssuccesssuccesssuccesssuccesssuccesssuccesssuccesssuccesssuccess"))
	// })
	// http.ListenAndServe(":8080", srv)

	// client

	// what will happen

	// content-length :- 100 bytes

	// make([]byte, 0, len(res.content-length))

	// fourByteData := make([]byte, 0, 4)
	// fourByteData = append(fourByteData, 255, 255, 97, 97)

	// fmt.Println(fourByteData)         // [97 97 97 97]
	// fmt.Println(string(fourByteData)) // aaaa
	// fmt.Println(len(fourByteData))    // 4
	// fmt.Println(cap(fourByteData))    // 4

	// confusion

	// fmt.Println("===================")
	// intData := 10000

	// fmt.Println([]byte(strconv.Itoa(intData)))
	// fmt.Println(len(strconv.Itoa(intData)))

	// fmt.Println(reflect.TypeOf(intData).Size())

	// memory layout

	// // intData := int64(10)   // use fixed size
	// buf := make([]byte, 10, 100) // 8 bytes for int64

	// buf = append(buf, byte('a'), byte('a'))

	// fmt.Println("Binary bytes:", buf)
	// fmt.Println("Length:", len(buf)) // always 8

	// server := http.NewServeMux()

	// server.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	// w.Header().Set("Content-Type", "application/octet-stream")
	// 	// json.NewEncoder(w).Encode(buf)

	// 	w.Write(make([]byte, 10))
	// })

	// http.ListenAndServe(":8080", server)

	fmt.Println(reflect.TypeOf(1).Size())

}
