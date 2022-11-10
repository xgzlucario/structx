package test

import (
	"testing"

	"github.com/liyiheng/zset"
	"github.com/xgzlucario/structx"
)

const NUM = 1000

func Benchmark_ZSet1(b *testing.B) {
	s := structx.NewSkipList[struct{}, int]()
	for i := 0; i < NUM; i++ {
		s.Add(i)
	}

	for i := 0; i < b.N; i++ {
		s.Range(0, -1, func(key struct{}, value int) {})
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
