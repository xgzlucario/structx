package structx

import (
	"fmt"

	"github.com/bytedance/sonic"
	"golang.org/x/exp/constraints"
)

// types
type Value constraints.Ordered

// errors
func errOutOfBounds(index int) error {
	return fmt.Errorf("error: index[%d] out of bounds", index)
}

func errKeyNotFound(key any) error {
	return fmt.Errorf("error: key[%v] not found", key)
}

// marshal
func marshalJSON(data any) ([]byte, error) {
	return sonic.Marshal(data)
}

func unmarshalJSON(src []byte, data any) error {
	return sonic.Unmarshal(src, data)
}
