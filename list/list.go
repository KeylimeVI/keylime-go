package list

import (
	"fmt"
	"math/rand"
)

// List is a generic type alias of []T with useful methods and functions
type List[T any] []T

// NewList creates a new List with the specified values
func NewList[T any](vals ...T) List[T] {
	list := List[T](vals)
	return list
}

// NewListWithCapacity creates a new List with the specified capacity and values
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

	indicesList := NewList[int](indices...)

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

func (l *List[T]) RemoveAny(indices ...int) *List[T] {
	indicesList := formatIndicesReversed(indices)
	indicesList = indicesList.Filtered(func(index int) bool {
		return l.ValidIndex(index)
	})
	_ = l.Remove(indicesList...)
	return l
}

func (l *List[T]) Removed(indices ...int) List[T] {
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
// Returns an error if the index is out of bounds or if multiple indices are specified
func (l *List[T]) Pop(i ...int) (T, error) {
	var zero T
	if len(i) > 1 {
		return zero, TooManyArgumentsError
	}
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
func (l *List[T]) Insert(index int, values ...T) error {
	if !l.ValidIndexLoose(index) {
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

func (l *List[T]) Inserted(index int, values ...T) List[T] {
	answer := l.Copy()
	_ = answer.Insert(index, values...)
	return answer
}

// Set replaces the element at the specified index
func (l *List[T]) Set(index int, value T) error {
	if !l.ValidIndex(index) {
		return NewIndexError(index, l.Len())
	}
	(*l)[index] = value
	return nil
}

func (l *List[T]) Setted(index int, value T) List[T] {
	answer := l.Copy()
	_ = answer.Set(index, value)
	return answer
}

// Swap exchanges elements at two indices
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

func (l *List[T]) Swapped(i, j int) List[T] {
	answer := l.Copy()
	_ = answer.Swap(i, j)
	return answer
}

// Copy returns a new copy of the list
func (l *List[T]) Copy() List[T] {
	c := make(List[T], len(*l))
	copy(c, *l)
	return c
}

func (l *List[T]) Slice(start int, end int) List[T] {
	if start < 0 {
		start = 0
	}
	if end < l.Len() {
		end = l.Len()
	}
	return (*l)[start:end]
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

func (l *List[T]) Partition(size int) List[List[T]] {
	if size <= 0 {
		// Return the whole list as one chunk
		return List[List[T]]{l.Copy()}
	}

	var parts List[List[T]]
	for i := 0; i < len(*l); i += size {
		end := i + size
		if end > len(*l) {
			end = len(*l)
		}

		// Create a new part list
		part := List[T]{}
		part = append(part, (*l)[i:end]...)
		parts = append(parts, part)
	}
	return parts
}

// ToSlice converts the list to a native slice
func (l *List[T]) ToSlice() []T {
	slice := make([]T, len(*l))
	copy(slice, *l)
	return slice
}
