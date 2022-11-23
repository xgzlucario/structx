package test

import (
	"strconv"
	"testing"

	"github.com/patrickmn/go-cache"
	"github.com/xgzlucario/structx"
)

func getCache1() *structx.Cache[string, float64] {
	s := structx.NewCache[string, float64]()
	for i := 0; i < NUM; i++ {
		s.Store(strconv.Itoa(i), float64(i))
	}
	return s
}

func getCache2() *cache.Cache {
	c := cache.New(cache.NoExpiration, structx.GCDuration)
	for i := 0; i < NUM; i++ {
		c.Set(strconv.Itoa(i), float64(i), cache.NoExpiration)
	}
	return c
}

// ========= Store =========
func Benchmark_CacheStore1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache1()
	}
}

func Benchmark_CacheStore2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache2()
	}
}

// ========= Load =========
func Benchmark_CacheLoad1(b *testing.B) {
	s := getCache1()
	for i := 0; i < b.N; i++ {
		s.Load(strconv.Itoa(i % NUM))
	}
}

func Benchmark_CacheLoad2(b *testing.B) {
	s := getCache2()
	for i := 0; i < b.N; i++ {
		s.Get(strconv.Itoa(i % NUM))
	}
}

// ========= Range =========
func Benchmark_CacheRange1(b *testing.B) {
	s := getCache1()
	for i := 0; i < b.N; i++ {
		s.Range(func(key string, value float64) bool {
			return false
		})
	}
}

func Benchmark_CacheRange2(b *testing.B) {
	s := getCache2()
	for i := 0; i < b.N; i++ {
		s.Items()
	}
}
