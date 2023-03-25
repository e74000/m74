package m74

// Complex is the set of all complex numbers. Includes int, uint, float, and complex
type Complex interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128
}

// Real is the set of all real numbers. Includes int, uint, and float
type Real interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Integer is the set of integers. Includes, int and uint
type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Natural is the set of all integers >= 0
// They are assumed to include zero
type Natural interface {
	uint | uint8 | uint16 | uint32 | uint64
}
