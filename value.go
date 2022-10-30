package structx

type Value interface {
	string | float64 | float32 | int64 | int32 | int | uint | byte
}
