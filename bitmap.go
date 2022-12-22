package structx

type BitMap []byte

const byteSize = 8 // byte is 8 bit

// NewBitMap
func NewBitMap(n uint) BitMap {
	return make([]byte, n/byteSize+1)
}

// Set
func (bm BitMap) Set(n uint) {
	if n/byteSize > uint(len(bm)) {
		return
	}
	bm[n/byteSize] |= 1 << (n % byteSize)
}

// Delete
func (bm BitMap) Delete(n uint) {
	if n/byteSize > uint(len(bm)) {
		return
	}
	bm[n/byteSize] &= 0 << (n % byteSize)
}

// Exist
func (bm BitMap) Exist(n uint) bool {
	if n/byteSize > uint(len(bm)) {
		return false
	}
	return bm[n/byteSize]&(1<<(n%byteSize)) != 0
}
