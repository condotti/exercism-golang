package collatzconjecture

type CollatzError struct {
	Message string
}

func (e *CollatzError) Error() string {
	return e.Message
}

func CollatzConjecture(n int) (int, error) {
	if n > 0 {
		steps := 0
		for ; n != 1; steps++ {
			if n & 1 == 0 {
				n >>= 1
			} else {
				n = n << 1 + n + 1
			}
		}
		return steps, nil
	}
	return 0, &CollatzError{ "Argument must be positive integer" }
}
