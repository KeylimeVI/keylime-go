package essentials

import (
	"github.com/KeylimeVI/keylime-go/list"
	"github.com/KeylimeVI/keylime-go/set"
)

// List is a generic type alias of []T with useful methods and functions
type List[T any] = kl.List[T]

// NewList creates a new List with the specified items
func NewList[T any](items ...T) List[T] {
	return kl.NewList(items...)
}

// Set is a generic set of comparable elements implemented as map[T]struct{}.
type Set[T comparable] = ks.Set[T]

// NewSet creates a new Set with the specified items
func NewSet[T comparable](items ...T) Set[T] {
	return ks.NewSet(items...)
}
