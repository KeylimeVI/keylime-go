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

// Reduce applies a reducer function over the slice-like list, starting at initial, and returns the accumulated result.
func Reduce[T any, U any, S ~[]T](list S, initial U, reducer func(accumulator U, value T) U) U {
	result := initial
	for _, item := range list {
		result = reducer(result, item)
	}
	return result
}

// Flatten concatenates a slice of slices into a single slice.
func Flatten[T any, S ~[]T, S2 ~[]S](list S2) S {
	totalLen := 0
	for i := range list {
		totalLen += len(list[i])
	}

	result := make(S, 0, totalLen)
	for i := range list {
		result = append(result, list[i]...)
	}

	return result
}

// FlatMap maps each element to a slice and concatenates the results.
func FlatMap[T1 any, T2 any, S1 ~[]T1, S2 ~[]T2](list S1, iteratee func(item T1) S2) S2 {
	result := make(S2, 0, len(list))

	for i := range list {
		result = append(result, iteratee(list[i])...)
	}
	return result
}

// Map transforms a slice of T1 into a slice of T2 using iteratee.
func Map[T1 any, T2 any, S1 ~[]T1, S2 ~[]T2](list S1, iteratee func(item T1) T2) S2 {
	result := make(S2, len(list))

	for i := range list {
		result[i] = iteratee(list[i])
	}

	return result
}

// Sort sorts a slice in place
func Sort[T cmp.Ordered, S ~[]T](list S) {
	if list == nil {
		return
	}
	if IsSorted(list) {
		return
	}
	slices.Sort[S, T](list)
}

// IsSorted returns true if the list is sorted
func IsSorted[T cmp.Ordered, S ~[]T](list S) bool {
	if list == nil || len(list) <= 1 {
		return true
	}
	return slices.IsSorted[S, T](list)
}

// Min returns the smallest element of the slice.
func Min[T cmp.Ordered, S ~[]T](list S) T {
	return slices.Min(list)
}

// Max returns the largest element of the slice.
func Max[T cmp.Ordered, S ~[]T](list S) T {
	return slices.Max(list)
}

// Sum adds up all elements in the slice and returns the total.
func Sum[T cmp.Ordered, S ~[]T](list S) T {
	var s T
	for _, item := range list {
		s += item
	}
	return s
}

// RemoveDuplicates removes duplicate values from the slice in place, preserving the first occurrence.
func RemoveDuplicates[T comparable, S ~[]T](list *S) {
	if len(*list) <= 1 {
		return
	}
	result := ListOf[T]()
	for _, v := range *list {
		if !result.singleContains(v) {
			result.Add(v)
		}
	}
	*list = result.ToSlice()
	return
}
