package kl

import (
	"errors"
	"fmt"
	"math/rand"
)

// List is a generic type definition of Slice that provides useful methods and ease of use
type List[T any] []T

func NewList[T any](vals ...T) *List[T] {
	list := List[T](vals)
	return &list
}

func NewListWithCapacity[T any](capacity int, vals ...T) *List[T] {
	list := make(List[T], 0, capacity)
	list = append(list, vals...)
	return &list
}

// Append vals to the end of the list
func (l *List[T]) Append(vals ...T) *List[T] {
	*l = append(*l, vals...)
	return l
}

// Len returns the length of the list
func (l *List[T]) Len() int {
	return len(*l)
}

func (l *List[T]) Capacity() int {
	return cap(*l)
}

func (l *List[T]) String() string {
	return fmt.Sprintf("%v", *l)
}

// Get the item at index i, or false if the index is out of bounds
func (l *List[T]) Get(i int) (T, error) {
	if l.IsEmpty() {
		var zero T
		return zero, errors.New("list is empty")
	}
	if !l.ValidIndex(i) {
		var zero T
		return zero, errors.New("index out of range")
	}
	return (*l)[i], nil
}

// Remove the item at index i
func (l *List[T]) Remove(i int) error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	if !l.ValidIndex(i) {
		return errors.New("index out of range")
	}
	*l = append((*l)[:i], (*l)[i+1:]...)
	return nil
}

// IsEmpty returns true if the list is empty
func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

// ValidIndex checks if the index is within the list bounds
func (l *List[T]) ValidIndex(index int) bool {
	return index >= 0 && index < l.Len()
}

// ValidIndexLoose checks if the index is within the list bounds, allows one index past the end
func (l *List[T]) ValidIndexLoose(index int) bool {
	return index >= 0 && index <= l.Len()
}

// Clear the list
func (l *List[T]) Clear() *List[T] {
	*l = []T{}
	return l
}

// Pop the last item in the list, or the last n items if n is specified
// Returns address of the popped item
func (l *List[T]) Pop() (T, error) {
	if l.IsEmpty() {
		var zero T
		return zero, errors.New("list is empty")
	}

	lastIndex := l.Len() - 1
	item := (*l)[lastIndex]
	*l = (*l)[:lastIndex]
	return item, nil
}

// Dequeue the first item in the list
func (l *List[T]) Dequeue() (T, error) {
	if l.IsEmpty() {
		var zero T
		return zero, errors.New("list is empty")
	}
	item, _ := l.Get(0)
	return item, l.Remove(0)
}

// Reverse reverses the elements of the list in-place
func (l *List[T]) Reverse() {
	if l.IsEmpty() || l.Len() <= 1 {
		return
	}

	arr := *l
	for i, j := 0, l.Len()-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Shuffle randomizes the order of elements in the list (in-place)
func (l *List[T]) Shuffle() *List[T] {
	if l.IsEmpty() || l.Len() <= 1 {
		return l
	}
	arr := *l
	for i := l.Len() - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return l
}

// Insert inserts values at the specified index
func (l *List[T]) Insert(index int, values ...T) error {
	if !l.ValidIndexLoose(index) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	if index == l.Len() {
		l.Append(values...)
		return nil
	}
	prev := (*l)[:index]
	after := (*l)[index:]
	*l = append(prev, values...)
	l.Append(after...)
	return nil
}

// Set replaces the element at the specified index
func (l *List[T]) Set(index int, value T) error {
	if !l.ValidIndex(index) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	(*l)[index] = value
	return nil
}

// Swap exchanges elements at two indices
func (l *List[T]) Swap(i, j int) error {
	if !l.ValidIndex(i) || !l.ValidIndex(j) {
		return fmt.Errorf("indices %d or %d out of bounds", i, j)
	}
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	return nil
}

// Copy returns a new copy of the list
func (l *List[T]) Copy() List[T] {
	if *l == nil {
		return nil
	}
	c := make(List[T], l.Len())
	copy(c, *l)
	return c
}

// Slice returns a new list containing the elements from start (inclusive) to end (exclusive)
func (l *List[T]) Slice(start int, end int) (List[T], error) {
	if !l.ValidIndex(start) || !l.ValidIndexLoose(end) {
		return nil, errors.New("invalid index")
	}
	if end < start {
		return nil, errors.New("end index must be greater than start index")
	}
	return (*l)[start:end], nil
}

func (l *List[T]) Filter(f func(T) bool) List[T] {
	result := *NewList[T]()
	for _, item := range *l {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func (l *List[T]) ThereExists(f func(T) bool) bool {
	for _, item := range *l {
		if f(item) {
			return true
		}
	}
	return false
}

func (l *List[T]) ForAll(f func(T) bool) bool {
	for _, item := range *l {
		if !f(item) {
			return false
		}
	}
	return true
}

// Find returns the first item for which f returns true, or (T, false) if none match
func (l *List[T]) Find(f func(T) bool) (T, bool) {
	for _, item := range *l {
		if f(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func (l *List[T]) ForEach(f func(T)) *List[T] {
	for _, item := range *l {
		f(item)
	}
	return l
}

func (l *List[T]) Chunk(size int) []*List[T] {
	if size <= 0 {
		// Return the whole list as one chunk
		return []*List[T]{l}
	}

	var chunks []*List[T]
	for i := 0; i < len(*l); i += size {
		end := i + size
		if end > len(*l) {
			end = len(*l)
		}

		// Create a new chunk list
		chunk := &List[T]{}
		*chunk = append(*chunk, (*l)[i:end]...)
		chunks = append(chunks, chunk)
	}
	return chunks
}
