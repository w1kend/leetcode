package subsets

func subsets(nums []int) [][]int {
	s := solver{nums: nums}
	return s.cascading()
}

type solver struct {
	nums []int
}

func (s solver) cascading() [][]int {
	res := make([][]int, 0, pow2(len(s.nums)))
	res = append(res, []int{})

	for _, n := range s.nums {
		l := len(res)
		for i := 0; i < l; i++ {
			var set []int
			if len(res[i]) > 0 {
				set = make([]int, len(res[i]), len(res[i])+1)
				copy(set, res[i])
			}
			set = append(set, n)
			res = append(res, set)
		}
	}

	return res
}

func (s solver) backtracking() [][]int {
	return nil
}

func (s solver) lexicographic() [][]int {
	return nil
}

func pow2(n int) int {
	v := 1
	for n > 0 {
		v *= 2
		n--
	}

	return v
}
