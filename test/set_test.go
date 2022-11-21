package test

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/xgzlucario/structx"
)

const SUM = 28

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
		s.Range(func(k int) bool {
			return false
		})
	}
}

// ============ Pop ============
func Benchmark_MapSetPop(b *testing.B) {
	s := getMapSet()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func Benchmark_LSetRop(b *testing.B) {
	s := getListSet()
	for i := 0; i < b.N; i++ {
		s.RandomPop()
	}
}

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

// ============ Difference ============
func Benchmark_MapSetDiff(b *testing.B) {
	s1 := getMapSet()
	s2 := getMapSet()
	for i := 0; i < b.N; i++ {
		s1.Difference(s2)
	}
}

func Benchmark_LSetDiff(b *testing.B) {
	s1 := getListSet()
	s2 := getListSet()
	for i := 0; i < b.N; i++ {
		s1.Difference(s2)
	}
}

// func BenchmarkTest(b *testing.B) {
// 	s := structx.NewLSet[int]()
// 	s.Add(2)
// 	s.Add(5)
// 	s.Add(9)
// 	s.Add(1)
// 	s.Remove(3)
// 	s.Remove(2)
// 	s.Print()
// }
