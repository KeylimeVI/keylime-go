package main

import (
	"fmt"

	kl "github.com/KeylimeVI/keylime-go/list"
	ks "github.com/KeylimeVI/keylime-go/set"
)

func main() {
	l := kl.ListOf(1, 2, 3)
	s := ks.SetOf(4, 5, 6)
	fmt.Printf("%v, %v", l, s)
}
