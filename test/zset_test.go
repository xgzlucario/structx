package test

import (
	"testing"

	"github.com/liyiheng/zset"
	"github.com/sourcegraph/conc/pool"
	"github.com/xgzlucario/structx"
)

const NUM = 1000

func getZSet1() *structx.ZSet[int64, float64] {
	s := structx.NewZSet[int64, float64]()
	for i := 0; i < NUM; i++ {
		s.Incr(0, float64(i))
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

// ========= Delete =========
func Benchmark_ZSetDelete1(b *testing.B) {
	s := getZSet1()
	for i := 0; i < b.N; i++ {
		s.Delete(int64(i))
	}
}

func Benchmark_ZSetDelete2(b *testing.B) {
	s := getZSet2()
	for i := 0; i < b.N; i++ {
		s.Delete(int64(i))
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
		s.GetByRank(3899)
	}
}

func Benchmark_ZSetRank2(b *testing.B) {
	s := getZSet2()
	for i := 0; i < b.N; i++ {
		s.GetDataByRank(3899, true)
	}
}

func BenchmarkPool1(b *testing.B) {
	p := structx.NewPool().WithMaxGoroutines(100)
	for i := 0; i < b.N; i++ {
		p.Go(func() {})
	}
	p.Wait()
}

func BenchmarkPool2(b *testing.B) {
	p := pool.New().WithMaxGoroutines(100)
	for i := 0; i < b.N; i++ {
		p.Go(func() {})
	}
	p.Wait()
}
