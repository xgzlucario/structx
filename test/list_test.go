package test

import (
	"testing"
)

func Benchmark_List1(b *testing.B) {
	ls := make(map[int]string)
	for i := 0; i < b.N; i++ {
		ls[i] = "abc"
	}
}

func Benchmark_List2(b *testing.B) {
	ls := make(map[int]string, 8)
	for i := 0; i < b.N; i++ {
		ls[i] = "abc"
	}
}

func Benchmark_List3(b *testing.B) {
	ls := make(map[int]string, 32)
	for i := 0; i < b.N; i++ {
		ls[i] = "abc"
	}
}

func Benchmark_List4(b *testing.B) {
	ls := make(map[int]string, 128)
	for i := 0; i < b.N; i++ {
		ls[i] = "abc"
	}
}
