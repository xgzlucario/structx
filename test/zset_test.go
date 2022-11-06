package test

import (
	"testing"

	"github.com/liyiheng/zset"
	"github.com/xgzlucario/structx"
)

func Benchmark_ZSet1(b *testing.B) {
	s := structx.NewZSet[int, float64]()
	s.Set(1, 0)
	s.Set(2, 0)
	s.Set(3, 0)

	for i := 0; i < b.N; i++ {
		s.IncrBy(1, float64(i))
		s.IncrBy(2, float64(i*2))
		s.IncrBy(3, float64(i*3))
	}
}

func Benchmark_ZSet2(b *testing.B) {
	s := zset.New()
	s.Set(0, 1, nil)
	s.Set(0, 2, nil)
	s.Set(0, 3, nil)

	for i := 0; i < b.N; i++ {
		s.IncrBy(float64(i), 1)
		s.IncrBy(float64(i*2), 2)
		s.IncrBy(float64(i*3), 3)
	}
}

func Benchmark_ZSet3(b *testing.B) {
	s := structx.NewZSet[string, float64]()
	s.Set("x1", 234)
	s.IncrBy("x1", 1)

	s.Set("x2", 1)
	s.IncrBy("x2", 10)

	s.Print()
}
