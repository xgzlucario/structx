package test

import (
	"fmt"
	"testing"

	"github.com/xgzlucario/structx"
)

func Benchmark_HMap(b *testing.B) {
	m := structx.NewHMap[string, int]()
	for i := 0; i < NUM; i++ {
		for j := 0; j < NUM; j++ {
			m.Set(fmt.Sprintf("key-%d", i), fmt.Sprintf("field-%d", j), i*j)
		}
	}
}
