package test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_Test1(b *testing.B) {
	l := structx.NewList(1)

	for i := 0; i < 10; i++ {
		l.RPush(i % 6)
	}

	l.SetLess(func(i, j int) bool {
		return l.Index(i) < l.Index(j)
	})

	for i := 0; i < b.N; i++ {
		l.Sort()
	}
	fmt.Println(l)
}

func Benchmark_Test2(b *testing.B) {
	l := sort.IntSlice{1}

	for i := 0; i < 100; i++ {
		l = append(l, i)
	}

	for i := 0; i < b.N; i++ {
		l.Sort()
	}
}
