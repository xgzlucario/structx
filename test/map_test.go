package test

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/xgzlucario/structx"
)

const SUM = 1000

func getMapSet() mapset.Set[int] {
	s := mapset.NewThreadUnsafeSet[int]()
	for i := 0; i < SUM; i++ {
		s.Add(i)
	}
	return s
}

func getListSet() *structx.LSet[int] {
	s := structx.NewLSet[int]()
	for i := 0; i < SUM; i++ {
		s.Add(i)
	}
	return s
}

// func getZSet() *structx.ZSet {
// 	s := structx.New()
// 	for i := 0; i < SUM; i++ {
// 		s.IncrBy(float64(i), int64(i))
// 	}
// 	return s
// }

// ============ Range ============
func Benchmark_MapSetRange(b *testing.B) {
	s := getMapSet()
	for i := 0; i < b.N; i++ {
		s.Each(func(i int) bool {
			return false
		})
	}
}

func Benchmark_LSetRange(b *testing.B) {
	s := getListSet()
	for i := 0; i < b.N; i++ {
		s.Range(func(k int) {
		})
	}
}

// func Benchmark_ZSetRange(b *testing.B) {
// 	s := getZSet()
// 	for i := 0; i < b.N; i++ {
// 		s.Range(0, s.Len(), func(f float64, i int64, a any) {})
// 	}
// }

// ============ Remove ============
func Benchmark_MapSetRemove(b *testing.B) {
	s := getMapSet()
	for i := 0; i < b.N; i++ {
		s.Remove(i)
	}
}

func Benchmark_LSetRemove(b *testing.B) {
	s := getListSet()
	for i := 0; i < b.N; i++ {
		s.Remove(i)
	}
}

// func Benchmark_ZSetRemove(b *testing.B) {
// 	s := getZSet()
// 	for i := 0; i < b.N; i++ {
// 		s.Delete(int64(i))
// 	}
// }

// ============ Add ============
func Benchmark_MapSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMapSet()
	}
}

func Benchmark_LSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getListSet()
	}
}

// func Benchmark_ZSetAdd(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		getZSet()
// 	}
// }

// ============ Union ============
func Benchmark_MapSetUnion(b *testing.B) {
	s1 := getMapSet()
	s2 := getMapSet()
	for i := 0; i < b.N; i++ {
		s1.Union(s2)
	}
}

func Benchmark_LSetUnion(b *testing.B) {
	s1 := getListSet()
	s2 := getListSet()
	for i := 0; i < b.N; i++ {
		s1.Union(s2)
	}
}

// ============ Intersect ============
func Benchmark_MapSetIntersect(b *testing.B) {
	s1 := getMapSet()
	s2 := getMapSet()
	for i := 0; i < b.N; i++ {
		s1.Intersect(s2)
	}
}

func Benchmark_LSetIntersect(b *testing.B) {
	s1 := getListSet()
	s2 := getListSet()
	for i := 0; i < b.N; i++ {
		s1.Intersect(s2)
	}
}
