package test

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/xgzlucario/structx"
)

const ADD_NUM = 99999

func Benchmark_Map1(b *testing.B) {
	s := mapset.NewSet[int]()
	for i := 0; i < ADD_NUM; i++ {
		s.Add(i)
	}

	for i := 0; i < b.N; i++ {
		s.Contains(i % ADD_NUM)
	}
}

func Benchmark_Map2(b *testing.B) {
	maps := structx.NewMap[int, struct{}]()
	for i := 0; i < ADD_NUM; i++ {
		maps.Store(i, struct{}{})
	}
}

func Benchmark_Map3(b *testing.B) {
	maps := structx.NewSet[int]()
	for i := 0; i < ADD_NUM; i++ {
		maps.Add(i)
	}
}
