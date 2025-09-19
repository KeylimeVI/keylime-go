package kl

type Queue[T any] struct {
	list List[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{list: NewList[T]()}
}

func (q *Queue[T]) Enqueue(items ...T) *Queue[T] {
	q.list.Add(items...)
	return q
}

func (q *Queue[T]) Dequeue() (T, bool) {
	item, ok := q.list.Pop(0)
	return item, ok
}

func (q *Queue[T]) Peek() (T, bool) {
	return q.list.Get(0)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Clear() {
	q.list.Clear()
}

func (q *Queue[T]) Copy() *Queue[T] {
	return &Queue[T]{list: q.list.Copy()}
}

func (q *Queue[T]) ToList() List[T] {
	return q.list.Copy()
}

func (q *Queue[T]) ToSlice() []T {
	return q.list.ToSlice()
}
