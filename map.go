package structx

type Map[K Value, V AnyValue] map[K]V

func NewMap[K Value, V AnyValue]() Map[K, V] {
	return make(Map[K, V], MAKE_SIZE)
}

func (m Map[K, V]) Store(k K, v V) {
	m[k] = v
}

func (m Map[K, V]) Load(k K) V {
	return m[k]
}

func (m Map[K, V]) Delete(key K) {
	delete(m, key)
}

func (m Map[K, V]) Range(f func(k K, v V)) {
	for k, v := range m {
		f(k, v)
	}
}

func (m Map[K, V]) Len() int {
	return len(m)
}
