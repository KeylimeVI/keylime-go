package main

import (
	"fmt"
	. "github.com/KeylimeVI/kl"
)

func main() {
	l := NewList(3, 2, 1, 1, 2, 3)
	RemoveDuplicates(l)
	fmt.Println(l)
}
