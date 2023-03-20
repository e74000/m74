package m74

type Complex interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128
}

type Real interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Natural numbers are assumed to include zero
type Natural interface {
	uint | uint8 | uint16 | uint32 | uint64
}
