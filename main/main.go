package main

import (
	"fmt"
	. "github.com/KeylimeVI/kl/list"
	"github.com/KeylimeVI/kl/sequence"
)

func main() {
	l := NewList(3, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	RemoveDuplicates(&l)
	fmt.Println(l)
	Sort(&l)
	fmt.Println(l)
	for i := range sequence.Sequence(6, 0, -2) {
		fmt.Println(i)
	}
}
