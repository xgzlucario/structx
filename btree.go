package structx

type Node[T Value] struct {
	Pre  *Node[T]
	Next *Node[T]
	Val  T
}
