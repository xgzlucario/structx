package test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/xgzlucario/structx"
)

const TrieNum = 99999

func getTrie() *structx.Trie[struct{}] {
	t := structx.NewTrie[struct{}]()
	for i := 0; i < TrieNum; i++ {
		t.Insert(strconv.Itoa(i))
	}
	return t
}

func getMap() structx.Map[string, struct{}] {
	t := structx.NewMap[string, struct{}]()
	for i := 0; i < TrieNum; i++ {
		t.Set(strconv.Itoa(i), struct{}{})
	}
	return t
}

// Add
func Benchmark_TrieAdd(b *testing.B) {
	t := getTrie()
	for i := 0; i < b.N; i++ {
		t.Insert(strconv.Itoa(i))
	}
}

func Benchmark_MapAdd(b *testing.B) {
	t := getMap()
	for i := 0; i < b.N; i++ {
		t.Set(strconv.Itoa(i), struct{}{})
	}
}

// Search
func Benchmark_TrieSearch(b *testing.B) {
	t := getTrie()
	for i := 0; i < b.N; i++ {
		t.Search(strconv.Itoa(i))
	}
}

func Benchmark_MapSearch(b *testing.B) {
	t := getMap()
	for i := 0; i < b.N; i++ {
		t.Get(strconv.Itoa(i))
	}
}

// Delete
func Benchmark_TrieDelete(b *testing.B) {
	t := getTrie()
	for i := 0; i < b.N; i++ {
		t.Delete(strconv.Itoa(i))
	}
}

func Benchmark_MapDelete(b *testing.B) {
	t := getMap()
	for i := 0; i < b.N; i++ {
		t.Delete(strconv.Itoa(i))
	}
}

func Benchmark_TrieTest(b *testing.B) {
	t := structx.NewTrie[int]()
	t.Insert("12358")
	t.Insert("124673")
	t.Insert("12458")
	t.Insert("12458", 1)
	t.Insert("12458", 2)
	t.Insert("12458", 3)
	t.Insert("1249")
	t.Insert("1243")

	t.Search("1245").PrintChildren()
	fmt.Println()
}
