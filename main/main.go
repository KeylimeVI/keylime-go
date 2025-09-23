package main

import (
	"fmt"
	. "github.com/KeylimeVI/kl/list"
)

func main() {
	l := NewList(1, 2, 3)
	s := NewList(1, 2, 3)
	v := NewList(1, 2, 2)
	fmt.Println(l.Equals(s), l.Equals(v), l.Equals(l))
	l.All(func(i int) bool {
		return i == s[i]
	})
}
