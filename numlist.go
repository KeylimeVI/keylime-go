package kv

import (
	"golang.org/x/exp/constraints"
	"math/rand"
)

// Sort sorts the list in-place using the quicksort algorithm
func Sort[T constraints.Ordered](list *[]T) {
	if len(*list) <= 1 {
		return
	}
	if IsSorted(list) {
		return
	}
	quicksort(list, 0, len(*list)-1)
}

// quicksort is the recursive implementation of the quicksort algorithm
func quicksort[T constraints.Ordered](arr *[]T, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		quicksort[T](arr, low, pivotIndex-1)
		quicksort[T](arr, pivotIndex+1, high)
	}
}

// partition rearranges the array and returns the pivot index
func partition[T constraints.Ordered](arr *[]T, low, high int) int {
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
func IsSorted[T constraints.Ordered](list *[]T) bool {
	for i := 1; i < len(*list); i++ {
		if (*list)[i] < (*list)[i-1] {
			return false
		}
	}
	return true
}
