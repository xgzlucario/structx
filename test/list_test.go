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
	return l
}

func getSlice() sort.IntSlice {
	l := sort.IntSlice{}
	for i := 0; i < NUM; i++ {
		l = append(l, i%32)
	}
	return l
}

func Benchmark_Sort1(b *testing.B) {
	l := getList()
	for i := 0; i < b.N; i++ {
		l.Sort(func(i, j int) bool {
			return i < j
		})
	}
}

func Benchmark_Sort2(b *testing.B) {
	l := getSlice()
	for i := 0; i < b.N; i++ {
		l.Sort()
	}
}

func Benchmark_Max(b *testing.B) {
	l := getList()
	for i := 0; i < b.N; i++ {
		l.Max(func(i, j int) bool {
			return i < j
		})
	}
}
