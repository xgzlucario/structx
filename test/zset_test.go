package test

import (
	"testing"

	"github.com/liyiheng/zset"
	"github.com/xgzlucario/structx"
)

func Benchmark_ZSet1(b *testing.B) {
	s := structx.NewZSet[int, float64]()
	s.Set(0, 234, nil)
	s.Set(0, 456, nil)

	for i := 0; i < b.N; i++ {
		s.IncrBy(234, float64(i))
		s.IncrBy(456, float64(i*2))
	}
}

func Benchmark_ZSet2(b *testing.B) {
	s := zset.New()
	s.Set(0, 234, nil)
	s.Set(0, 345, nil)

	for i := 0; i < b.N; i++ {
		s.IncrBy(float64(i), 234)
		s.IncrBy(float64(i*2), 456)
	}
}

// func Benchmark_ZSet3(b *testing.B) {
// 	ss := structx.NewZSet[string, int]()
// 	ss.Incr("a1", 3)
// 	ss.Print()
// 	ss.Incr("a2", 8)
// 	ss.Print()
// 	ss.Incr("a3", 5)
// 	ss.Print()
// 	ss.Incr("a1", 10)
// 	ss.Print()
// }
