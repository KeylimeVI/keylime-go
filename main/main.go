package main

import (
	"fmt"
	. "github.com/KeylimeVI/kl"
)

func main() {
	l := NewList[int](1, 2, 3, 4, 5)
	l.Shuffle()
	fmt.Println(l)
	Sort(&l)
	fmt.Println(l)
	fmt.Println(IsSorted(&l))
	fmt.Println(Max(l))
	fmt.Println(Sum(l))
	fmt.Println(Min(l))
}
