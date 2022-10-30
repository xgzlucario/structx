package structx

type Chan[T Value] chan T

func NewChan[T Value](size ...int) *Chan[T] {
	if len(size) > 0 {
		ch := make(Chan[T], size[0])
		return &ch
	}
	ch := make(Chan[T])
	return &ch
}
