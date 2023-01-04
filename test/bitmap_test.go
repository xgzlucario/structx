package test

import (
	"testing"

	"github.com/xgzlucario/structx"
)

const bitSize = 100000000

func getBitMap() *structx.BitMap {
	bm := structx.NewBitMap()
	for i := 0; i < bitSize; i++ {
		bm.Add(uint(i))
	}
	return bm
}

func Benchmark_BitMapAdd(b *testing.B) {
	bm := structx.NewBitMap()
	for i := 0; i < b.N; i++ {
		bm.Add(uint(i))
	}
}

func Benchmark_BitMapExist(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.Exist(uint(i))
	}
}

func Benchmark_BitMapRemove(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.Remove(uint(i))
	}
}

func Benchmark_BitMapGetMax(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.GetMax()
	}
}

func Benchmark_BitMapGetMin(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.GetMin()
	}
}

func Benchmark_BitMapMarshal(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.Marshal()
	}
}

func Benchmark_BitMapUnmarshal(b *testing.B) {
	bm := getBitMap()
	buf, _ := bm.Marshal()
	for i := 0; i < b.N; i++ {
		bm.Unmarshal(buf)
	}
}
