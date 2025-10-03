package kl

import (
	"errors"
	"fmt"
	"math/rand"
)

// List is a generic type alias of []T with useful methods and functions
type List[T any] []T

// ListOf creates a new List with the specified values
func ListOf[T any](items ...T) List[T] {
	list := List[T](items)
	return list
}

// ListWithCap creates a new List with the specified capacity and values
func ListWithCap[T any](capacity int, items ...T) List[T] {
	list := make(List[T], 0, capacity)
	list = append(list, items...)
	return list
}

// Add items to the end of the list
func (l *List[T]) Add(items ...T) *List[T] {
	*l = append(*l, items...)
	return l
}

// Len returns the len of the list
func (l *List[T]) Len() int {
	return len(*l)
}

// Cap returns the capacity of the list
func (l *List[T]) Cap() int {
	return cap(*l)
}

// String returns the string representation of the list
func (l *List[T]) String() string {
	return fmt.Sprintf("%v", *l)
}

// Get the item at index i, or false if the index is out of bounds
func (l *List[T]) Get(i int) (T, bool) {
	if l.IsEmpty() {
		var zero T
		return zero, false
	}
	if !l.ValidIndex(i) {
		var zero T
		return zero, false
	}
	return (*l)[i], true
}

// Remove the items at indices, gives up and returns an error if any of the indices are out of bounds
func (l *List[T]) Remove(indices ...int) error {
	if len(indices) == 0 {
		return nil
	}

	if len(indices) == 1 {
		index := indices[0]
		if !l.ValidIndex(index) {
			return NewIndexError(index, l.Len())
		}
		*l = append((*l)[:index], (*l)[index+1:]...)
		return nil
	}

	if !indicesAreFormattedReversed(indices) {
		indices = formatIndicesReversed(indices)
	}

	indicesList := ListOf[int](indices...)

	if !indicesList.All(func(index int) bool {
		return l.ValidIndex(index)
	}) {
		// Return context for the first offending index
		for _, idx := range indicesList {
			if !l.ValidIndex(idx) {
				return NewIndexError(idx, l.Len())
			}
		}
		return IndexOutOfBoundsError // fallback; shouldn't reach
	}

	*l = append((*l)[:indices[0]], (*l)[indices[0]+1:]...)

	return l.Remove(indicesList[1:]...)
}

// RemoveAny removes any items at the specified indices.
// Supports method chaining
func (l *List[T]) RemoveAny(indices ...int) *List[T] {
	indicesList := formatIndicesReversed(indices)
	indicesList.Filter(func(index int) bool {
		return l.ValidIndex(index)
	})
	_ = l.Remove(indicesList...)
	return l
}

// IsEmpty returns true if the list is empty
func (l *List[T]) IsEmpty() bool {
	return len(*l) == 0
}

// ValidIndex checks if the index is within the list bounds
func (l *List[T]) ValidIndex(index int) bool {
	return index >= 0 && index < len(*l)
}

// Clear the list
func (l *List[T]) Clear() *List[T] {
	*l = []T{}
	return l
}

// Pop removes and returns the last item in the list, or the item at index i if specified.
//
// Errors: IndexError, EmptyListError
func (l *List[T]) Pop(i ...int) (T, error) {
	var zero T
	if len(i) == 1 {
		idx := i[0]
		if !l.ValidIndex(idx) {
			return zero, NewIndexError(idx, l.Len())
		}
		item := (*l)[idx]
		return item, nil
	}
	// No index provided: emptiness is the only precondition here
	if l.IsEmpty() {
		return zero, EmptyListError
	}
	return (*l)[len(*l)-1], nil
}

// Concatenate concatenates the inputs to the list
//
// Supports method chaining
func (l *List[T]) Concatenate(lists ...List[T]) *List[T] {
	toConcatenate := Flatten[T, List[T], []List[T]](lists)
	*l = append(*l, toConcatenate...)
	return l
}

// Reverse reverses the order of the list
//
// Supports method chaining
func (l *List[T]) Reverse() *List[T] {
	if l.IsEmpty() || len(*l) <= 1 {
		return l
	}

	arr := *l
	for i, j := 0, len(*l)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return l
}

// Shuffle randomizes the order of the list
//
// Supports method chaining
func (l *List[T]) Shuffle() *List[T] {
	if l.IsEmpty() || len(*l) <= 1 {
		return l
	}
	arr := *l
	for i := len(*l) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return l
}

// Insert values at the specified index
//
// Errors: IndexError
func (l *List[T]) Insert(index int, values ...T) error {
	if !l.validIndexLoose(index) {
		return NewIndexError(index, l.Len())
	}
	if index == len(*l) {
		l.Add(values...)
		return nil
	}
	prev := (*l)[:index]
	after := (*l)[index:]
	*l = append(prev, values...)
	l.Add(after...)
	return nil
}

// Set replaces the element at the specified index with value
//
// Errors: IndexError
func (l *List[T]) Set(index int, value T) error {
	if !l.ValidIndex(index) {
		return NewIndexError(index, l.Len())
	}
	(*l)[index] = value
	return nil
}

// Swap switches the elements at two indices
//
// Errors: IndexError
func (l *List[T]) Swap(i, j int) error {
	if !l.ValidIndex(i) {
		return NewIndexError(i, l.Len())
	}
	if !l.ValidIndex(j) {
		return NewIndexError(j, l.Len())
	}
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	return nil
}

// Copy returns a new deep copy of the list
func (l *List[T]) Copy() List[T] {
	c := make(List[T], len(*l))
	copy(c, *l)
	return c
}

// Slice returns a sublist from start (inclusive) to end (exclusive).
// Errors: IndexError if bounds are invalid.
func (l *List[T]) Slice(start int, end int) (List[T], error) {
	if start < 0 {
		return nil, NewIndexError(start, l.Len())
	}
	if end > l.Len() {
		return nil, NewIndexError(end, l.Len())
	}
	if start > end {
		return nil, errors.New("list.slice: start must be less than or equal to end")
	}
	s := l.Copy()
	return s[start:end], nil
}

// Grow increases the list's capacity by amount
//
// Supports method chaining
func (l *List[T]) Grow(amount int) *List[T] {
	newList := ListWithCap[T](amount+l.Cap(), *l...)
	*l = newList
	return l
}

// Filter keeps elements for which predicate returns true. Supports method chaining.
func (l *List[T]) Filter(predicate func(T) bool) *List[T] {
	toRemove := ListOf[int]()
	for index, item := range *l {
		if !predicate(item) {
			toRemove.Add(index)
		}
	}
	_ = l.Remove(toRemove...)
	return l
}

// Map applies f to each element in place. Supports method chaining.
func (l *List[T]) Map(f func(T) T) *List[T] {
	for i, item := range *l {
		(*l)[i] = f(item)
	}
	return l
}

// FlatMap maps each element to a slice and replaces the list with the concatenated result.
// Supports method chaining.
func (l *List[T]) FlatMap(f func(T) []T) *List[T] {
	result := ListOf[T]()
	for _, item := range *l {
		result.Add(f(item)...)
	}
	*l = result
	return l
}

// ForEach calls f for each element in the list. Supports method chaining.
func (l *List[T]) ForEach(f func(T)) *List[T] {
	for _, item := range *l {
		f(item)
	}
	return l
}

// Any returns true if predicate returns true for any element.
func (l *List[T]) Any(predicate func(T) bool) bool {
	for _, item := range *l {
		if predicate(item) {
			return true
		}
	}
	return false
}

// All returns true if predicate returns true for every element.
func (l *List[T]) All(predicate func(T) bool) bool {
	for _, item := range *l {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// FindBy returns the first item for which f returns true: (item, index, ok)
func (l *List[T]) FindBy(predicate func(T) bool) (T, int, bool) {
	for i, item := range *l {
		if predicate(item) {
			return item, i, true
		}
	}
	var zero T
	return zero, -1, false
}

// Partition splits the list into a slice of chunks of the given size.
// If size <= 0, returns a single chunk copy of the list.
func (l *List[T]) Partition(size int) [][]T {
	if size <= 0 {
		return [][]T{l.Copy()}
	}

	var parts [][]T
	for i := 0; i < len(*l); i += size {
		end := i + size
		if end > len(*l) {
			end = len(*l)
		}

		var part []T
		part = append(part, (*l)[i:end]...)
		parts = append(parts, part)
	}
	return parts
}
