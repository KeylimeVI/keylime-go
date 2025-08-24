package main

import (
	"fmt"
)

func main() {
	fmt.Println("Running List tests...")

	testAppend()
	testLen()
	testGet()
	testRemove()
	testIsEmpty()
	testEdgeCases()

	fmt.Println("All tests passed! ✅")
}
