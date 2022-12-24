package structx

import (
	"fmt"

	"github.com/bytedance/sonic"
	"golang.org/x/exp/constraints"
)

// =============== types ===============
type Signed constraints.Signed

type Unsigned constraints.Unsigned

type Integer constraints.Integer

type Float constraints.Float

type Number interface{ Integer | Float }

type Value constraints.Ordered

// =============== errors ===============
func errOutOfBounds(index int) error {
	return fmt.Errorf("error: index[%d] out of bounds", index)
}

func errKeyNotFound(key any) error {
	return fmt.Errorf("error: key[%v] not found", key)
}

// =============== Marshal ===============
func marshalJSON(data any) ([]byte, error) {
	return sonic.Marshal(data)
}

func unmarshalJSON(src []byte, data any) error {
	return sonic.Unmarshal(src, data)
}
