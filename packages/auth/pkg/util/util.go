package util

// BoolToInt converts a boolean to any type of integer with value 0 or 1
func BoolToInt[IntType int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](input bool) IntType {
	result := IntType(0)

	if input {
		result = 1
	}

	return result
}

// ToPointer converts an rvalue to an lvalue pointer. Very niche.
func ToPointer[T any](input T) *T {
	return &input
}
