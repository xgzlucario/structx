package test

import (
	"structx"
	"testing"
)

func Benchmark1(b *testing.B) {
	l1 := structx.NewList[int]()
	for i := 0; i < b.N; i++ {
		if i%10 == 0 {
			l1 = structx.NewList[int]()
		}
		l1.RPush(i)
	}
}

func Benchmark2(b *testing.B) {
	l1 := structx.NewList[int]()
	for i := 0; i < b.N; i++ {
		l1.RPush(i)
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
