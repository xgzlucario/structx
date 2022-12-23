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

// Set
func (bm *BitMap) Set(num uint) {
	word, bit := num/bitSize, uint(num%bitSize)
	for word >= uint(len(bm.words)) {
		bm.words = append(bm.words, 0)
	}

	// if not exist
	if bm.words[word]&(1<<bit) == 0 {
		bm.words[word] |= 1 << bit
		bm.len++
	}
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
