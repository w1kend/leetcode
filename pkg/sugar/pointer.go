package sugar

// P - returns pointer to the value
func P[T any](v T) *T {
	return &v
}
