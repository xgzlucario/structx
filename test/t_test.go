package test

import (
	"math"
	"testing"
)

func Benchmark_Test1(b *testing.B) {
	var a int64
	for i := 0; i < b.N; i++ {
		a = -1
	}
	if a > 0 {
	}
}

func Benchmark_Test2(b *testing.B) {
	var a int64
	for i := 0; i < b.N; i++ {
		a = math.MaxInt64
	}
	if a > 0 {
	}
}
