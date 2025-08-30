package kl

import (
	"cmp"
	"math/rand"
)

// Sort sorts any slice-like type in-place using the quicksort algorithm
func Sort[S ~[]T, T cmp.Ordered](list *S) {
	if list == nil || len(*list) <= 1 {
		return
	}
	if IsSorted(list) {
		return
	}
	quicksort(list, 0, len(*list)-1)
}

// quicksort is the recursive implementation of the quicksort algorithm
func quicksort[S ~[]T, T cmp.Ordered](arr *S, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quicksort[S, T](arr, low, pivotIndex-1)
		quicksort[S, T](arr, pivotIndex+1, high)
	}
}

// partition rearranges the array and returns the pivot index
func partition[S ~[]T, T cmp.Ordered](arr *S, low, high int) int {
	pivotIndex := low + rand.Intn(high-low+1)
	pivot := (*arr)[pivotIndex]

	(*arr)[pivotIndex], (*arr)[high] = (*arr)[high], (*arr)[pivotIndex]

	i := low

	for j := low; j < high; j++ {
		if (*arr)[j] <= pivot {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			i++
		}
	}

	(*arr)[i], (*arr)[high] = (*arr)[high], (*arr)[i]
	return i
}

// IsSorted returns true if the list is sorted
func IsSorted[S ~[]T, T cmp.Ordered](list *S) bool {
	if list == nil || len(*list) <= 1 {
		return true
	}
	for i := 1; i < len(*list); i++ {
		if (*list)[i] < (*list)[i-1] {
			return false
		}
	}
	return true
}
