package test

import (
	"sort"
	"testing"

	"github.com/xgzlucario/structx"
)

func getList() *structx.List[int] {
	l := structx.NewList[int]()
	for i := 0; i < NUM; i++ {
		l.RPush(i % 32)
	}
	l.SetLess(func(i, j int) bool {
		return i < j
	})
	return l
}

func Benchmark_Sort1(b *testing.B) {
	l := getList()
	for i := 0; i < b.N; i++ {
		l.Sort()
	}
}

func Benchmark_Sort2(b *testing.B) {
	l := sort.IntSlice{}
	for i := 0; i < NUM; i++ {
		l = append(l, i%32)
	}

	for i := 0; i < b.N; i++ {
		l.Sort()
	}
}

func Benchmark_Max(b *testing.B) {
	l := getList()
	for i := 0; i < b.N; i++ {
		l.Max()
	}
}
