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

func errLengthNotEqual(len1, len2 int) error {
	return fmt.Errorf("length not equal: [%d] [%d]", len1, len2)
}