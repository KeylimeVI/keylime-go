package kl

// ToSlice converts the list to a native slice
func (l *List[T]) ToSlice() []T {
	slice := make([]T, len(*l))
	copy(slice, *l)
	return slice
}

// ToMap converts the list to a map where the key is the index and the value is the element.
func (l *List[T]) ToMap() map[int]T {
	result := map[int]T{}
	for index, item := range *l {
		result[index] = item
	}
	return result
}
