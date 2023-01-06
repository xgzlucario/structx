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
	for word >= bm.wordLen() {
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

// AddRange
func (bm *BitMap) AddRange(start uint, end uint) *BitMap {
	for i := start; i < end; i++ {
		bm.Add(i)
	}
	return bm
}

// Remove
func (bm *BitMap) Remove(num uint) bool {
	word, bit := num/bitSize, uint(num%bitSize)
	if word >= bm.wordLen() {
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
	return word < bm.wordLen() && (bm.words[word]&(1<<bit)) != 0
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
func (bm *BitMap) Union(target *BitMap, inplace ...bool) *BitMap {
	// modify inplace
	if len(inplace) > 0 && inplace[0] {
		// append
		for target.wordLen() > bm.wordLen() {
			bm.words = append(bm.words, 0)
		}

		for i, v := range target.words {
			bm.words[i] |= v
		}
		return nil

	} else {
		min, max := bm.compareLength(target)
		// copy max object
		max = max.Copy()

		for i, v := range min.words {
			max.words[i] |= v
		}

		return max
	}
}

// Intersect
func (bm *BitMap) Intersect(target *BitMap, inplace ...bool) *BitMap {
	// modify inplace
	if len(inplace) > 0 && inplace[0] {
		// bm is min
		if bm.wordLen() < target.wordLen() {
			for i := range bm.words {
				bm.words[i] &= target.words[i]
			}

		} else {
			for i, v := range target.words {
				bm.words[i] &= v
			}
			// set 0
			for i := target.wordLen(); i < bm.wordLen(); i++ {
				bm.words[i] &= 0x00
			}
		}
		return nil

	} else {
		min, max := bm.compareLength(target)
		// copy min object
		min = min.Copy()

		for i, v := range max.words {
			if i >= int(min.wordLen()) {
				break
			}
			min.words[i] &= v
		}
		return min
	}
}

// Difference
func (bm *BitMap) Difference(target *BitMap, inplace ...bool) *BitMap {
	// modify inplace
	if len(inplace) > 0 && inplace[0] {
		// append
		for bm.wordLen() < target.wordLen() {
			bm.words = append(bm.words, 0)
		}

		for i, v := range target.words {
			bm.words[i] ^= v
		}
		return nil

	} else {
		min, max := bm.compareLength(target)
		// copy max object
		max = max.Copy()

		// append
		for min.wordLen() < max.wordLen() {
			min.words = append(min.words, 0)
		}

		for i := range max.words {
			max.words[i] ^= min.words[i]
		}
		return max
	}
}

// Len
func (bm *BitMap) Len() uint64 {
	return bm.len
}

// wordLen
func (bm *BitMap) wordLen() uint {
	return uint(len(bm.words))
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
	if bm1.wordLen() < bm2.wordLen() {
		return bm1, bm2
	}
	return bm2, bm1
}
