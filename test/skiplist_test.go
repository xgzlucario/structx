package test

import (
	"testing"

	"github.com/huandu/skiplist"
	"github.com/xgzlucario/structx"
)

func Benchmark_SkipList1(b *testing.B) {
	ls := structx.NewSkipList[byte, int]()
	for i := 0; i < b.N; i++ {
		ls.Add(0, i)
	}
}

func Benchmark_SkipList2(b *testing.B) {
	list := skiplist.New(skiplist.Int)
	for i := 0; i < b.N; i++ {
		list.Set(i, struct{}{})
	}
}

func Benchmark_SkipList3(b *testing.B) {
	ls := structx.NewSkipList[byte, int]()
	ls.Add(0, 2)
	ls.Add(1, 3)
	ls.Add(6, 5)
	ls.Add(4, 8)
	ls.Add(9, 2)
	ls.Add(2, 3)
	ls.Print()
}
