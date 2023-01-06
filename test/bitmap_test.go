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

func BenchmarkBmAdd(b *testing.B) {
	bm := structx.NewBitMap()
	for i := 0; i < b.N; i++ {
		bm.Add(uint(i))
	}
}

func BenchmarkBmContains(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.Contains(uint(i))
	}
}

func BenchmarkBmRemove(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.Remove(uint(i))
	}
}

func BenchmarkBmMax(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.Max()
	}
}

func BenchmarkBmMin(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.Min()
	}
}

func BenchmarkBmUnion(b *testing.B) {
	bm := structx.NewBitMap().AddRange(0, 1000)
	bm1 := structx.NewBitMap().AddRange(500, 1500)
	for i := 0; i < b.N; i++ {
		bm.Union(bm1)
	}
}

func BenchmarkBmUnionInplace(b *testing.B) {
	bm := structx.NewBitMap().AddRange(0, 1000)
	bm1 := structx.NewBitMap().AddRange(500, 1500)
	for i := 0; i < b.N; i++ {
		bm.Union(bm1, true)
	}
}

func BenchmarkBmIntersect(b *testing.B) {
	bm := structx.NewBitMap().AddRange(0, 1000)
	bm1 := structx.NewBitMap().AddRange(500, 1500)
	for i := 0; i < b.N; i++ {
		bm.Intersect(bm1)
	}
}

func BenchmarkBmIntersectInplace(b *testing.B) {
	bm := structx.NewBitMap().AddRange(0, 1000)
	bm1 := structx.NewBitMap().AddRange(500, 1500)
	for i := 0; i < b.N; i++ {
		bm.Intersect(bm1, true)
	}
}

func BenchmarkBmDifference(b *testing.B) {
	bm := structx.NewBitMap().AddRange(0, 1000)
	bm1 := structx.NewBitMap().AddRange(500, 1500)
	for i := 0; i < b.N; i++ {
		bm.Difference(bm1)
	}
}

func BenchmarkBmDifferenceInplace(b *testing.B) {
	bm := structx.NewBitMap().AddRange(0, 1000)
	bm1 := structx.NewBitMap().AddRange(500, 1500)
	for i := 0; i < b.N; i++ {
		bm.Difference(bm1, true)
	}
}

func BenchmarkBmMarshal(b *testing.B) {
	bm := getBitMap()
	for i := 0; i < b.N; i++ {
		bm.Marshal()
	}
}
