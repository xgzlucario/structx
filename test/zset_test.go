package test

import (
	"testing"

	"github.com/liyiheng/zset"
	"github.com/xgzlucario/structx"
)

func Benchmark_ZSet1(b *testing.B) {
	s := structx.NewZSet[int, int]()
	for i := 0; i < b.N; i++ {
		s.IncrBy(i, i)
	}
}

func Benchmark_ZSet2(b *testing.B) {
	s := zset.New()
	for i := 0; i < b.N; i++ {
		s.Set(float64(i), int64(i), nil)
	}
}

func Benchmark_ZSet3(b *testing.B) {
	s := structx.NewSortList[struct{}, int]()
	for i := 0; i < b.N; i++ {
		s.Insert(i)
	}
}
