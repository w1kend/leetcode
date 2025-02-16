package poweroftwo

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	if n == 2 {
		return true
	}

	return n%2 == 0 && isPowerOfTwo(n/2)
}
