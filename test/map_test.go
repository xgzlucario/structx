package test

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

func Benchmark_Map1(b *testing.B) {
	ls := make(map[int]struct{})
	for i := 0; i < b.N; i++ {
		ls[i] = struct{}{}
	}
}

func Benchmark_Map2(b *testing.B) {
	s := mapset.NewSet[int]()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}
