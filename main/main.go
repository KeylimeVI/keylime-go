package main

import (
	"fmt"
	. "github.com/KeylimeVI/kl"
)

func main() {
	l := NewList[string]("a", "b", "c", "d", "e", "f")
	fmt.Println(l)
	l.Remove(1, 3, 5)
	fmt.Println(l)
	l.RemoveAny(-5, 7, 8, 0)
	fmt.Println(l)
}
