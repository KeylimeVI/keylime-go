package kv

type LinearCollection[T any] interface {
	Append(...T)
	Len() int
	Capacity() int
	Get(int) (T, error)
	Set(int, T) (T, error)
	Remove(int)
	ForEach(func(T) bool)
	Copy()
	ThereExists(func(T) bool) bool
	ForAll(func(T) bool) bool
	Slice(int, int) (List[T], error)
	Reverse()
	Shuffle()
	Clear()
	String() string
	IsEmpty() bool
	ValidIndex(int) bool
	ValidIndexLoose(int) bool
	Insert(int, ...T) error
	Pop() (T, error)
	Dequeue() (T, error)
	Swap(int, int)
	Filter(func(T) bool)
	Find(func(T) bool) (int, bool)
}
