package test

import (
	"testing"

	"github.com/liyiheng/zset"
	"github.com/xgzlucario/structx"
)

const NUM = 1000

func Benchmark_ZSet1(b *testing.B) {
	s := structx.NewSkipList[int64, float64]()
	for i := 0; i < NUM; i++ {
		s.Add(float64(i))
	}

	for i := 0; i < b.N; i++ {
		s.Range(0, -1, func(key int64, value float64) {})
	}
}

// func Benchmark_ZSet2(b *testing.B) {
// 	s := structx.NewZSet[int, int]()
// 	for i := 0; i < b.N; i++ {
// 		s.Incr(i, i)
// 	}
// }

func Benchmark_ZSet3(b *testing.B) {
	s := zset.New()
	for i := 0; i < NUM; i++ {
		s.Set(float64(i), int64(i), nil)
	}

	for i := 0; i < b.N; i++ {
		s.Range(0, -1, func(f float64, i1 int64, i2 interface{}) {})
	}
}

// func Benchmark_ZSet4(b *testing.B) {
// 	s := structx.NewZSet[string, int]()
// 	s.Incr("asd", 123)
// 	s.Incr("rr", 12)
// 	s.Incr("gg", 173)
// 	s.Incr("ww", 14)
// 	s.Incr("ww", 17)
// 	s.Print()

// 	fmt.Println(s.GetByRank(2))
// }
