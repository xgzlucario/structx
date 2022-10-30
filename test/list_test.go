package test

import (
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_List1(b *testing.B) {
	ls := make([]int, 0)
	for i := 0; i < b.N; i++ {
		ls = append(ls, i)
	}
}

func Benchmark_List2(b *testing.B) {
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
