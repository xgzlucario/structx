package structx

type Value interface {
	string | float32 | float64 | int64 | int32 | int16 | int | uint32 | uint16 | uint | byte
}

type Number interface {
	float32 | float64 | int64 | int32 | int16 | int | uint32 | uint16 | uint | byte
}
