package kv

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Set is an alias of []T that behaves like a mathematical set
type Set[T comparable] []T

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
	for _, v := range *s {
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
			*s = append(*s, v)
		}
	}
	return s
}

func (s *Set[T]) Insert(index int, vals ...T) error {
	if !s.ValidIndexLoose(index) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	if index == s.Len() {
		s.Append(vals...)
		return nil
	}
	prev := (*s)[:index]
	after := (*s)[index:]
	*s = append(prev, vals...)
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
	(*s)[index] = val
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
	var intersection Set[T]
	for _, v := range *s {
		if other.Contains(v) {
			intersection.Append(v)
		}
	}
	return intersection
}

func (s *Set[T]) Union(other Set[T]) Set[T] {
	var union Set[T]
	union.Append(*s...)
	union.Append(other...)
	return union
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) Capacity() int {
	return cap(*s)
}

func (s *Set[T]) Get(i int) T {
	if !s.ValidIndex(i) {
		var zero T
		return zero
	}
	return (*s)[i]
}

func (s *Set[T]) Remove(i int) error {
	if !s.ValidIndex(i) {
		return fmt.Errorf("index %d out of bounds", i)
	}
	// Remove the element by creating a new slice without it
	*s = append((*s)[:i], (*s)[i+1:]...)
	return nil
}

func (s *Set[T]) ForEach(f func(T)) {
	for _, item := range *s {
		f(item)
	}
}

func (s *Set[T]) Copy() *Set[T] {
	newSet := make(Set[T], len(*s))
	copy(newSet, *s)
	return &newSet
}

func (s *Set[T]) ThereExists(f func(T) bool) bool {
	for _, item := range *s {
		if f(item) {
			return true
		}
	}
	return false
}

func (s *Set[T]) ForAll(f func(T) bool) bool {
	for _, item := range *s {
		if !f(item) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Slice(start int, end int) (Set[T], error) {
	if !s.ValidIndex(start) || end < start || end > s.Len() {
		return nil, fmt.Errorf("invalid slice range [%d:%d]", start, end)
	}
	return (*s)[start:end], nil
}

func (s *Set[T]) Reverse() {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func (s *Set[T]) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*s), func(i, j int) {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	})
}

func (s *Set[T]) Clear() {
	*s = (*s)[:0]
}

func (s *Set[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Set[")
	for i, item := range *s {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", item))
	}
	sb.WriteString("]")
	return sb.String()
}

func (s *Set[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Set[T]) ValidIndex(i int) bool {
	return i >= 0 && i < len(*s)
}

func (s *Set[T]) ValidIndexLoose(i int) bool {
	return i >= 0 && i <= len(*s)
}

func (s *Set[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("cannot pop from empty set")
	}
	lastIndex := len(*s) - 1
	value := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return value, nil
}

func (s *Set[T]) Dequeue() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("cannot dequeue from empty set")
	}
	value := (*s)[0]
	*s = (*s)[1:]
	return value, nil
}

func (s *Set[T]) Swap(i int, j int) error {
	if !s.ValidIndex(i) || !s.ValidIndex(j) {
		return fmt.Errorf("invalid indices for swap: %d, %d", i, j)
	}
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	return nil
}

func (s *Set[T]) Filter(f func(T) bool) {
	newSet := Set[T]{}
	for _, item := range *s {
		if f(item) {
			newSet = append(newSet, item)
		}
	}
	*s = newSet
}

func (s *Set[T]) Find(f func(T) bool) (int, bool) {
	for i, item := range *s {
		if f(item) {
			return i, true
		}
	}
	return -1, false
}

func (s *Set[T]) Chunk(size int) List[Set[T]] {
	if size <= 0 {
		return List[Set[T]]{*s}
	}

	var chunks List[Set[T]]
	setSlice := *s

	for i := 0; i < len(setSlice); i += size {
		end := i + size
		if end > len(setSlice) {
			end = len(setSlice)
		}

		chunkSet := Set[T](setSlice[i:end])
		chunks = append(chunks, chunkSet)
	}
	return chunks
}
