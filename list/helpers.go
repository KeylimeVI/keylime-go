package list

import (
	"cmp"
	"errors"
	"fmt"
	"reflect"
)

// Exported sentinel errors (preferred names)
var (
	EmptyListError        = errors.New("list is empty")
	IndexOutOfBoundsError = errors.New("index out of bounds")
)

// IndexError provides structured context for invalid index operations.
// It unwraps to the IndexOutOfBoundsError sentinel so callers can use errors.Is.
type IndexError struct {
	index  int
	length int
}

func (e IndexError) Error() string {
	if e.index < 0 {
		return fmt.Sprintf("index out of bounds: index = %d", e.index)
	}
	if e.length == 0 {
		return fmt.Sprintf("list is empty: length of list = %d", e.length)
	}
	return fmt.Sprintf("index out of bounds: index = %d, length of list = %d", e.index, e.length)

}

// Unwrap enables errors.Is(err, IndexOutOfBoundsError).
func (e IndexError) Unwrap() error { return IndexOutOfBoundsError }

// Accessors for structured data without exporting fields.
func (e IndexError) Index() int  { return e.index }
func (e IndexError) Length() int { return e.length }

// NewIndexError constructs a typed index error with context.
func NewIndexError(index int, length int) error {
	return IndexError{index: index, length: length}
}

func quickSelect[T cmp.Ordered](arr []T, left, right, k int) T {
	if left == right {
		return arr[left]
	}

	pivotIndex := partition(arr, left, right)

	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelect[T](arr, left, pivotIndex-1, k)
	} else {
		return quickSelect[T](arr, pivotIndex+1, right, k)
	}
}

func partition[T cmp.Ordered](arr []T, left, right int) int {
	pivot := arr[right] // Use the rightmost element as pivot
	i := left

	// Partition the array around the pivot
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Move pivot to its final position
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func formatIndicesReversed(indices List[int]) List[int] {
	if indices.Len() <= 1 {
		return indices
	}
	RemoveDuplicates(&indices)
	Sort[int, List[int]](&indices)
	indices.Reverse()
	return indices
}

func indicesAreFormattedReversed(indices List[int]) bool {
	if indices.Len() <= 1 {
		return true
	}
	for index, item := range indices {
		if index == indices.Len()-1 {
			return true
		}
		if item <= indices[index+1] {
			return false
		}
	}
	return true
}

func singleContains[S ~[]T, T comparable](list S, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func (l *List[T]) singleSet(index int, value T) bool {
	if !l.ValidIndex(index) {
		return false
	}
	(*l)[index] = value
	return true
}

func copyList[S ~[]T, T any](list S) S {
	result := make(S, len(list))
	copy(result, list)
	return result
}

func isComparable[T any](value T) bool {
	valueAny := any(value)
	t := reflect.TypeOf(valueAny)
	return t.Comparable()
}
