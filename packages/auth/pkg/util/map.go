package util

// Map is a higher order function that applies a fun to each element of input by value
// Useful when dealing with pointers or any primitives
func Map[Input, Result any](input []Input, fun func(Input) Result) []Result {
	us := make([]Result, len(input))

	for i := range input {
		us[i] = fun(input[i])
	}

	return us
}

// MapByReference is a higher order function that applies a fun to each element of input by pointer
// Useful when dealing with structs and such when you don't want to map by value
func MapByReference[Input, Result any](input []Input, fun func(*Input) Result) []Result {
	us := make([]Result, len(input))

	for i := range input {
		us[i] = fun(&input[i])
	}

	return us
}
