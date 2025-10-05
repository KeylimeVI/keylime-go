package main

import (
	. "github.com/KeylimeVI/keylime-go/essentials"
)

func main() {
	l := NewList(1, 2, 3)
	l.Add(2).Filter(func(i int) bool {
		return i > 1
	}).FlatMap(func(i int) []int {
		return []int{i, i + 1, i * i}
	})
}
