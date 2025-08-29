package kl

// Contains returns true if the list contains the value
func Contains[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value { // ✅ This now works
			return true
		}
	}
	return false
}

func Reduce[T any, U any](list []T, initial U, f func(accumulator U, value T) U) U {
	result := initial
	for _, item := range list {
		result = f(result, item)
	}
	return result
}
