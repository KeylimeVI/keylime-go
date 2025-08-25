package kv

import "math/rand"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Sort sorts the list in-place using the quicksort algorithm
func Sort[T Number](list *List[T]) {
	if list.IsEmpty() || list.Len() <= 1 {
		return
	}
	if IsSorted(*list) {
		return
	}
	quicksort(*list, 0, len(*list)-1)
}

// quicksort is the recursive implementation of the quicksort algorithm
func quicksort[T Number](arr List[T], low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		quicksort[T](arr, low, pivotIndex-1)
		quicksort[T](arr, pivotIndex+1, high)
	}
}

// partition rearranges the array and returns the pivot index
func partition[T Number](arr List[T], low, high int) int {
	pivotIndex := low + rand.Intn(high-low+1)
	pivot := arr[pivotIndex]

	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	i := low

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}

// IsSorted returns true if the list is sorted
func IsSorted[T Number](list List[T]) bool {
	for i := 1; i < len(list); i++ {
		if list[i] < list[i-1] {
			return false
		}
	}
	return true
}
