package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"testing"
)

func main2() {
	res, err := http.Get("https://jsonplacehoolder")
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Status, res.StatusCode)

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	customStruct := struct {
		UserId int    `json:"userId"`
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}{}

	if err := json.Unmarshal(byteData, &customStruct); err != nil {
		panic(err)
	}

	defer res.Body.Close()
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(200)
	}
}

func ContextTest(ctx context.Context) (context.Context, func()) {
	newCtx, cancelFuncCallback := context.WithCancel(ctx)
	return newCtx, cancelFuncCallback
}
