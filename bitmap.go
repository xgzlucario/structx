package structx

const bitSize = 64 // uint64 is 64 bits

type BitMap struct {
	words []uint64
	len   int
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

	// is exist
	if bm.words[word]&(1<<bit) != 0 {
		bm.words[word] &^= 1 << bit
		bm.len--
		return true
	}

	return false
}

// Exist
func (bm *BitMap) Exist(num uint) bool {
	word, bit := num/bitSize, uint(num%bitSize)
	return word < uint(len(bm.words)) && (bm.words[word]&(1<<bit)) != 0
}

// GetMax
func (bm *BitMap) GetMax() int {
	if bm.len == 0 {
		return -1
	}

	n := len(bm.words) - 1
	word := bm.words[n]
	for i := bitSize - 1; i >= 0; i-- {
		if word&(1<<i) != 0 {
			return bitSize*n + i
		}
	}
	return -1
}

// Len
func (bm *BitMap) Len() int {
	return bm.len
}

// ToSlice
func (bm *BitMap) ToSlice() []uint {
	arr := make([]uint, 0, bm.len)
	for i, v := range bm.words {
		if v == 0 {
			continue
		}
		for j := uint(0); j < bitSize; j++ {
			if v&(1<<j) != 0 {
				arr = append(arr, bitSize*uint(i)+j)
			}
		}
	}
	return arr
}
