package main

import (
	. "github.com/KeylimeVI/kv"
)

func main() {
	list := *NewList[int](1, 2, 3, 4, 5)
	list.Append(6, 7, 8, 9, 10)
}
