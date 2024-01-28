package confusingnumber

func confusingNumber(n int) bool {
	rotated := [10]int{0, 1, -1, -1, -1, -1, 9, -1, 8, 6}

	oldNum := n
	num := 0

	for n != 0 {
		c := n % 10

		if rotated[c] == -1 {
			return false
		}

		num = num*10 + rotated[c]
		n /= 10
	}

	return oldNum != num
}
