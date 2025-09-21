package essentials

import (
	"github.com/KeylimeVI/kl/list"
	hashmap "github.com/KeylimeVI/kl/map"
	"github.com/KeylimeVI/kl/set"
)

// Re-export List type and constructors
type List[T any] = list.List[T]

var NewList = list.NewList
var NewListWithCapacity = list.NewListWithCapacity

// Re-export Map type and constructors
type Map[K comparable, V any] = hashmap.Map[K, V]

var NewMap = hashmap.NewMap
var NewMapWithCapacity = hashmap.NewMapWithCapacity
var NewMapFrom = hashmap.NewMapFrom

// Re-export Set type and constructors
type Set[T comparable] = set.Set[T]

var NewSet = set.NewSet
var NewSetWithCapacity = set.NewSetWithCapacity

// Re-export list functions
var Contains = list.Contains
var ContainsAny = list.ContainsAny
var Reduce = list.Reduce
var Flatten = list.Flatten
var Sort = list.Sort
var Sorted = list.Sorted
var IsSorted = list.IsSorted
var Min = list.Min
var Max = list.Max
var Sum = list.Sum
var Average = list.Average
var Find = list.Find
var BinarySearch = list.BinarySearch
var Median = list.Median
var RemoveDuplicates = list.RemoveDuplicates
var DuplicatesRemoved = list.DuplicatesRemoved

// Add other essential functions as needed
