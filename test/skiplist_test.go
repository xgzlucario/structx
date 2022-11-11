package test

import (
	"fmt"
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

func Benchmark_SkipList3(b *testing.B) {
	ls := structx.NewSkipList[string, int]()
	ls.Add(123, "xgz")
	ls.Add(12, "xgz1")
	ls.Add(124, "xgz2")
	ls.Add(56, "xgz3")
	ls.Add(199, "xgz4")
	ls.Add(116, "xgz5")
	ls.RangeByScores(60, 200, func(key string, value int) {
		fmt.Println(key, value)
	})
	// ls.Print()
}
