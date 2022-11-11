package structx

import (
	"fmt"
)

func errOutOfBounds(index int) error {
	return fmt.Errorf("index[%d] out of bounds", index)
}

func errKeyNotFound(key any) error {
	return fmt.Errorf("key[%v] not found", key)
}
