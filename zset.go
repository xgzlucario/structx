package structx

type zslNode[K comparable, V Value] struct {
	key   K
	value V
}

func (z *zslNode[K, V]) Key() K {
	return z.key
}

func (z *zslNode[K, V]) Value() V {
	return z.value
}

type ZSet[K comparable, V Value] struct {
	zsl *Skiplist[K, V]
	m   Map[K, *zslNode[K, V]]
}

// NewZSet
func NewZSet[K comparable, V Value]() *ZSet[K, V] {
	return &ZSet[K, V]{
		zsl: NewSkipList[K, V](),
		m:   Map[K, *zslNode[K, V]]{},
	}
}

// Set: set key and value
func (z *ZSet[K, V]) Set(key K, value V) bool {
	n, ok := z.m[key]

	// value not change
	if value == n.value {
		return false
	}
	if ok {
		n.value = value
		z.zsl.Delete(n.value, key)
		z.zsl.Add(n.value, key)

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
	z.zsl.Delete(n.value, key)
	n.value += value
	z.zsl.Add(n.value, key)

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
func (z *ZSet[K, V]) GetByRank(rank int) (K, V, error) {
	if rank < 0 || rank > z.Len() {
		var k K
		var v V
		return k, v, errOutOfBounds(rank)
	}
	return z.zsl.GetByRank(rank)
}

// GetScore
func (z *ZSet[K, V]) GetScore(key K) (V, error) {
	node, ok := z.m[key]
	if !ok {
		var v V
		return v, errKeyNotFound(key)
	}
	return node.value, nil
}

// Copy
func (z *ZSet[K, V]) Copy() *ZSet[K, V] {
	newZSet := NewZSet[K, V]()
	z.Range(0, -1, func(key K, value V) {
		newZSet.Set(key, value)
	})
	return z
}

// Union
func (z *ZSet[K, V]) Union(target *ZSet[K, V]) {
	target.Range(0, -1, func(key K, value V) {
		z.Incr(key, value)
	})
}

// Range
func (z *ZSet[K, V]) Range(start, end int, f func(key K, value V)) {
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
	return z.zsl.Add(value, key)
}

// make sure that key exist!
func (z *ZSet[K, V]) deleteNode(key K, value V) {
	delete(z.m, key)
	z.zsl.Delete(value, key)
}

// DEBUG
func (z *ZSet[K, V]) Print() {
	z.zsl.Print()
}
