package main

import (
	"fmt"
	. "github.com/KeylimeVI/kl"
)

func main() {
	list := NewList[int](1, 2, 3, 4, 5)
	list.Append(6, 7, 8, 9, 10)
	list.Reverse()
	fmt.Println(list)
	list.Shuffle()
	fmt.Println(list)
	Sort[List[int], int](&list)
	fmt.Println(list)
}
