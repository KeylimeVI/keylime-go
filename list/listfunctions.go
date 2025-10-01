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

const notFound = -1

// Contains returns true if the list contains the value
func Contains[T comparable, S ~[]T](list S, values ...T) bool {
	for _, value := range values {
		if !singleContains(list, value) {
			return false
		}
	}
	return true
}

func ContainsAny[T comparable, S ~[]T](list S, values ...T) bool {
	for _, value := range values {
		if singleContains(list, value) {
			return true
		}
	}
	return false
}

func Reduce[T any, U any, S ~[]T](slice S, initial U, reducer func(accumulator U, value T) U) U {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

func Flatten[T any, S ~[]T, S2 ~[]S](collection S2) S {
	totalLen := 0
	for i := range collection {
		totalLen += len(collection[i])
	}

	result := make(S, 0, totalLen)
	for i := range collection {
		result = append(result, collection[i]...)
	}

	return result
}

func FlatMap[T1 any, T2 any, S1 ~[]T1, S2 ~[]T2](collection S1, iteratee func(item T1) S2) S2 {
	result := make(S2, 0, len(collection))

	for i := range collection {
		result = append(result, iteratee(collection[i])...)
	}
	return result
}

func Map[T1 any, T2 any, S1 ~[]T1, S2 ~[]T2](collection S1, iteratee func(item T1) T2) S2 {
	result := make(S2, len(collection))

	for i := range collection {
		result[i] = iteratee(collection[i])
	}

	return result
}

// Sort sorts any slice-like type in-place using the quicksort algorithm
func Sort[T cmp.Ordered, S ~[]T](list *S) {
	if list == nil {
		return
	}
	if IsSorted(*list) {
		return
	}
	slices.Sort[S, T](*list)
}

// IsSorted returns true if the list is sorted
func IsSorted[T cmp.Ordered, S ~[]T](list S) bool {
	if list == nil || len(list) <= 1 {
		return true
	}
	return slices.IsSorted[S, T](list)
}

func Min[T cmp.Ordered, S ~[]T](list S) T {
	return slices.Min(list)
}

func Max[T cmp.Ordered, S ~[]T](list S) T {
	return slices.Max(list)
}

func Sum[T cmp.Ordered, S ~[]T](list S) T {
	var s T
	for _, item := range list {
		s += item
	}
	return s
}

func Average[T RealNumber, S ~[]T](list S) float64 {
	if len(list) == 0 {
		return 0
	}
	floatList := NewWithCap[float64](len(list))
	for _, item := range list {
		floatList.Add(float64(item))
	}
	return Sum(floatList) / float64(len(list))
}

// Find - always uses linear search
func Find[T comparable, S ~[]T](slice S, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return notFound
}

func BinarySearch[T cmp.Ordered, S ~[]T](slice S, value T) int {
	if len(slice) == 0 {
		return notFound
	}
	if !IsSorted(slice) {
		sortedSlice := NewList[T](slice...)
		Sort(&sortedSlice)
		i, ok := slices.BinarySearch(sortedSlice, value)
		if ok {
			return i
		}
		return notFound
	}
	i, ok := slices.BinarySearch(slice, value)
	if ok {
		return i
	}
	return notFound
}

func Median[T cmp.Ordered, S ~[]T](list S) T {
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
	// For even len, get both middle elements
	left := quickSelect(arr, 0, n-1, n/2-1)
	quickSelect(arr, 0, n-1, n/2)

	// Since we can't assume arithmetic operations, return the lower one
	// Or if T is numeric, we could average them
	return left
}

func RemoveDuplicates[T comparable, S ~[]T](list *S) {
	if len(*list) <= 1 {
		return
	}
	set := make(map[T]struct{}, len(*list))
	for _, item := range *list {
		set[item] = struct{}{}
	}
	result := make(S, 0, len(set))
	for item := range set {
		result = append(result, item)
	}
	*list = result
}
