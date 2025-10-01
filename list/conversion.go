package list

// ToSlice converts the list to a native slice
func (l *List[T]) ToSlice() []T {
	slice := make([]T, len(*l))
	copy(slice, *l)
	return slice
}
