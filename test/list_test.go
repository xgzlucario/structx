package test

import (
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_List1(b *testing.B) {
	ls := structx.NewList[int]()
	for i := 0; i < 9000; i++ {
		ls.LPush(i)
	}
}

func Benchmark_List2(b *testing.B) {
	ls := []int{}
	for i := 0; i < 9000; i++ {
		ls = append([]int{i}, ls...)
	}
}
