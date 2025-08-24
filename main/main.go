package main

import (
	"fmt"
	. "kvlist"
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

func testAppend() {
	fmt.Println("\n=== Testing Append ===")

	// Test 1: Append to empty list
	var list List[int]
	list.Append(1, 2, 3)
	if list.Len() != 3 || list[0] != 1 || list[1] != 2 || list[2] != 3 {
		panic("Append to empty list failed")
	}
	fmt.Println("✓ Append to empty list works")

	// Test 2: Append to existing list
	list.Append(4, 5)
	if list.Len() != 5 || list[3] != 4 || list[4] != 5 {
		panic("Append to existing list failed")
	}
	fmt.Println("✓ Append to existing list works")

	// Test 3: Method chaining
	var list2 List[string]
	list2.Append("a").Append("b", "c")
	if list2.Len() != 3 || list2[0] != "a" || list2[1] != "b" || list2[2] != "c" {
		panic("Method chaining failed")
	}
	fmt.Println("✓ Method chaining works")
}

func testLen() {
	fmt.Println("\n=== Testing Len ===")

	// Test 1: Empty list
	var empty List[float64]
	if empty.Len() != 0 {
		panic("Empty list length should be 0")
	}
	fmt.Println("✓ Empty list length is 0")

	// Test 2: Non-empty list
	list := List[int]{1, 2, 3, 4, 5}
	if list.Len() != 5 {
		panic("List length should be 5")
	}
	fmt.Println("✓ Non-empty list length correct")

	// Test 3: After modifications
	list.Append(6)
	if list.Len() != 6 {
		panic("Length after append should be 6")
	}
	fmt.Println("✓ Length updates after append")
}

func testGet() {
	fmt.Println("\n=== Testing Get ===")

	list := List[string]{"apple", "banana", "cherry"}

	// Test 1: Valid indices
	if val, ok := list.Get(0); !ok || val != "apple" {
		panic("Get(0) failed")
	}
	if val, ok := list.Get(1); !ok || val != "banana" {
		panic("Get(1) failed")
	}
	if val, ok := list.Get(2); !ok || val != "cherry" {
		panic("Get(2) failed")
	}
	fmt.Println("✓ Get with valid indices works")

	// Test 2: Invalid indices
	if _, ok := list.Get(-1); ok {
		panic("Get(-1) should fail")
	}
	if _, ok := list.Get(3); ok {
		panic("Get(3) should fail")
	}
	if _, ok := list.Get(100); ok {
		panic("Get(100) should fail")
	}
	fmt.Println("✓ Get with invalid indices fails correctly")

	// Test 3: Empty list
	var empty List[int]
	if _, ok := empty.Get(0); ok {
		panic("Get on empty list should fail")
	}
	fmt.Println("✓ Get on empty list fails correctly")
}

func testRemove() {
	fmt.Println("\n=== Testing Remove ===")

	list := List[int]{1, 2, 3, 4, 5}

	// Test 1: Remove from middle
	err := list.Remove(2) // Remove value 3
	if err != nil || list.Len() != 4 || list[0] != 1 || list[1] != 2 || list[2] != 4 || list[3] != 5 {
		panic("Remove from middle failed")
	}
	fmt.Println("✓ Remove from middle works")

	// Test 2: Remove from beginning
	err = list.Remove(0) // Remove value 1
	if err != nil || list.Len() != 3 || list[0] != 2 || list[1] != 4 || list[2] != 5 {
		panic("Remove from beginning failed")
	}
	fmt.Println("✓ Remove from beginning works")

	// Test 3: Remove from end
	err = list.Remove(2) // Remove value 5
	if err != nil || list.Len() != 2 || list[0] != 2 || list[1] != 4 {
		panic("Remove from end failed")
	}
	fmt.Println("✓ Remove from end works")

	// Test 4: Invalid indices
	err = list.Remove(-1)
	if err == nil {
		panic("Remove(-1) should fail")
	}
	err = list.Remove(2)
	if err == nil {
		panic("Remove(2) on length 2 list should fail")
	}
	fmt.Println("✓ Remove with invalid indices fails correctly")

	// Test 5: Empty list
	var empty List[string]
	err = empty.Remove(0)
	if err == nil || err.Error() != "list is empty" {
		panic("Remove from empty list should fail")
	}
	fmt.Println("✓ Remove from empty list fails correctly")
}

func testIsEmpty() {
	fmt.Println("\n=== Testing IsEmpty ===")

	// Test 1: Empty list
	var empty List[float64]
	if !empty.IsEmpty() {
		panic("Empty list should report IsEmpty() = true")
	}
	fmt.Println("✓ Empty list reports IsEmpty() = true")

	// Test 2: Non-empty list
	list := List[int]{1}
	if list.IsEmpty() {
		panic("Non-empty list should report IsEmpty() = false")
	}
	fmt.Println("✓ Non-empty list reports IsEmpty() = false")

	// Test 3: After modifications
	list.Append(2)
	if list.IsEmpty() {
		panic("List with items should report IsEmpty() = false")
	}

	list.Remove(0)
	list.Remove(0)
	if !list.IsEmpty() {
		panic("List after removing all items should report IsEmpty() = true")
	}
	fmt.Println("✓ IsEmpty() updates correctly after modifications")
}

func testEdgeCases() {
	fmt.Println("\n=== Testing Edge Cases ===")

	// Test 1: Empty list operations
	var empty List[complex128]
	if empty.Len() != 0 || !empty.IsEmpty() {
		panic("Empty list state incorrect")
	}
	if _, ok := empty.Get(0); ok {
		panic("Empty list Get should fail")
	}
	if err := empty.Remove(0); err == nil {
		panic("Empty list Remove should fail")
	}
	fmt.Println("✓ Empty list operations work correctly")

	// Test 2: Single element list
	single := List[string]{"only"}
	if single.Len() != 1 || single.IsEmpty() {
		panic("Single element list state incorrect")
	}
	if val, ok := single.Get(0); !ok || val != "only" {
		panic("Single element Get failed")
	}
	if err := single.Remove(0); err != nil || !single.IsEmpty() {
		panic("Single element Remove failed")
	}
	fmt.Println("✓ Single element list operations work correctly")

	// Test 3: Type compatibility
	intList := List[int]{1, 2, 3}
	stringList := List[string]{"a", "b", "c"}

	if intList.Len() != 3 || stringList.Len() != 3 {
		panic("Generic types not working correctly")
	}
	fmt.Println("✓ Generic types work correctly")

	fmt.Println("✓ All edge cases passed")
}
