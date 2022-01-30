package internal

// MaybePanic panics if the given error is not nil.
func MaybePanic(err error) {
	if err != nil {
		panic(err)
	}
}
