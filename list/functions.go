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
