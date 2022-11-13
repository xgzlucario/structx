package test

import (
	"testing"

	"github.com/liyiheng/zset"
	"github.com/xgzlucario/structx"
)

const NUM = 1000

func getZSet1() *structx.Skiplist[int64, float64] {
	s := structx.NewSkipList[int64, float64]()
	for i := 0; i < NUM; i++ {
		s.Add(0, float64(i))
	}
	return s
}

func getZSet2() *zset.SortedSet {
	s := zset.New()
	for i := 0; i < NUM; i++ {
		s.Set(float64(i), int64(i), nil)
	}
	return s
}

// ========= Add =========
func Benchmark_ZSetAdd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getZSet1()
	}
}

func Benchmark_ZSetAdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getZSet2()
	}
}

// ========= Range =========
func Benchmark_ZSetRange1(b *testing.B) {
	s := getZSet1()
	for i := 0; i < b.N; i++ {
		s.Range(0, -1, func(key int64, value float64) bool {
			return false
		})
	}
}

func Benchmark_ZSetRange2(b *testing.B) {
	s := getZSet2()
	for i := 0; i < b.N; i++ {
		s.Range(0, -1, func(f float64, i1 int64, i2 interface{}) {})
	}
}

// ========= Rank =========
func Benchmark_ZSetRank1(b *testing.B) {
	s := getZSet1()
	for i := 0; i < b.N; i++ {
		s.GetByRank(76)
	}
}

func Benchmark_ZSetRank2(b *testing.B) {
	s := getZSet2()
	for i := 0; i < b.N; i++ {
		s.GetDataByRank(76, true)
	}
}
