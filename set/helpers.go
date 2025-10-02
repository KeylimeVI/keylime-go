package ks

// singleContains Check if set contains item
func (s *Set[T]) singleContains(item T) bool {
	_, exists := (*s)[item]
	return exists
}
