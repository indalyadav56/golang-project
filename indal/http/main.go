package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("GET /users", MyCustomHandler)

	http.ListenAndServe(":8080", server)

}

func MyCustomHandler(w http.ResponseWriter, r *http.Request) {

}

func Testnewhandle() {

	w := httptest.NewRecorder()

	r := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer([]byte("indal")))

	MyCustomHandler(w, r)

	bodyData, _ := io.ReadAll(w.Body)

	fmt.Println(string(bodyData))

}
