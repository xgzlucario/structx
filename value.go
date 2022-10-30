package structx

type Value interface {
	string | float64 | float32 | int64 | int32 | int | uint | byte
}

type AnyValue interface {
	string | float64 | float32 | int64 | int32 | int | uint | byte | any
}

type Number interface {
	float64 | float32 | int64 | int32 | int | uint
}

type Int interface {
	int64 | int32 | int | uint
}

type Float interface {
	float64 | float32
}
