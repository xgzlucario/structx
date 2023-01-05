package structx

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const bitSize = 64 // uint64 is 64 bit

type BitMap struct {
	words []uint64
	len   uint64
}

// NewBitMap
func NewBitMap(nums ...uint) *BitMap {
	bm := new(BitMap)
	for _, num := range nums {
		bm.Add(num)
	}
	return bm
}

// Add
func (bm *BitMap) Add(num uint) bool {
	word, bit := num/bitSize, uint(num%bitSize)
	for word >= uint(len(bm.words)) {
		bm.words = append(bm.words, 0)
	}

	// if not exist
	if bm.words[word]&(1<<bit) == 0 {
		bm.words[word] |= 1 << bit
		bm.len++
		return true
	}
	return false
}

// Remove
func (bm *BitMap) Remove(num uint) bool {
	word, bit := num/bitSize, uint(num%bitSize)
	if word >= uint(len(bm.words)) {
		return false
	}

	// if exist
	if bm.words[word]&(1<<bit) != 0 {
		bm.words[word] &^= 1 << bit
		bm.len--
		return true
	}
	return false
}

// Contains
func (bm *BitMap) Contains(num uint) bool {
	word, bit := num/bitSize, uint(num%bitSize)
	return word < uint(len(bm.words)) && (bm.words[word]&(1<<bit)) != 0
}

// Min
func (bm *BitMap) Min() int {
	min := -1
	bm.Range(func(u uint) bool {
		min = int(u)
		return true
	})
	return min
}

// Max
func (bm *BitMap) Max() int {
	max := -1
	bm.RevRange(func(u uint) bool {
		max = int(u)
		return true
	})
	return max
}

// Union
func (bm *BitMap) Union(target *BitMap) *BitMap {
	min, max := bm.compareLength(target)
	// should copy max object
	max = max.Copy()

	min.Range(func(v uint) bool {
		max.Add(v)
		return false
	})
	return max
}

// Intersect
func (bm *BitMap) Intersect(target *BitMap) *BitMap {
	min, max := bm.compareLength(target)
	// should copy min object
	min = min.Copy()

	min.Range(func(v uint) bool {
		if !max.Contains(v) {
			min.Remove(v)
		}
		return false
	})
	return min
}

// Difference
func (bm *BitMap) Difference(target *BitMap) *BitMap {
	newBm := NewBitMap()

	bm.Range(func(u uint) bool {
		if !target.Contains(u) {
			newBm.Add(u)
		}
		return false
	})
	target.Range(func(u uint) bool {
		if !bm.Contains(u) {
			newBm.Add(u)
		}
		return false
	})

	return newBm
}

// Len
func (bm *BitMap) Len() uint64 {
	return bm.len
}

// Copy
func (bm *BitMap) Copy() *BitMap {
	return &BitMap{
		words: slices.Clone(bm.words),
		len:   bm.len,
	}
}

// Range: Not recommended for poor performance
func (bm *BitMap) Range(f func(uint) bool) {
	for i, v := range bm.words {
		if v == 0 {
			continue
		}
		for j := uint(0); j < bitSize; j++ {
			if v&(1<<j) != 0 {
				if f(bitSize*uint(i) + j) {
					return
				}
			}
		}
	}
}

// RevRange: Not recommended for poor performance
func (bm *BitMap) RevRange(f func(uint) bool) {
	for i := len(bm.words) - 1; i >= 0; i-- {
		v := bm.words[i]
		if v == 0 {
			continue
		}
		for j := bitSize - 1; j >= 0; j-- {
			if v&(1<<j) != 0 {
				if f(bitSize*uint(i) + uint(j)) {
					return
				}
			}
		}
	}
}

// Marshal
func (bm *BitMap) Marshal() ([]byte, error) {
	return marshalJSON(append(bm.words, bm.len))
}

// Unmarshal
func (bm *BitMap) Unmarshal(src []byte) error {
	if err := unmarshalJSON(src, &bm.words); err != nil {
		return err
	}
	if len(bm.words) == 0 {
		return fmt.Errorf("unmarshal bitmap error")
	}

	n := len(bm.words)
	bm.len = bm.words[n-1]
	bm.words = bm.words[:n-1]
	return nil
}

// Compare two bitmap length and return (*min, *max)
func (bm1 *BitMap) compareLength(bm2 *BitMap) (*BitMap, *BitMap) {
	if bm1.Len() < bm2.Len() {
		return bm1, bm2
	}
	return bm2, bm1
}
