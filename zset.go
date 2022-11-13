package structx

type zslNode[K, V Value] struct {
	key   K
	value V
}

type ZSet[K, V Value] struct {
	zsl *Skiplist[K, V]
	m   Map[K, *zslNode[K, V]]
}

// NewZSet
func NewZSet[K, V Value]() *ZSet[K, V] {
	return &ZSet[K, V]{
		zsl: NewSkipList[K, V](),
		m:   Map[K, *zslNode[K, V]]{},
	}
}

// Set: set key and value
func (z *ZSet[K, V]) Set(key K, value V) bool {
	n, ok := z.m[key]
	if ok {
		// value not change
		if value == n.value {
			return false
		}
		n.value = value
		z.zsl.Delete(key, n.value)
		z.zsl.Add(key, n.value)

	} else {
		z.insertNode(key, value)
	}
	return true
}

// Incr: Increment value by key
func (z *ZSet[K, V]) Incr(key K, value V) V {
	n, ok := z.m[key]
	// not exist
	if !ok {
		z.insertNode(key, value)
		return value
	}
	// exist
	z.zsl.Delete(key, n.value)
	n.value += value
	z.zsl.Add(key, n.value)

	return n.value
}

// Delete: delete keys
func (z *ZSet[K, V]) Delete(keys ...K) error {
	for _, key := range keys {
		n, ok := z.m[key]
		if !ok {
			return errKeyNotFound(key)
		}
		z.deleteNode(n.key, n.value)
	}
	return nil
}

// GetByRank: get value by rank
func (z *ZSet[K, V]) GetByRank(rank int) (k K, v V, err error) {
	if rank < 0 || rank > z.Len() {
		return k, v, errOutOfBounds(rank)
	}
	return z.zsl.GetByRank(rank)
}

// GetScore
func (z *ZSet[K, V]) GetScore(key K) (v V, err error) {
	node, ok := z.m[key]
	if !ok {
		return v, errKeyNotFound(key)
	}
	return node.value, nil
}

// Copy
func (z *ZSet[K, V]) Copy() *ZSet[K, V] {
	newZSet := NewZSet[K, V]()
	z.Range(0, -1, func(key K, value V) bool {
		return newZSet.Set(key, value)
	})
	return z
}

// Union
func (z *ZSet[K, V]) Union(target *ZSet[K, V]) {
	target.Range(0, -1, func(key K, value V) bool {
		z.Incr(key, value)
		return false
	})
}

// Range
func (z *ZSet[K, V]) Range(start, end int, f func(key K, value V) bool) {
	z.zsl.Range(start, end, f)
}

// RangeByScores
func (z *ZSet[K, V]) RangeByScores(min, max V, f func(key K, value V)) {
	z.zsl.RangeByScores(min, max, f)
}

func (z *ZSet[K, V]) Len() int {
	return len(z.m)
}

// make sure that key is not exist!
func (z *ZSet[K, V]) insertNode(key K, value V) *skiplistNode[K, V] {
	z.m[key] = &zslNode[K, V]{
		key:   key,
		value: value,
	}
	return z.zsl.Add(key, value)
}

// make sure that key exist!
func (z *ZSet[K, V]) deleteNode(key K, value V) {
	delete(z.m, key)
	z.zsl.Delete(key, value)
}

// DEBUG
func (z *ZSet[K, V]) Print() {
	z.zsl.Print()
}
