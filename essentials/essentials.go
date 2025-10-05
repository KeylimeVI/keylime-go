package essentials

import (
	"github.com/KeylimeVI/keylime-go/list"
	"github.com/KeylimeVI/keylime-go/set"
)

type List[T any] = kl.List[T]

func NewList[T any](items ...T) List[T] {
	return kl.NewList(items...)
}

type Set[T comparable] = ks.Set[T]

func NewSet[T comparable](items ...T) Set[T] {
	return ks.NewSet(items...)
}
