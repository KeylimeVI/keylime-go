package main

import (
	"fmt"

	kl "github.com/KeylimeVI/keylime-go/list"
	ks "github.com/KeylimeVI/keylime-go/set"
)

type Container[T any] interface {
	Add(vals ...T) *kl.List[T]
}

func AddZero[T any](c Container[T]) {
	var zero T
	c.Add(zero)
}

func main() {
	l := kl.ListOf(1, 2, 3)
	s := ks.SetOf(4, 5, 6)
	fmt.Printf("%v, %v", l, s)

	AddZero[int](&l)
}

//Testing stuff

type Counter[T kl.RealNumber] interface {
	Increment()
}

type ints interface {
	int | int32
}

type CountInt[T ints] struct {
	i T
}

func (c *CountInt[T]) Increment() {
	c.i++
}
