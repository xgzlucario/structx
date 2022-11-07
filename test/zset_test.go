package test

import (
	"testing"

	"github.com/liyiheng/zset"
	"github.com/xgzlucario/structx"
)

const NUM = 49999

func Benchmark_ZSet1(b *testing.B) {
	s := structx.NewSkipList[struct{}, int]()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func Benchmark_ZSet2(b *testing.B) {
	s := structx.NewZSet[int, int]()
	for i := 0; i < b.N; i++ {
		s.Incr(i, i)
	}
}

func Benchmark_ZSet3(b *testing.B) {
	s := zset.New()
	for i := 0; i < b.N; i++ {
		s.Set(float64(i), int64(i), nil)
	}
}
