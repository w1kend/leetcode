package maximumerasurevalue

func maximumUniqueSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j := 0, 1

	// track uniq nums
	m := make(map[int]bool, 100)
	m[nums[i]] = true
	max := nums[i]
	sum := max
	for j < len(nums) {
		inum, jnum := nums[i], nums[j]

		if m[jnum] == true {
			// if we already have this num in the window
			// we have to move left border until it(the num) disapeer
			for i < j {
				inum = nums[i]
				m[inum] = false
				i++
				sum -= inum

				if inum == jnum { // if it was needed value
					if i == j {
						// we didn't include num at position j, so have to do it here
						j++
						m[inum] = true
						sum = inum
						if sum > max {
							max = sum
						}
					}
					break
				}
			}

			continue
		}

		m[jnum] = true
		sum += jnum
		if sum > max {
			max = sum
		}
		j++
	}

	return max
}
