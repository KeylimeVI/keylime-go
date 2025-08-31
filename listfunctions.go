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

const NOT_FOUND = -1

// Contains returns true if the list contains the value
func Contains[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
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

func Min[T cmp.Ordered](list []T) T {
	return slices.Min(list)
}

func Max[T cmp.Ordered](list []T) T {
	return slices.Max(list)
}

func Sum[T cmp.Ordered](list []T) T {
	var s T
	for _, item := range list {
		s += item
	}
	return s
}

func Average[T RealNumber](list []T) float64 {
	if len(list) == 0 {
		return 0
	}
	floatList := NewListWithCapacity[float64](len(list))
	for _, item := range list {
		floatList.Append(float64(item))
	}
	return Sum(floatList) / float64(len(list))
}

// Find - always uses linear search
func Find[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return NOT_FOUND
}

func BinarySearch[T cmp.Ordered](slice []T, value T) int {
	if len(slice) == 0 {
		return NOT_FOUND
	}
	if !IsSorted(slice) {
		sortedSlice := NewList[T](slice...)
		Sort(sortedSlice)
		i, ok := slices.BinarySearch(sortedSlice, value)
		if ok {
			return i
		}
		return NOT_FOUND
	}
	i, ok := slices.BinarySearch(slice, value)
	if ok {
		return i
	}
	return NOT_FOUND
}

func Median[T cmp.Ordered](list []T) T {
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
