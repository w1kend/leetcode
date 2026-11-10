package xmath

type ints interface {
	int | int16 | int32 | int64
}

func Abs[T ints](val T) T {
	if val < 0 {
		return -val
	}

	return val
}
