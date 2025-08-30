package kl

import (
	"cmp"
	"slices"
)

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
func Sort[S ~[]T, T cmp.Ordered](list *S) {
	if list == nil || len(*list) <= 1 {
		return
	}
	if IsSorted(list) {
		return
	}
	slices.Sort[S, T](*list)
}

// IsSorted returns true if the list is sorted
func IsSorted[S ~[]T, T cmp.Ordered](list *S) bool {
	if list == nil || len(*list) <= 1 {
		return true
	}
	return slices.IsSorted[S, T](*list)
}
