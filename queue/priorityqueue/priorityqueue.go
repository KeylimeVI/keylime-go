package priorityqueue

import (
	"github.com/KeylimeVI/kl/list"
	"github.com/KeylimeVI/kl/map"
	"github.com/KeylimeVI/kl/queue"
	"github.com/KeylimeVI/kl/set"
)

type Priority = int

const (
	Critical Priority = iota
	High
	Medium
	Low
)

type PriorityQueue[T any] struct {
	queues     *hashmap.Map[Priority, *queue.Queue[T]]
	priorities set.Set[Priority]
}

func New[T any]() PriorityQueue[T] {
	pq := PriorityQueue[T]{
		queues:     &hashmap.Map[Priority, *queue.Queue[T]]{},
		priorities: set.NewSet(Critical, High, Medium, Low), // order matters
	}
	// init queues for each priority
	for p := range pq.priorities {
		q := queue.NewQueue[T]()
		pq.queues.Set(p, &q)
	}
	return pq
}

func NewCustom[T any](priorities set.Set[Priority]) PriorityQueue[T] {
	pq := PriorityQueue[T]{
		queues:     &hashmap.Map[Priority, *queue.Queue[T]]{},
		priorities: priorities,
	}
	// init queues for each priority
	for p := range pq.priorities {
		q := queue.NewQueue[T]()
		pq.queues.Set(p, &q)
	}
	return pq
}

func (pq *PriorityQueue[T]) Enqueue(p Priority, items ...T) {
	q, _ := pq.queues.Get(p)
	q.Enqueue(items...)
}

func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	for p := range pq.priorities {
		q, _ := pq.queues.Get(p)
		if !q.IsEmpty() {
			return q.Dequeue()
		}
	}
	var zero T
	return zero, false
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	for p := range pq.priorities {
		q, _ := pq.queues.Get(p)
		if !q.IsEmpty() {
			return q.Peek()
		}
	}
	var zero T
	return zero, false
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	for p := range pq.priorities {
		q, _ := pq.queues.Get(p)
		if !q.IsEmpty() {
			return false
		}
	}
	return true
}

func (pq *PriorityQueue[T]) Len() int {
	count := 0
	for p := range pq.priorities {
		q, _ := pq.queues.Get(p)
		count += q.Len()
	}
	return count
}

// Clear removes all elements from all queues.
func (pq *PriorityQueue[T]) Clear() {
	for p := range pq.priorities {
		q, _ := pq.queues.Get(p)
		q.Clear()
	}
}

// Copy returns a deep copy of the PriorityQueue.
func (pq *PriorityQueue[T]) Copy() *PriorityQueue[T] {
	newPQ := &PriorityQueue[T]{
		queues:     &hashmap.Map[Priority, *queue.Queue[T]]{},
		priorities: pq.priorities, // copy priorities list
	}
	for p := range pq.priorities {
		q, _ := pq.queues.Get(p)
		newPQ.queues.Set(p, q.Copy())
	}
	return newPQ
}

// ToList returns all elements from all priorities as a List, ordered by priority.
func (pq *PriorityQueue[T]) ToList() list.List[T] {
	result := list.NewList[T]()
	for p := range pq.priorities {
		q, _ := pq.queues.Get(p)
		result.Add(q.ToList()...)
	}
	return result
}

// ToSlice returns all elements from all priorities as a slice, ordered by priority.
func (pq *PriorityQueue[T]) ToSlice() []T {
	result := list.NewList[T]()
	for p := range pq.priorities {
		q, _ := pq.queues.Get(p)
		result.Add(q.ToList()...)
	}
	return result.ToSlice()
}
