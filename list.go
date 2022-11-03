package structx

const MAKE_SIZE = 8

type List[T Value] struct {
	Values[T]
}

// NewList: return new List
func NewList[T Value](values ...T) *List[T] {
	if len(values) > 0 {
		return &List[T]{Values: values}
	}
	return &List[T]{
		Values: make(Values[T], 0, MAKE_SIZE),
	}
}

// AddToSet: add to the set
func (ls *List[T]) AddToSet(value T) bool {
	if ls.Index(value) < 0 {
		ls.RPush(value)
		return true
	}
	return false
}

func (ls *List[T]) LPush(value T) {
	ls.RPush(value)
	ls.RShift()
}

func (ls *List[T]) RPush(value T) {
	ls.Values = append(ls.Values, value)
}

func (ls *List[T]) LPop() T {
	ls.LShift()
	return ls.RPop()
}

func (ls *List[T]) Set(index int, value T) {
	ls.Values[index] = value
}

func (ls *List[T]) RPop() T {
	val := ls.Values[ls.Len()-1]
	ls.Values = ls.Values[:ls.Len()-1]
	return val
}

func (ls *List[T]) RemoveElem(elem T) {
	for i, v := range ls.Values {
		if v == elem {
			ls.Values = append(ls.Values[:i], ls.Values[i+1:]...)
			return
		}
	}
}
