package test

import (
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_List1(b *testing.B) {
	ls := structx.NewList[int]()
	for i := 0; i < b.N; i++ {
		ls.RPush(i)
	}
	structx.Max(ls.Array...)
}
