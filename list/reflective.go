package kl

import "reflect"

func (l *List[T]) Contains(values ...T) bool {
	if l.IsEmpty() {
		return len(values) == 0
	}

	for _, value := range values {
		if !l.singleContains(value) {
			return false
		}
	}
	return true
}

func (l *List[T]) ContainsAny(values ...T) bool {
	if l.IsEmpty() {
		return false
	}

	for _, value := range values {
		if l.singleContains(value) {
			return true
		}
	}
	return false
}

// Equals compares two lists to determine if they are equal
//
// Warning: uses reflect.DeepEqual for non-comparable types, may impact performance.
// Pass in an optional Comparator function for better performance.
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

func (l *List[T]) IndexOf(value T) (int, bool) {
	if l.IsEmpty() {
		return -1, false
	}
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

func (l *List[T]) singleContains(value T) bool {
	if l.IsEmpty() {
		return false
	}
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
}

func (l *List[T]) indexOfFunc(predicate func(T) bool) (int, bool) {
	for i, item := range *l {
		if predicate(item) {
			return i, true
		}
	}
	return -1, false
}
