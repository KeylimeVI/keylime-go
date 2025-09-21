package set

import (
	"fmt"
	"github.com/KeylimeVI/kl/list"
)

type Set[T comparable] map[T]struct{}

// NewSet Constructor - returns Set value
func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func NewSetWithCapacity[T comparable](capacity int, items ...T) Set[T] {
	s := make(Set[T], capacity)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

// IsEmpty Check if set is empty
func (s *Set[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Add items to set
func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		(*s)[item] = struct{}{}
	}
}

// Remove items from set
func (s *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(*s, item)
	}
}

// Contains Check if set contains item
func (s *Set[T]) Contains(item T) bool {
	_, exists := (*s)[item]
	return exists
}

// ContainsAll Check if set contains all items
func (s *Set[T]) ContainsAll(items ...T) bool {
	for _, item := range items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

// Len Get size of set
func (s *Set[T]) Len() int {
	return len(*s)
}

// Clear all items from set
func (s *Set[T]) Clear() {
	for item := range *s {
		delete(*s, item)
	}
}

// ToSlice Convert to slice
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(*s))
	for item := range *s {
		slice = append(slice, item)
	}
	return slice
}

func (s *Set[T]) ToList() list.List[T] {
	return list.NewList[T](s.ToSlice()...)
}

// SubsetOf Check if this set is a subset of another set
func (s *Set[T]) SubsetOf(other *Set[T]) bool {
	if s.IsEmpty() {
		return true // Empty set is subset of any set
	}

	for item := range *s {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// SupersetOf Check if this set is a superset of another set
func (s *Set[T]) SupersetOf(other *Set[T]) bool {
	return other.SubsetOf(s)
}

// Union of two sets (returns Set value)
func (s *Set[T]) Union(other *Set[T]) Set[T] {
	result := NewSet[T]()
	for item := range *s {
		result[item] = struct{}{}
	}
	for item := range *other {
		result[item] = struct{}{}
	}
	return result
}

// Intersection of two sets (returns Set value)
func (s *Set[T]) Intersection(other *Set[T]) Set[T] {
	result := NewSet[T]()
	// Iterate over the smaller set for efficiency
	if len(*s) < len(*other) {
		for item := range *s {
			if other.Contains(item) {
				result[item] = struct{}{}
			}
		}
	} else {
		for item := range *other {
			if s.Contains(item) {
				result[item] = struct{}{}
			}
		}
	}
	return result
}

// Equals Check if two sets are equal
func (s *Set[T]) Equals(other *Set[T]) bool {
	if len(*s) != len(*other) {
		return false
	}
	return s.SubsetOf(other)
}

// Clone the set (returns Set value)
func (s *Set[T]) Clone() Set[T] {
	result := NewSet[T]()
	for item := range *s {
		result[item] = struct{}{}
	}
	return result
}

// String representation (for debugging)
func (s *Set[T]) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
