package main

import "testing"

func BenchmarkLeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeapYear(2020)
	}
}
