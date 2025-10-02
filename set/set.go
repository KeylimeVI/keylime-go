package ks

import "fmt"

// Set is a generic set of comparable elements implemented as a map[T]struct{}.
type Set[T comparable] map[T]struct{}

// SetOf Constructor - returns Set value
func SetOf[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

// SetWithCap creates a set with the given initial capacity and optional items.
func SetWithCap[T comparable](capacity int, items ...T) Set[T] {
	s := make(Set[T], capacity)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

// Add values to set
func (s *Set[T]) Add(items ...T) *Set[T] {
	for _, item := range items {
		(*s)[item] = struct{}{}
	}
	return s
}

// IsEmpty Check if set is empty
func (s *Set[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Remove items from set
func (s *Set[T]) Remove(items ...T) *Set[T] {
	for _, item := range items {
		delete(*s, item)
	}
	return s
}

// Pop removes and returns an arbitrary element from the set.
// Returns an error if the set is empty.
func (s *Set[T]) Pop() (T, error) {
	for item := range *s {
		delete(*s, item)
		return item, nil
	}
	var zero T
	return zero, fmt.Errorf("pop: empty set")
}

// Contains checks if set contains all items
func (s *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if !s.singleContains(item) {
			return false
		}
	}
	return true
}

// ContainsAny checks if set contains at least one item from items
func (s *Set[T]) ContainsAny(items ...T) bool {
	for _, item := range items {
		if s.singleContains(item) {
			return true
		}
	}
	return false
}

// Len Get size of set
func (s *Set[T]) Len() int {
	return len(*s)
}

// Clear all items from set
func (s *Set[T]) Clear() *Set[T] {
	for item := range *s {
		delete(*s, item)
	}
	return s
}

// Copy the set (returns Set value)
func (s *Set[T]) Copy() Set[T] {
	result := SetOf[T]()
	for item := range *s {
		result[item] = struct{}{}
	}
	return result
}

// String representation (for debugging)
func (s *Set[T]) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}

// Equals returns true if both sets contain exactly the same elements.
func (s *Set[T]) Equals(other Set[T]) bool {
	if len(*s) != len(other) {
		return false
	}
	return s.SupersetOf(other) && s.SubsetOf(other)
}

// SubsetOf Check if this set is a subset of another set
func (s *Set[T]) SubsetOf(other Set[T]) bool {
	if s.IsEmpty() {
		return true // Empty set is subset of any set
	}

	for item := range *s {
		if !other.singleContains(item) {
			return false
		}
	}
	return true
}

// SupersetOf Check if this set is a superset of another set
func (s *Set[T]) SupersetOf(other Set[T]) bool {
	return other.SubsetOf(*s)
}

// Union of two sets
func (s *Set[T]) Union(other Set[T]) Set[T] {
	result := SetOf[T]()
	for item := range *s {
		result[item] = struct{}{}
	}
	for item := range other {
		result[item] = struct{}{}
	}
	return result
}

// Intersection of two sets
func (s *Set[T]) Intersection(other Set[T]) Set[T] {
	result := SetOf[T]()
	// Iterate over the smaller set for efficiency
	if len(*s) < len(other) {
		for item := range *s {
			if other.singleContains(item) {
				result[item] = struct{}{}
			}
		}
	} else {
		for item := range other {
			if s.singleContains(item) {
				result[item] = struct{}{}
			}
		}
	}
	return result
}

// Difference returns a set of elements that are in either s or other but not both.
func (s *Set[T]) Difference(other Set[T]) Set[T] {
	result := SetOf[T]()
	for item := range *s {
		if !other.singleContains(item) {
			result.Add(item)
		}
	}
	for item := range other {
		if !s.singleContains(item) {
			result.Add(item)
		}
	}
	return result
}

// Filter removes elements for which predicate returns false. Supports method chaining.
func (s *Set[T]) Filter(predicate func(T) bool) *Set[T] {
	for item := range *s {
		if !predicate(item) {
			delete(*s, item)
		}
	}
	return s
}

func (s *Set[T]) Map(f func(T) T) *Set[T] {
	result := SetOf[T]()
	for item := range *s {
		applied := f(item)
		result.Add(applied)
	}
	*s = result
	return s
}

func (s *Set[T]) FlatMap(f func(T) []T) *Set[T] {
	result := SetOf[T]()
	for item := range *s {
		applied := f(item)
		result.Add(applied...)
	}
	*s = result
	return s
}

func (s *Set[T]) ForEach(f func(T)) *Set[T] {
	for item := range *s {
		f(item)
	}
	return s
}

func (s *Set[T]) Any(predicate func(T) bool) bool {
	for item := range *s {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (s *Set[T]) All(predicate func(T) bool) bool {
	for item := range *s {
		if !predicate(item) {
			return false
		}
	}
	return true
}
