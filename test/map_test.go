package test

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/xgzlucario/structx"
)

func Benchmark_Map1(b *testing.B) {
	s := mapset.NewSet[int]()
	for i := 0; i < b.N; i++ {
		s.Add(i)
		s.Contains(i)
	}
}

func Benchmark_Map2(b *testing.B) {
	maps := structx.NewMap[int, struct{}]()
	for i := 0; i < b.N; i++ {
		maps.Store(i, struct{}{})
		maps.Load(i)
	}
}

func Benchmark_Map3(b *testing.B) {
	maps := structx.NewSet[int]()
	for i := 0; i < b.N; i++ {
		maps.Add(i)
		maps.Exist(i)
	}
}
