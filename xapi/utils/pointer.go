package utils

// Pointer gets a pointer to given argument
func Pointer[T any](value T) *T {
	return &value
}
