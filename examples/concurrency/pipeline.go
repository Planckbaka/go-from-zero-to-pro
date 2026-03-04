package concurrency

// Pipeline is a placeholder for concurrency chapter examples.
func Pipeline(input []int) []int {
	out := make([]int, len(input))
	copy(out, input)
	return out
}
