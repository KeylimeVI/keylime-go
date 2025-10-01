package kl

import "reflect"

func (l *List[T]) Contains(value T, optionalComparator ...func(T, T) bool) bool {
	if l.IsEmpty() {
		return false
	}
	if len(optionalComparator) == 0 {
		valueAny := any(value)
		if isComparable(value) {
			for _, item := range *l {
				itemAny := any(item)
				if itemAny == valueAny {
					return true
				}
			}
			return false
		} else {
			return l.Any(func(item T) bool {
				return reflect.DeepEqual(item, value)
			})
		}
	} else {
		return l.containsFunc(value, optionalComparator[0])
	}
}

// Equals compares two lists to determine if they are equal
//
// Warning: uses reflect.DeepEqual, may hinder performance.
// Use List.All with a custom equality function for better performance.
func (l *List[T]) Equals(other List[T], optionalComparator ...func(T, T) bool) bool {
	if len(*l) != len(other) {
		return false
	}
	if l.IsEmpty() && other.IsEmpty() {
		return true
	}
	if len(optionalComparator) == 0 {
		anyValue := any((*l)[0])
		if isComparable(anyValue) {
			for index, item := range *l {
				itemAny := any(item)
				otherAny := any(other[index])
				if itemAny != otherAny {
					return false
				}
			}
			return true
		}
		return reflect.DeepEqual([]T(*l), other)
	} else {
		return l.equalsFunc(other, optionalComparator[0])
	}
}

func (l *List[T]) IndexOf(value T, optionalPredicate ...func(T) bool) (int, bool) {
	if l.IsEmpty() {
		return -1, false
	}
	if len(optionalPredicate) == 0 {
		if isComparable(value) {
			valueAny := any(value)
			for index, item := range *l {
				anyItem := any(item)
				if anyItem == valueAny {
					return index, true
				}
			}
			return -1, false
		}
		for index, item := range *l {
			if reflect.DeepEqual(item, value) {
				return index, true
			}
		}
		return -1, false
	} else {
		return l.indexOfFunc(optionalPredicate[0])
	}
}

func (l *List[T]) equalsFunc(other []T, f func(T, T) bool) bool {
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

func (l *List[T]) containsFunc(value T, f func(T, T) bool) bool {
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

func defaultEquals[T any](a, b T) bool {
	aVal := any(a)
	bVal := any(b)

	// Handle nil cases
	if aVal == nil || bVal == nil {
		return aVal == bVal
	}

	// Check if the type is comparable using reflection
	t := reflect.TypeOf(aVal)
	if t.Comparable() {
		// Safe to use == operator
		return aVal == bVal
	}

	// Fall back to DeepEqual for non-comparable types
	return reflect.DeepEqual(aVal, bVal)
}
