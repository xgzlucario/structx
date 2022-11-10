package test

import (
	"testing"

	"github.com/huandu/skiplist"
	"github.com/xgzlucario/structx"
)

func Benchmark_SkipList1(b *testing.B) {
	ls := structx.NewSkipList[struct{}, int]()
	for i := 0; i < b.N; i++ {
		ls.Add(i)
	}
}

func Benchmark_SkipList2(b *testing.B) {
	list := skiplist.New(skiplist.Int)
	for i := 0; i < b.N; i++ {
		list.Set(i, struct{}{})
	}
}
