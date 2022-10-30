package test

import (
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_SortSet1(b *testing.B) {
	ss := structx.NewSortSet[string, int]()
	ss.Incr("a1", 3)
	ss.Print()
	ss.Incr("a2", 8)
	ss.Print()
	ss.Incr("a3", 5)
	ss.Print()
	ss.Incr("a1", 10)
	ss.Print()
}

func Benchmark_SortSet2(b *testing.B) {
	ls := structx.NewList[int]()
	for i := 0; i < b.N; i++ {
		ls.RPush(i)
	}
}

// func Benchmark3(b *testing.B) {
// 	l1 := util.NewListx[int]()

// 	for i := 0; i < 16; i++ {
// 		fmt.Println(l1)
// 		l1.RPush(i + 1)
// 	}
// 	fmt.Println(l1)
// }
