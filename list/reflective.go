package list

import "reflect"

func (l *List[T]) Contains(value T) bool {
	return l.Any(func(item T) bool {
		return reflect.DeepEqual(item, value)
	})
}

// Equals compares two lists to determine if they are equal
//
// Warning: uses reflect.DeepEqual, may hinder performance.
// Use List.All with a custom equality function for better performance.
func (l *List[T]) Equals(other []T) bool {
	if len(*l) != len(other) {
		return false
	}
	return reflect.DeepEqual([]T(*l), other)
}

func (l *List[T]) IndexOf(value T) (int, bool) {
	for i, item := range *l {
		if reflect.DeepEqual(item, value) {
			return i, true
		}
	}
	return -1, false
}
