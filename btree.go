package structx

import "github.com/xgzlucario/structx/base"

type Node[T base.Value] struct {
	Pre  *Node[T]
	Next *Node[T]
	Val  T
}
