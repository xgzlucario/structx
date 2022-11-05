package structx

type List[T comparable] struct {
	Values[T]
}

// NewList: return new List
func NewList[T comparable](values ...T) *List[T] {
	return &List[T]{Values: values}
}

func (ls *List[T]) LPush(values ...T) {
	ls.Values = append(values, ls.Values...)
}

func (ls *List[T]) RPush(values ...T) {
	ls.Values = append(ls.Values, values...)
}

func (ls *List[T]) LPop() T {
	val := ls.Values[0]
	ls.Values = ls.Values[1:]
	return val
}

// Insert
func (ls *List[T]) Insert(i int, value T) {
	if i <= 0 {
		ls.LPush(value)

	} else if i >= ls.Len() {
		ls.RPush(value)

	} else {
		ls.Values = append(append(ls.Values[0:i], value), ls.Values[i:]...)
	}
}

// RPop
func (ls *List[T]) RPop() T {
	val := ls.Values[ls.Len()-1]
	ls.Values = ls.Values[:ls.Len()-1]
	return val
}

// RemoveElem
func (ls *List[T]) RemoveElem(elem T) {
	for i, v := range ls.Values {
		if v == elem {
			ls.Values = append(ls.Values[:i], ls.Values[i+1:]...)
			return
		}
	}
}
