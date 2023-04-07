package search2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	rowN := searchBinary(len(matrix), func(i int) compareResult {
		arr := matrix[i]
		first, last := arr[0], arr[len(arr)-1]

		// if target is inside an array
		if target >= first && target <= last {
			return equal
		}

		if target < first {
			return less
		}

		return more
	})

	if rowN == -1 {
		return false
	}

	arrN := searchBinary(len(matrix[rowN]), func(i int) compareResult {
		elem := matrix[rowN][i]
		if elem == target {
			return equal
		}
		if target < elem {
			return less
		}

		return more
	})

	return arrN != -1
}

type comparer func(idx int) compareResult

type compareResult int8

const (
	equal compareResult = 0
	less  compareResult = 1
	more  compareResult = 2
)

func searchBinary(length int, comparer comparer) (index int) {
	l := 0
	r := length - 1

	for l <= r {
		mid := (l + r) / 2

		cmp := comparer(mid)
		switch cmp {
		case equal:
			return mid
		case less:
			r = mid - 1
		case more:
			l = mid + 1
		}
	}

	return -1
}
