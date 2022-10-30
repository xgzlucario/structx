package structx

import "structx/base"

type Node[T base.Value] struct {
	Pre  *Node[T]
	Next *Node[T]
	Val  T
}
