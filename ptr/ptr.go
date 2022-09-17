package ptr

// To returns a pointer to v
func To[T any](v T) *T {
	p := &v
	return p
}

// From dereferences v if possible, else returning the zero value
func From[T any](v *T) T {
	if v != nil {
		return *v
	}

	var zero T
	return zero
}
