package ks

import kl "github.com/KeylimeVI/keylime-go/list"

// ToSlice converts the set to a slice
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(*s))
	for item := range *s {
		slice = append(slice, item)
	}
	return slice
}

// ToList converts the set to a list
func (s *Set[T]) ToList() kl.List[T] {
	return kl.NewList[T](s.ToSlice()...)
}
