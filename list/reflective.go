package kl

import "reflect"

// Contains returns true if the list contains all of the provided items.
func (l *List[T]) Contains(items ...T) bool {
	if l.IsEmpty() {
		return len(items) == 0
	}

	for _, value := range items {
		if !l.singleContains(value) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the list contains at least one of the provided items.
func (l *List[T]) ContainsAny(items ...T) bool {
	if l.IsEmpty() {
		return false
	}

	for _, value := range items {
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

// IndexOf returns the index of the first occurrence of item and true; otherwise (-1, false).
func (l *List[T]) IndexOf(item T) (int, bool) {
	if l.IsEmpty() {
		return -1, false
	}
	if isComparable(item) {
		valueAny := any(item)
		for index, item := range *l {
			anyItem := any(item)
			if anyItem == valueAny {
				return index, true
			}
		}
		return -1, false
	}
	for index, item := range *l {
		if reflect.DeepEqual(item, item) {
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

func (l *List[T]) singleContains(item T) bool {
	if l.IsEmpty() {
		return false
	}
	valueAny := any(item)
	if isComparable(item) {
		for _, val := range *l {
			itemAny := any(val)
			if itemAny == valueAny {
				return true
			}
		}
		return false
	} else {
		return l.Any(func(val T) bool {
			return reflect.DeepEqual(val, item)
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
