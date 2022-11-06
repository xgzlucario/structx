package test

import (
	"fmt"
	"testing"

	"github.com/liyiheng/zset"
	"github.com/xgzlucario/structx"
)

func Benchmark_ZSet1(b *testing.B) {
	s := structx.NewZSet[int, float64]()
	s.Set(0, 234)
	s.Set(0, 345)

	for i := 0; i < b.N; i++ {
		s.IncrBy(234, float64(i))
		s.IncrBy(345, float64(i*2))
	}
	fmt.Println("a", s.Len())
}

func Benchmark_ZSet2(b *testing.B) {
	s := zset.New()
	s.Set(0, 234, nil)
	s.Set(0, 345, nil)

	for i := 0; i < b.N; i++ {
		s.IncrBy(float64(i), 234)
		s.IncrBy(float64(i*2), 345)
	}
	fmt.Println("b", s.Length())
}

func Benchmark_ZSet3(b *testing.B) {
	s := structx.NewZSet[string, float64]()
	s.Set("xgz", 234)
	s.IncrBy("xgz", 1)

	s.Set("lxs", 1)
	s.IncrBy("lxs", 10)

	s.Print()
}
