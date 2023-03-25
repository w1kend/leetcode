package cmpr

type Comparable interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64 | string
}

func Max[V Comparable](a, b V) V {
	if a > b {
		return a
	}
	return b
}

func Min[V Comparable](a, b V) V {
	if a < b {
		return a
	}
	return b
}
