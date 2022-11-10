package structx

type List[T comparable] struct {
	Array[T]
}

// NewList: return new List
func NewList[T comparable](values ...T) *List[T] {
	return &List[T]{Array: values}
}

func (ls *List[T]) LPush(values ...T) {
	ls.Array = append(values, ls.Array...)
}

func (ls *List[T]) RPush(values ...T) {
	ls.Array = append(ls.Array, values...)
}

func (ls *List[T]) LPop() T {
	val := ls.Array[0]
	ls.Array = ls.Array[1:]
	return val
}

// Insert
func (ls *List[T]) Insert(i int, value T) {
	if i <= 0 {
		ls.LPush(value)

	} else if i >= ls.Len() {
		ls.RPush(value)

	} else {
		ls.Array = append(append(ls.Array[0:i], value), ls.Array[i:]...)
	}
}

// RPop
func (ls *List[T]) RPop() T {
	val := ls.Array[ls.Len()-1]
	ls.Array = ls.Array[:ls.Len()-1]
	return val
}

// RemoveElem
func (ls *List[T]) RemoveElem(elem T) {
	for i, v := range ls.Array {
		if v == elem {
			ls.Array = append(ls.Array[:i], ls.Array[i+1:]...)
			return
		}
	}
}
