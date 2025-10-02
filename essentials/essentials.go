package essentials

import (
	kl "github.com/KeylimeVI/keylime-go/list"
	kp "github.com/KeylimeVI/keylime-go/pair"
	ks "github.com/KeylimeVI/keylime-go/set"
)

type List[T any] = kl.List[T]

func ListOf[T any](items ...T) List[T] {
	return kl.ListOf(items...)
}

type Set[T comparable] = ks.Set[T]

func SetOf[T comparable](items ...T) Set[T] {
	return ks.SetOf(items...)
}

type Pair[A any, B any] = kp.Pair[A, B]

func PairOf[A any, B any](a A, b B) Pair[A, B] {
	return kp.PairOf(a, b)
}
