package kl

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

// Set is an alias of []T that behaves like a mathematical set
type Set[T comparable] struct {
	list List[T]
}

// NewSet creates a new set from the given values
func NewSet[T comparable](vals ...T) Set[T] {
	newList := NewListWithCapacity[T](len(vals))
	for _, v := range vals {
		newList.Append(v)
	}
	newSet := Set[T]{newList}
	return newSet
}

// Contains returns true if the set contains the value
func (s *Set[T]) Contains(val T) bool {
	for _, v := range s.list {
		if v == val {
			return true
		}
	}
	return false
}

// Append adds the given values to the set if they are not already present
func (s *Set[T]) Append(vals ...T) *Set[T] {
	for _, v := range vals {
		if !s.Contains(v) {
			s.list.Append(v)
		}
	}
	return s
}

func (s *Set[T]) Extend(vals []T) *Set[T] {
	return s.Append(vals...)
}

func (s *Set[T]) Insert(index int, vals ...T) error {
	if !s.ValidIndexLoose(index) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	if index == s.Len() {
		s.Append(vals...)
		return nil
	}
	prev := (s.list)[:index]
	after := (s.list)[index:]
	s.list = append(prev, vals...)
	s.Append(after...)
	return nil
}

func (s *Set[T]) Set(index int, val T) error {
	if s.Contains(val) {
		return nil
	}
	if !s.ValidIndex(index) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	(s.list)[index] = val
	return nil
}

func (s *Set[T]) SupersetOf(other Set[T]) bool {
	if other.ForAll(func(v T) bool { return s.Contains(v) }) {
		return true
	}
	return false
}

func (s *Set[T]) SubsetOf(other Set[T]) bool {
	if s.ForAll(func(v T) bool { return other.Contains(v) }) {
		return true
	}
	return false
}

func (s *Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := NewSet[T]()
	for _, v := range s.list {
		if other.Contains(v) {
			intersection.Append(v)
		}
	}
	return intersection
}

func (s *Set[T]) Union(other Set[T]) Set[T] {
	union := NewSet[T]()
	union.Append(s.list...)
	union.Append(other.list...)
	return union
}

func (s *Set[T]) Len() int {
	return len(s.list)
}

func (s *Set[T]) Capacity() int {
	return cap(s.list)
}

func (s *Set[T]) Get(i int) (T, error) {
	if !s.ValidIndex(i) {
		var zero T
		return zero, errors.New("index out of range")
	}
	return (s.list)[i], nil
}

func (s *Set[T]) Remove(i int) error {
	if !s.ValidIndex(i) {
		return fmt.Errorf("index %d out of bounds", i)
	}
	// Remove the element by creating a new slice without it
	s.list = append((s.list)[:i], (s.list)[i+1:]...)
	return nil
}

func (s *Set[T]) ForEach(f func(T)) {
	for _, item := range s.list {
		f(item)
	}
}

func (s *Set[T]) Copy() Set[T] {
	newSet := NewSet[T](s.list...)
	copy(newSet.list, s.list)
	return newSet
}

func (s *Set[T]) ThereExists(f func(T) bool) bool {
	for _, item := range s.list {
		if f(item) {
			return true
		}
	}
	return false
}

func (s *Set[T]) ForAll(f func(T) bool) bool {
	for _, item := range s.list {
		if !f(item) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Slice(start int, end int) (Set[T], error) {
	if !s.ValidIndex(start) || end < start || end > s.Len() {
		var zero Set[T]
		return zero, fmt.Errorf("invalid slice range [%d:%d]", start, end)
	}
	newSet := NewSet[T]()
	newSet.list = (s.list)[start:end]
	return newSet, nil
}

func (s *Set[T]) Reverse() {
	for i, j := 0, len(s.list)-1; i < j; i, j = i+1, j-1 {
		(s.list)[i], (s.list)[j] = (s.list)[j], (s.list)[i]
	}
}

func (s *Set[T]) Shuffle() {
	for i := range s.list {
		j := rand.Intn(i + 1)
		(s.list)[i], (s.list)[j] = (s.list)[j], (s.list)[i]
	}
}

func (s *Set[T]) Clear() {
	s.list = (s.list)[:0]
}

func (s *Set[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Set[")
	for i, item := range s.list {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", item))
	}
	sb.WriteString("]")
	return sb.String()
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *Set[T]) ValidIndex(i int) bool {
	return i >= 0 && i < len(s.list)
}

func (s *Set[T]) ValidIndexLoose(i int) bool {
	return i >= 0 && i <= len(s.list)
}

func (s *Set[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("cannot pop from empty set")
	}
	lastIndex := len(s.list) - 1
	value := (s.list)[lastIndex]
	s.list = (s.list)[:lastIndex]
	return value, nil
}

func (s *Set[T]) Swap(i int, j int) error {
	if !s.ValidIndex(i) || !s.ValidIndex(j) {
		return fmt.Errorf("invalid indices for swap: %d, %d", i, j)
	}
	(s.list)[i], (s.list)[j] = (s.list)[j], (s.list)[i]
	return nil
}

func (s *Set[T]) Filter(f func(T) bool) Set[T] {
	newSet := Set[T]{}
	for _, item := range s.list {
		if f(item) {
			newSet.Append(item)
		}
	}
	return newSet
}

func (s *Set[T]) Find(f func(T) bool) int {
	for i, item := range s.list {
		if f(item) {
			return i
		}
	}
	return -1
}

func (s *Set[T]) Chunk(size int) List[Set[T]] {
	chunkOfLists := s.list.Chunk(size)
	chunkOfSets := NewList[Set[T]]()
	for _, list := range chunkOfLists {
		chunkOfSets.Append(NewSet[T](list...))
	}
	return chunkOfSets
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, len(s.list))
	copy(slice, s.list)
	return slice
}

func (s *Set[T]) ToList() List[T] {
	return NewList[T](s.list...)
}

func (s *Set[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, item := range s.list {
			ch <- item
		}
	}()
	return ch
}
