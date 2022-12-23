package test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/xgzlucario/structx"
)

func getCache() *structx.Cache[string, int] {
	s := structx.NewCache[string, int]()
	for i := 0; i < NUM; i++ {
		s.Store(strconv.Itoa(i), i)
	}
	return s
}

func Benchmark_CacheSet(b *testing.B) {
	s := structx.NewCache[string, int]()
	for i := 0; i < b.N; i++ {
		s.Store(strconv.Itoa(i), i)
	}
}

func Benchmark_CacheGet(b *testing.B) {
	s := getCache()
	for i := 0; i < b.N; i++ {
		s.Load(strconv.Itoa(i % NUM))
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

func Benchmark_CacheTable(b *testing.B) {
	table := structx.NewCacheTable[string, int]()
	for i := 0; i < b.N; i++ {
		str := strconv.Itoa(i)
		table.Table(str).Store(str, i)
	}
}

func Benchmark_BitMapTest(b *testing.B) {
	bm := structx.NewBitMap()
	bm.Set(22)
	bm.Set(5)
	bm.Set(8)
	bm.Set(4)
	fmt.Println(bm.ToSlice())
}
