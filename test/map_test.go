package test

import (
	"strconv"
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_Map1(b *testing.B) {
	ls := make(map[string]int)
	for i := 0; i < b.N; i++ {
		ls[strconv.Itoa(i)] = i
	}
}

func Benchmark_Map2(b *testing.B) {
	ls := structx.NewMap[string, int]()
	for i := 0; i < b.N; i++ {
		ls[strconv.Itoa(i)] = i
	}
}

// func Benchmark3(b *testing.B) {
// 	l1 := util.NewListx[int]()

// 	for i := 0; i < 16; i++ {
// 		fmt.Println(l1)
// 		l1.RPush(i + 1)
// 	}
// 	fmt.Println(l1)
// }
