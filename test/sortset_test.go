package test

import (
	"fmt"
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_SortSet1(b *testing.B) {
	ss := structx.NewSortSet[string, int]()
	for i := 0; i < b.N; i++ {
		ss.Incr(fmt.Sprintf("%d", i%100), i)
	}
}

func Benchmark_SortSet2(b *testing.B) {
	ss := structx.NewSortSet[string, int]()
	for i := 0; i < b.N; i++ {
		ss.Incr(fmt.Sprintf("%d", i%100), i)
	}
}

// func Benchmark_SortSet3(b *testing.B) {
// 	ss := structx.NewSortSet[string, int]()
// 	ss.Incr("a1", 3)
// 	ss.Print()
// 	ss.Incr("a2", 8)
// 	ss.Print()
// 	ss.Incr("a3", 5)
// 	ss.Print()
// 	ss.Incr("a1", 10)
// 	ss.Print()
// }
