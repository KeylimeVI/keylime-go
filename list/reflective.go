package list

import "reflect"

func (l *List[T]) Contains(value T, optionalComparator ...func(A, B T) bool) bool {
	if len(optionalComparator) == 0 {
		return l.Any(func(item T) bool {
			return reflect.DeepEqual(item, value)
		})
	} else {
		return l.containsFunc(value, optionalComparator[0])
	}
}

// Equals compares two lists to determine if they are equal
//
// Warning: uses reflect.DeepEqual, may hinder performance.
// Use List.All with a custom equality function for better performance.
func (l *List[T]) Equals(other []T, optionalComparator ...func(A, B T) bool) bool {
	if len(*l) != len(other) {
		return false
	}
	if len(optionalComparator) == 0 {
		return reflect.DeepEqual([]T(*l), other)
	} else {
		return l.equalsFunc(other, optionalComparator[0])
	}
}

func (l *List[T]) IndexOf(value T, optionalPredicate ...func(T) bool) (int, bool) {
	if len(optionalPredicate) == 0 {
		for i, item := range *l {
			if reflect.DeepEqual(item, value) {
				return i, true
			}
		}
		return -1, false
	} else {
		return l.indexOfFunc(optionalPredicate[0])
	}
}

func (l *List[T]) equalsFunc(other []T, f func(A, B T) bool) bool {
	if len(*l) != len(other) {
		return false
	}
	for i, item := range *l {
		if !f(item, other[i]) {
			return false
		}
	}
	return true
}

func (l *List[T]) containsFunc(value T, f func(A, B T) bool) bool {
	return l.Any(func(item T) bool {
		return f(item, value)
	})
}

func (l *List[T]) indexOfFunc(predicate func(T) bool) (int, bool) {
	for i, item := range *l {
		if predicate(item) {
			return i, true
		}
	}
	return -1, false
}
