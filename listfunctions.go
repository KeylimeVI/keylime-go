package kl

import (
	"cmp"
	"slices"
)

type RealNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

const NotFound = -1

// Contains returns true if the list contains the value
func Contains[S ~[]T, T comparable](list S, values ...T) bool {
	for _, value := range values {
		if !singleContains(list, value) {
			return false
		}
	}
	return true
}

func ContainsAny[S ~[]T, T comparable](list S, values ...T) bool {
	for _, value := range values {
		if singleContains(list, value) {
			return true
		}
	}
	return false
}

func Reduce[T any, U any](list []T, initial U, f func(accumulator U, value T) U) U {
	result := initial
	for _, item := range list {
		result = f(result, item)
	}
	return result
}

// Sort sorts any slice-like type in-place using the quicksort algorithm
func Sort[S ~[]T, T cmp.Ordered](list S) {
	if list == nil {
		return
	}
	if IsSorted(list) {
		return
	}
	slices.Sort[S, T](list)
}

// IsSorted returns true if the list is sorted
func IsSorted[S ~[]T, T cmp.Ordered](list S) bool {
	if list == nil || len(list) <= 1 {
		return true
	}
	return slices.IsSorted[S, T](list)
}

func Min[S ~[]T, T cmp.Ordered](list S) T {
	return slices.Min(list)
}

func Max[S ~[]T, T cmp.Ordered](list S) T {
	return slices.Max(list)
}

func Sum[S ~[]T, T cmp.Ordered](list S) T {
	var s T
	for _, item := range list {
		s += item
	}
	return s
}

func Average[S ~[]T, T RealNumber](list S) float64 {
	if len(list) == 0 {
		return 0
	}
	floatList := NewListWithCapacity[float64](len(list))
	for _, item := range list {
		floatList.Add(float64(item))
	}
	return Sum(floatList) / float64(len(list))
}

// Find - always uses linear search
func Find[S ~[]T, T comparable](slice S, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return NotFound
}

func BinarySearch[S ~[]T, T cmp.Ordered](slice S, value T) int {
	if len(slice) == 0 {
		return NotFound
	}
	if !IsSorted(slice) {
		sortedSlice := NewList[T](slice...)
		Sort(sortedSlice)
		i, ok := slices.BinarySearch(sortedSlice, value)
		if ok {
			return i
		}
		return NotFound
	}
	i, ok := slices.BinarySearch(slice, value)
	if ok {
		return i
	}
	return NotFound
}

func Median[S ~[]T, T cmp.Ordered](list S) T {
	if len(list) == 0 {
		var zero T
		return zero
	}

	// Create a copy to work with
	arr := make([]T, len(list))
	copy(arr, list)

	n := len(arr)
	if n%2 == 1 {
		return quickSelect(arr, 0, n-1, n/2)
	}
	// For even length, get both middle elements
	left := quickSelect(arr, 0, n-1, n/2-1)
	quickSelect(arr, 0, n-1, n/2)

	// Since we can't assume arithmetic operations, return the lower one
	// Or if T is numeric, we could average them
	return left
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
	indicesSet := NewSet[int](indices...)
	indicesList := indicesSet.ToList()
	Sort[List[int], int](indicesList)
	indicesList.Reverse()
	return indicesList
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
