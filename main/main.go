package main

import (
	"fmt"
	. "kv"
)

func main() {
	fmt.Println("\n=== Testing Set ===")

	testNewSet()
	testSetContains()
	testSetAppend()
	testSetExtend()
	testSetInsert()
	testSetSet()
	testSetUniqueness()
	testSetEdgeCases()

	fmt.Println("All Set tests passed! ✅")
}

func testNewSet() {
	fmt.Println("Testing NewSet...")

	// Test empty set
	emptySet := NewSet[int]()
	if emptySet.Len() != 0 {
		panic("Empty set should have length 0")
	}

	// Test set with values
	set := NewSet(1, 2, 3, 2, 1) // Duplicates should be removed
	if set.Len() != 3 {
		panic("NewSet should remove duplicates")
	}
	fmt.Println("✓ NewSet works correctly")
}

func testSetContains() {
	fmt.Println("Testing Contains...")

	set := NewSet("apple", "banana", "cherry")

	// Test existing values
	if !set.Contains("apple") {
		panic("Should contain 'apple'")
	}
	if !set.Contains("banana") {
		panic("Should contain 'banana'")
	}

	// Test non-existing values
	if set.Contains("orange") {
		panic("Should not contain 'orange'")
	}
	if set.Contains("") {
		panic("Should not contain empty string")
	}
	fmt.Println("✓ Contains works correctly")
}

func testSetAppend() {
	fmt.Println("Testing Append...")

	set := NewSet[int]()

	// Append unique values
	set.Append(1, 2, 3)
	if set.Len() != 3 {
		panic("Append should add unique values")
	}
	set.Append(2, 3, 4) // Only 4 should be added
	if set.Len() != 4 || !set.Contains(4) {
		panic("Append should not add duplicates")
	}

	// Test method chaining
	set.Append(5).Append(6)
	if set.Len() != 6 {
		panic("Method chaining should work")
	}
	fmt.Println("✓ Append works correctly")
}

func testSetExtend() {
	fmt.Println("Testing Extend...")

	set := NewSet(1, 2, 3)
	newValues := []int{3, 4, 5} // 3 is duplicate

	// Extend with slice
	set.Extend(newValues)
	if set.Len() != 5 || !set.Contains(4) || !set.Contains(5) {
		panic("Extend should add unique values from slice")
	}

	// Test that duplicates weren't added
	count := 0
	for _, v := range set.List {
		if v == 3 {
			count++
		}
	}
	if count != 1 {
		panic("Extend should not add duplicates")
	}
	fmt.Println("✓ Extend works correctly")
}

func testSetInsert() {
	fmt.Println("Testing Insert...")

	set := NewSet(1, 3, 5) // [1, 3, 5]

	// Insert unique value at valid index
	err := set.Insert(1, 2) // Should become [1, 2, 3, 5]
	if err != nil {
		panic("Insert should work at valid index")
	}
	if set.Len() != 4 || !set.Contains(2) {
		panic("Insert should add unique value")
	}

	// Check position
	if val, ok := set.Get(1); !ok || val != 2 {
		panic("Insert should place value at correct position")
	}

	// Insert duplicate value
	err = set.Insert(0, 1) // 1 already exists
	if err != nil {
		panic("Insert should not fail for duplicates, just skip them")
	}
	if set.Len() != 4 { // Length should not change
		panic("Insert should not add duplicates")
	}

	// Insert at invalid index
	err = set.Insert(10, 6)
	if err == nil {
		panic("Insert should fail for out-of-bounds index")
	}

	// Insert multiple values
	set = NewSet(1, 4)        // [1, 4]
	err = set.Insert(1, 2, 3) // Should become [1, 2, 3, 4]
	if err != nil || set.Len() != 4 {
		panic("Insert with multiple values should work")
	}
	fmt.Println("✓ Insert works correctly")
}

func testSetSet() {
	fmt.Println("Testing Set...")

	set := NewSet(1, 2, 3) // [1, 2, 3]

	// Set with unique value
	err := set.Set(1, 4) // Should set index 1 to 4: [1, 4, 3]
	if err != nil {
		panic("Set should work with unique value")
	}
	if !set.Contains(4) || set.Contains(2) {
		panic("Set should replace value and maintain uniqueness")
	}

	// Set with duplicate value
	err = set.Set(0, 3) // Try to set index 0 to 3 (which already exists at index 2)
	if err != nil {
		panic("Set should not fail for duplicates, just skip")
	}
	// Value should not change
	if val, ok := set.Get(0); !ok || val != 1 {
		panic("Set should not change value if duplicate")
	}

	// Set at invalid index
	err = set.Set(10, 5)
	if err == nil {
		panic("Set should fail for out-of-bounds index")
	}
	fmt.Println("✓ Set works correctly")
}

func testSetUniqueness() {
	fmt.Println("Testing Uniqueness...")

	set := NewSet(1, 2, 3, 2, 1, 3, 2, 1)

	// Check that only unique values exist
	if set.Len() != 3 {
		panic("Set should maintain uniqueness")
	}

	// Count occurrences of each value
	counts := make(map[int]int)
	for _, v := range set.List {
		counts[v]++
	}

	for _, count := range counts {
		if count != 1 {
			panic("Each value should appear exactly once")
		}
	}
	fmt.Println("✓ Uniqueness is maintained")
}

func testSetEdgeCases() {
	fmt.Println("Testing Edge Cases...")

	// Test empty set operations
	emptySet := NewSet[int]()
	if emptySet.Contains(1) {
		panic("Empty set should not contain anything")
	}

	// Test with different types
	stringSet := NewSet("a", "b", "a", "c")
	if stringSet.Len() != 3 {
		panic("String set should maintain uniqueness")
	}

	// Test with struct values
	type Point struct{ X, Y int }
	pointSet := NewSet(Point{1, 2}, Point{1, 2}, Point{3, 4})
	if pointSet.Len() != 2 { // Duplicate Point{1, 2} should be removed
		panic("Struct set should maintain uniqueness")
	}

	// Test method inheritance from List
	set := NewSet(1, 2, 3)
	if set.Len() != 3 {
		panic("Should inherit Len method from List")
	}

	// Test that we can still use List methods
	set.List.Append(4) // Bypass Set's uniqueness check
	if set.Len() != 4 {
		panic("Should be able to use embedded List methods")
	}
	// But this breaks uniqueness!
	set.List.Append(4) // Add duplicate directly
	if set.Len() != 5 {
		panic("Direct List access should bypass uniqueness")
	}
	fmt.Println("✓ Edge cases handled correctly")
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
