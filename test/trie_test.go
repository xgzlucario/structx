package test

import (
	"strconv"
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_TrieSearch(b *testing.B) {
	t := structx.NewTrie[struct{}]()
	for i := 0; i < b.N; i++ {
		t.Insert(strconv.Itoa(i))
	}
}

func Benchmark_MapSearch(b *testing.B) {
	t := structx.NewMap[string, struct{}]()
	for i := 0; i < b.N; i++ {
		t.Set(strconv.Itoa(i), struct{}{})
	}
}

// func Benchmark_TrieTest(b *testing.B) {
// 	t := structx.NewTrie[struct{}]()
// 	t.Insert("12358")
// 	t.Insert("124673")
// 	t.Insert("12458")
// 	t.Insert("1246")

// 	t.Search("1246").PrintChildren()
// 	fmt.Println()
// }
