package main

import "testing"

func BenchmarkMyFunc(b *testing.B) {
	for range b.N {
		MyFunc()
	}
}

func MyFunc() {
	for i := range 100000 {
		_ = i * i
	}
}
