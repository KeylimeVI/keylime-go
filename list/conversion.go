package kl

// ToSlice converts the list to a native slice
func (l *List[T]) ToSlice() []T {
	slice := make([]T, len(*l))
	copy(slice, *l)
	return slice
}

func (l *List[T]) ToMap() map[int]T {
	result := map[int]T{}
	for index, item := range *l {
		result[index] = item
	}
	return result
}
