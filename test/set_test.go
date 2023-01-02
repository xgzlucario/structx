package test

import (
	"testing"

	"github.com/xgzlucario/structx"
)

const SUM = 1000

func getListSet() *structx.LSet[int] {
	s := structx.NewLSet[int]()
	for i := 0; i < SUM; i++ {
		s.Add(i)
	}
	return s
}

func Benchmark_Range(b *testing.B) {
	s := getListSet()
	for i := 0; i < b.N; i++ {
		s.Range(func(i int, val int) bool {
			return false
		})
	}
}

func Benchmark_Pop(b *testing.B) {
	s := getListSet()
	for i := 0; i < b.N; i++ {
		if s.Len() > 0 {
			s.RandomPop()
		}
	}
}

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getListSet()
	}
}

func Benchmark_Union(b *testing.B) {
	s1 := getListSet()
	s2 := getListSet()
	for i := 0; i < b.N; i++ {
		s1.Union(s2)
	}
}

func Benchmark_Intersect(b *testing.B) {
	s1 := getListSet()
	s2 := getListSet()
	for i := 0; i < b.N; i++ {
		s1.Intersect(s2)
	}
}

func Benchmark_Diff(b *testing.B) {
	s1 := getListSet()
	s2 := getListSet()
	for i := 0; i < b.N; i++ {
		s1.Difference(s2)
	}
}

func Benchmark_IsSubSet(b *testing.B) {
	s1 := getListSet()
	s2 := getListSet()
	for i := 0; i < b.N; i++ {
		s1.IsSubSet(s2)
	}
}

func Benchmark_Marshal(b *testing.B) {
	s1 := getListSet()
	for i := 0; i < b.N; i++ {
		s1.Marshal()
	}
}
