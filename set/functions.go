package ks

// Reduce applies a reducer function over all elements of the set, starting from initial, and returns the accumulated result.
// Note: iteration order over a set is undefined.
func Reduce[T comparable, U any](set Set[T], initial U, reducer func(accumulator U, value T) U) U {
	result := initial
	for item := range set {
		result = reducer(result, item)
	}
	return result
}

// Map transforms a set of T1 into a set of T2 using the provided iteratee.
// T2 must be comparable to be used as a set element.
func Map[T comparable, U comparable](set Set[T], iteratee func(item T) U) Set[U] {
	result := SetOf[U]()
	for item := range set {
		result.Add(iteratee(item))
	}
	return result
}

// FlatMap maps each element of the set to a slice of T2 and flattens the results into a single set of T2 (i.e., union of all mapped values).
func FlatMap[T comparable, U comparable, S ~[]U](set Set[T], iteratee func(item T) S) Set[U] {
	result := SetOf[U]()
	for item := range set {
		mapped := iteratee(item)
		if len(mapped) > 0 {
			result.Add(mapped...)
		}
	}
	return result
}

// Flatten unions a collection (slice-like) of sets into a single set.
// This mirrors list.Flatten (which flattens a slice of slices) but for sets we flatten a slice of sets because a set of sets is not representable in Go.
func Flatten[T comparable, C ~[]Set[T]](collection C) Set[T] {
	result := SetOf[T]()
	for i := range collection {
		for item := range collection[i] {
			result.Add(item)
		}
	}
	return result
}
