package test

import (
	"testing"

	"github.com/liyiheng/zset"
	"github.com/xgzlucario/structx"
)

func Benchmark_ZSet1(b *testing.B) {
	s := structx.New()
	for i := 0; i < b.N; i++ {
		s.IncrBy(1, int64(i%1000))
	}
}

func Benchmark_ZSet2(b *testing.B) {
	s := zset.New()
	for i := 0; i < b.N; i++ {
		s.IncrBy(1, int64(i%1000))
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
