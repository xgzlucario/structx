package test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_SortList2(b *testing.B) {
	s := structx.NewSortList[struct{}, int]()
	for i := 0; i < 10; i++ {
		s.Insert(rand.Intn(20))
	}

	s.Print()
	fmt.Println(s.Index(7))
	s.Delete(5)
	s.Print()
}

// func Benchmark_SortList2(b *testing.B) {
// 	s := structx.NewSortList[struct{}, int]()
// 	for i := 0; i < 10; i++ {
// 		s.Insert(rand.Intn(20))
// 	}
// 	s.Print()
// }

// func Benchmark_SortList3(b *testing.B) {
// 	s := structx.NewSortList[struct{}, int]()
// 	for i := 0; i < 10; i++ {
// 		s.Insert(rand.Intn(20))
// 	}
// 	s.Print()
// }
