package test

import (
	"strconv"
	"testing"

	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/xgzlucario/structx"
)

func getCMap() cmap.ConcurrentMap[string, string] {
	m := cmap.New[string]()
	for i := 0; i < NUM; i++ {
		j := strconv.Itoa(i)
		m.Set(j, j)
	}
	return m
}

func getSMap() structx.Map[string, string] {
	m := structx.NewMap[string, string]()
	for i := 0; i < NUM; i++ {
		j := strconv.Itoa(i)
		m.Set(j, j)
	}
	return m
}

func BenchmarkCMapAdd(b *testing.B) {
	m := cmap.New[string]()
	for i := 0; i < b.N; i++ {
		j := strconv.Itoa(i)
		m.Set(j, j)
	}
}

func BenchmarkSMapAdd(b *testing.B) {
	m := cmap.New[string]()
	for i := 0; i < b.N; i++ {
		j := strconv.Itoa(i)
		m.Set(j, j)
	}
}
