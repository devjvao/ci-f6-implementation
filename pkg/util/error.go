package util

// Ignore ignores the error from a function
func Ignore(f func() error) {
	_ = f()
}
