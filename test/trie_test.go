package test

import (
	"strconv"
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_TrieSearch(b *testing.B) {
	t := structx.NewTrie[struct{}]()
	for i := 0; i < 9999; i++ {
		t.Insert(strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		t.Search("675")
	}
}

func Benchmark_MapSearch(b *testing.B) {
	t := structx.NewMap[string, struct{}]()
	for i := 0; i < 9999; i++ {
		t.Set(strconv.Itoa(i), struct{}{})
	}

	for i := 0; i < b.N; i++ {
		t.Get("675")
	}
}
