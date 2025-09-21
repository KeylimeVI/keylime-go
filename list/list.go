package list

import (
	"fmt"
	"math/rand"
)

// List is a generic type definition of Slice that provides useful methods and ease of use
type List[T any] []T

func NewList[T any](vals ...T) List[T] {
	list := List[T](vals)
	return list
}

// NewListWithCapacity creates a new list with the specified capacity
func NewListWithCapacity[T any](capacity int, vals ...T) List[T] {
	list := make(List[T], 0, capacity)
	list = append(list, vals...)
	return list
}

// Add vals to the end of the list
func (l *List[T]) Add(vals ...T) *List[T] {
	*l = append(*l, vals...)
	return l
}

func (l *List[T]) Added(vals ...T) List[T] {
	return append(*l, vals...)
}

// Len returns the length of the list
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

// Remove the items at indices, gives up and returns false if any of the indices are out of bounds
func (l *List[T]) Remove(indices ...int) bool {

	if len(indices) == 0 {
		return true
	}

	if l.IsEmpty() {
		return false
	}

	if len(indices) == 1 {
		index := indices[0]

		if !l.ValidIndex(index) {
			return false
		}
		*l = append((*l)[:index], (*l)[index+1:]...)
		return true
	}

	if !indicesAreFormattedReversed(indices) {
		indices = formatIndicesReversed(indices)
	}

	indicesList := NewList[int](indices...)

	if !indicesList.All(func(index int) bool {
		return l.ValidIndex(index)
	}) {
		return false
	}

	*l = append((*l)[:indices[0]], (*l)[indices[0]+1:]...)

	return l.Remove(indicesList[1:]...)
}

func (l *List[T]) RemoveAny(indices ...int) *List[T] {
	indicesList := formatIndicesReversed(indices)
	indicesList = indicesList.Filtered(func(index int) bool {
		return l.ValidIndex(index)
	})
	l.Remove(indicesList...)
	return l
}

func (l *List[T]) RemovedAny(indices ...int) List[T] {
	answer := l.Copy()
	answer.RemoveAny(indices...)
	return answer
}

// IsEmpty returns true if the list is empty
func (l *List[T]) IsEmpty() bool {
	return len(*l) == 0
}

// ValidIndex checks if the index is within the list bounds
func (l *List[T]) ValidIndex(index int) bool {
	return index >= 0 && index < len(*l)
}

// ValidIndexLoose checks if the index is within the list bounds, allows one index past the end
func (l *List[T]) ValidIndexLoose(index int) bool {
	return index >= 0 && index <= len(*l)
}

// Clear the list
func (l *List[T]) Clear() *List[T] {
	*l = []T{}
	return l
}

// Pop the last item in the list, or the item at i if specified, and return it
// Returns false if the index is out of bounds or if multiple indices are specified
func (l *List[T]) Pop(i ...int) (T, bool) {
	var zero T
	if l.IsEmpty() {
		return zero, false
	}
	if len(i) > 1 {
		return zero, false
	}
	if len(i) == 1 {
		if !l.ValidIndex(i[0]) {
			return zero, false
		}
		item := (*l)[i[0]]
		return item, true
	}
	return (*l)[len(*l)-1], true
}

// Reverse reverses the elements of the list in-place
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

func (l *List[T]) Reversed() List[T] {
	answer := l.Copy()
	answer.Reverse()
	return answer
}

// Shuffle randomizes the order of elements in the list (in-place)
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

func (l *List[T]) Shuffled() List[T] {
	answer := l.Copy()
	answer.Shuffle()
	return answer
}

// Insert inserts values at the specified index
func (l *List[T]) Insert(index int, values ...T) bool {
	if !l.ValidIndexLoose(index) {
		return false
	}
	if index == len(*l) {
		l.Add(values...)
		return true
	}
	prev := (*l)[:index]
	after := (*l)[index:]
	*l = append(prev, values...)
	l.Add(after...)
	return true
}

func (l *List[T]) Inserted(index int, values ...T) List[T] {
	answer := l.Copy()
	answer.Insert(index, values...)
	return answer
}

// Set replaces the element at the specified index
func (l *List[T]) Set(index int, value T) bool {
	if !l.ValidIndex(index) {
		return false
	}
	(*l)[index] = value
	return true
}

func (l *List[T]) Setted(index int, value T) List[T] {
	answer := l.Copy()
	answer.Set(index, value)
	return answer
}

// Swap exchanges elements at two indices
func (l *List[T]) Swap(i, j int) bool {
	if !l.ValidIndex(i) || !l.ValidIndex(j) {
		return false
	}
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	return true
}

func (l *List[T]) Swapped(i, j int) List[T] {
	answer := l.Copy()
	answer.Swap(i, j)
	return answer
}

// Copy returns a new copy of the list
func (l *List[T]) Copy() List[T] {
	c := make(List[T], len(*l))
	copy(c, *l)
	return c
}

// Slice returns a new list containing the elements from start (inclusive) to end (exclusive)
func (l *List[T]) Slice(start int, end int) (List[T], bool) {
	if !l.ValidIndex(start) || !l.ValidIndexLoose(end) {
		return nil, false
	}
	if end < start {
		return nil, false
	}
	return (*l)[start:end], true
}

func (l *List[T]) Filter(predicate func(T) bool) *List[T] {
	*l = l.Filtered(predicate)
	return l
}

func (l *List[T]) Filtered(predicate func(T) bool) List[T] {
	result := NewList[T]()
	for _, item := range *l {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func (l *List[T]) Map(f func(T) T) *List[T] {
	for i, item := range *l {
		(*l)[i] = f(item)
	}
	return l
}

func (l *List[T]) Mapped(f func(T) T) List[T] {
	result := l.Copy()
	result.Map(f)
	return result
}

func (l *List[T]) Any(predicate func(T) bool) bool {
	for _, item := range *l {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (l *List[T]) All(predicate func(T) bool) bool {
	for _, item := range *l {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Find returns the first item for which f returns true, or (T, false) if none match
func (l *List[T]) Find(predicate func(T) bool) (T, bool) {
	for _, item := range *l {
		if predicate(item) {
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

func (l *List[T]) ToSlice() []T {
	slice := make([]T, len(*l))
	copy(slice, *l)
	return slice
}
