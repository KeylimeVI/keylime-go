package essentials

import (
	kl "github.com/KeylimeVI/keylime-go/list"
	kp "github.com/KeylimeVI/keylime-go/pair"
	ks "github.com/KeylimeVI/keylime-go/set"
)

type List[T any] = kl.List[T]

func NewList[T any](items ...T) List[T] {
	return kl.NewList(items...)
}

type Set[T comparable] = ks.Set[T]

func NewSet[T comparable](items ...T) Set[T] {
	return ks.NewSet(items...)
}

type Pair[A any, B any] = kp.Pair[A, B]

func NewPair[A any, B any](a A, b B) Pair[A, B] {
	return kp.NewPair(a, b)
}
