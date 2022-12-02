package structx

import (
	"golang.org/x/exp/constraints"
)

type Signed constraints.Signed

type Unsigned constraints.Unsigned

type Integer constraints.Integer

type Float constraints.Float

type Number interface {
	Integer | Float
}

type Value constraints.Ordered
