package structx

import "fmt"

type zslNode[K comparable, V Value] struct {
	key   K
	value V
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

// Delete: delete node by key
func (z *ZSet[K, V]) Delete(key K) bool {
	n, ok := z.m[key]
	if ok {
		z.deleteNode(n.key, n.value)
	}
	return ok
}

// GetByRank: get value by rank
func (z *ZSet[K, V]) GetByRank(rank int) V {
	p := z.zsl.head.forward[rank]
	return p.value
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
	// add zsl node
	return z.zsl.Add(value, key)
}

// make sure that key exist!
func (z *ZSet[K, V]) deleteNode(key K, value V) {
	delete(z.m, key)
	// delete zsl node
	z.zsl.Delete(value, key)
}

func (z *ZSet[K, V]) Print() {
	for k, v := range z.m {
		fmt.Println(k, v.value)
	}
}
