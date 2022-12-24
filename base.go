package structx

import (
	"bytes"
	"encoding/binary"
	"fmt"

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

func errScanBinaryData() error {
	return fmt.Errorf("error: scan binary data")
}

// =============== Marshal ===============
func marshalBin(data any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)
	return buf.Bytes(), err
}

func unmarshalBin(src []byte, data any) error {
	buf := new(bytes.Buffer)
	buf.Read(src)
	return binary.Read(buf, binary.BigEndian, data)
}
