package test

import (
	"math/rand"
	"testing"

	"github.com/xgzlucario/structx"
)

func Test_SortList2(b *testing.T) {
	s := structx.NewSortList[int]()
	for i := 0; i < 10; i++ {
		s.Insert(rand.Intn(20))
		s.Index(rand.Intn(10))
	}
	s.Len()
	s.Values()
	s.Print()
}

func Benchmark_SortList2(b *testing.B) {
	s := structx.NewSortList[int]()
	for i := 0; i < 10; i++ {
		s.Insert(rand.Intn(20))
	}
	s.Print()
}

func Benchmark_SortList3(b *testing.B) {
	s := structx.NewSortList[int]()
	for i := 0; i < 10; i++ {
		s.Insert(rand.Intn(20))
	}
	s.Print()
}
