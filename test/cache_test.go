package test

import (
	"strconv"
	"testing"

	"github.com/xgzlucario/structx"
)

func getCache() *structx.Cache[string, int] {
	s := structx.NewCache[int]()
	for i := 0; i < NUM; i++ {
		s.Set(strconv.Itoa(i), i)
	}
	return s
}

func Benchmark_CacheSet(b *testing.B) {
	s := structx.NewCache[int]()
	for i := 0; i < b.N; i++ {
		s.Set(strconv.Itoa(i), i)
	}
}

func Benchmark_CacheGet(b *testing.B) {
	s := getCache()
	for i := 0; i < b.N; i++ {
		s.Get(strconv.Itoa(i % NUM))
	}
}

func Benchmark_CacheDelete(b *testing.B) {
	s := getCache()
	for i := 0; i < b.N; i++ {
		s.Delete(strconv.Itoa(i % NUM))
	}
}

func Benchmark_CacheRange(b *testing.B) {
	s := getCache()
	for i := 0; i < b.N; i++ {
		s.Range(func(key string, value int) bool {
			return false
		})
	}
}

// func Benchmark_CacheTable(b *testing.B) {
// 	table := structx.NewCacheTable[string, int]()
// 	for i := 0; i < b.N; i++ {
// 		str := strconv.Itoa(i)
// 		table.Table(str).Store(str, i)
// 	}
// }
