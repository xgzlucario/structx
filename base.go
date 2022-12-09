package structx

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// types
type Signed constraints.Signed

type Unsigned constraints.Unsigned

type Integer constraints.Integer

type Float constraints.Float

type Number interface{ Integer | Float }

type Value constraints.Ordered

type Function interface {
	~func(...string) | ~func(string) |
		~func(...int) | ~func(int) |
		~func(...int32) | ~func(int32) |
		~func(...int64) | ~func(int64) |
		~func(...float32) | ~func(float32) |
		~func(...float64) | ~func(float64) |
		~func(...any) | ~func(any)
}

// errors
func errOutOfBounds(index int) error {
	return fmt.Errorf("index[%d] out of bounds", index)
}

func errKeyNotFound(key any) error {
	return fmt.Errorf("key[%v] not found", key)
}
