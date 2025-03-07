package utils

// Ptr is function convert Type to *Type
func Ptr[T any](v T) *T {
	return &v
}
