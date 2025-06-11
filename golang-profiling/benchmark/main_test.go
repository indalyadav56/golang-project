package main

import "testing"

func BenchmarkMyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunc()
	}
}

func MyFunc() {
	for i := 0; i < 10_000_00; i++ {
		_ = i * i
	}
}
