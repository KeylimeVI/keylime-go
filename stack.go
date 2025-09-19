package kl

type Stack[T any] struct {
	list List[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{list: NewList[T]()}
}

func (s *Stack[T]) Push(items ...T) {
	s.list.Add(items...)
}

func (s *Stack[T]) Pop() (T, bool) {
	item, successful := s.list.Pop()
	if !successful {
		var zero T
		return zero, false
	}
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	return s.list.Get(s.list.Len() - 1)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

func (s *Stack[T]) Clear() {
	s.list.Clear()
}

func (s *Stack[T]) Copy() *Stack[T] {
	return &Stack[T]{list: s.list.Copy()}
}

func (s *Stack[T]) ToList() List[T] {
	return s.list.Copy()
}

func (s *Stack[T]) ToSlice() []T {
	return s.list.ToSlice()
}
