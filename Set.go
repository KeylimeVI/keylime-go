package kv

import "fmt"

// Set is an alias of List that provides useful methods and ease of use
type Set[T comparable] struct {
	List[T]
}

// NewSet creates a new set from the given values
func NewSet[T comparable](vals ...T) *Set[T] {
	newset := Set[T]{}
	for _, v := range vals {
		newset.Append(v)
	}
	return &newset
}

// Contains returns true if the set contains the value
func (s *Set[T]) Contains(val T) bool {
	for _, v := range s.List {
		if v == val {
			return true
		}
	}
	return false
}

// Append adds the given values to the set, if they are not already present
func (s *Set[T]) Append(vals ...T) *Set[T] {
	for _, v := range vals {
		if !s.Contains(v) {
			s.List.Append(v)
		}
	}
	return s
}

// Extend is an alias for Append
func (s *Set[T]) Extend(vals []T) *Set[T] {
	return s.Append(vals...)
}

func (s *Set[T]) Insert(index int, vals ...T) error {
	if !s.ValidIndexLoose(index) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	for _, v := range vals {
		if !s.Contains(v) {
			_ = s.List.Insert(index, v)
			index++
		}
	}
	return nil
}

func (s *Set[T]) Set(index int, val T) error {
	if s.Contains(val) {
		return nil
	}
	return s.List.Set(index, val)
}
