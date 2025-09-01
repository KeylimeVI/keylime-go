package main

import (
	"fmt"
	. "github.com/KeylimeVI/kl"
)

func main() {
	// Create a new map
	m := NewMap[string, int]()

	// Use methods with pointer semantics
	m.Set("apple", 5)
	m.Set("banana", 3)

	// Get values
	if value, exists := m.Get("apple"); exists {
		fmt.Println("Apple count:", value)
	}

	// Modify the map
	m.Delete("banana")
	fmt.Println("Has banana:", m.Has("banana")) // false

	// Merge with another map
	other := NewMapFrom(map[string]int{"orange": 8, "grape": 12})
	m.Merge(other)

	// Iterate
	m.ForEach(func(key string, value int) {
		fmt.Printf("%s: %d\n", key, value)
	})

	// Filter
	filtered := m.Filter(func(key string, value int) bool {
		return value > 4
	})
	fmt.Println("Filtered:", filtered.ToNativeMap())

	// Clear
	m.Clear()
	fmt.Println("Is empty:", m.IsEmpty()) // true
}
