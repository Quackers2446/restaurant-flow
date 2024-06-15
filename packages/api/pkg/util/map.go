package util

// Map is a higher order function that applies a fun to each element of input
func Map[Input, Result any](input []Input, fun func(*Input) Result) []Result {
	us := make([]Result, len(input))

	for i := range input {
		us[i] = fun(&input[i])
	}

	return us
}
