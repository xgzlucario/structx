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
func NewBitMap() *BitMap {
	return new(BitMap)
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
	for i, v := range bm.words {
		if v == 0 {
			continue
		}
		for j := uint(0); j < bitSize; j++ {
			if v&(1<<j) != 0 {
				return int(bitSize*uint(i) + j)
			}
		}
	}
	return -1
}

// Max
func (bm *BitMap) Max() int {
	n := len(bm.words) - 1
	for i := n; i >= 0; i-- {
		v := bm.words[i]
		if v == 0 {
			continue
		}
		for j := bitSize - 1; j >= 0; j-- {
			if v&(1<<j) != 0 {
				return int(bitSize*uint(i) + uint(j))
			}
		}
	}
	return -1
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
