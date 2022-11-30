package test

import (
	"sort"
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_Sort1(b *testing.B) {
	l := structx.NewList(1, 2, 3)
	for i := 0; i < NUM; i++ {
		l.RPush(i % 32)
	}
	l.SetLess(func(i, j int) bool {
		return l.Index(i) < l.Index(j)
	})

	for i := 0; i < b.N; i++ {
		l.Sort()
	}
}

func Benchmark_Sort2(b *testing.B) {
	l := sort.IntSlice{1}
	for i := 0; i < NUM; i++ {
		l = append(l, i%32)
	}

	for i := 0; i < b.N; i++ {
		l.Sort()
	}
}
