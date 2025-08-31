package kl

type Priority int

const (
	Critical Priority = iota
	High
	Medium
	Low
)

type PriorityQueue[T any] struct {
	queues     *Map[Priority, *Queue[T]] // Use the new Map type
	priorities []Priority                // Ordered list of priorities
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		queues:     NewMap[Priority, *Queue[T]](),
		priorities: []Priority{Critical, High, Medium, Low}, // Highest priority first
	}

	// Initialize queues for each priority level
	for _, priority := range pq.priorities {
		pq.queues.Set(priority, &NewQueue[T]())
	}

	return pq
}

func (pq *PriorityQueue[T]) Enqueue(item T, priority Priority) {
	if queue, exists := pq.queues.Get(priority); exists {
		queue.Enqueue(item)
	} else {
		// If unknown priority, use Normal as default
		if normalQueue, exists := pq.queues.Get(Normal); exists {
			normalQueue.Enqueue(item)
		}
	}
}

func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	// Check queues in priority order (highest first)
	for _, priority := range pq.priorities {
		if queue, exists := pq.queues.Get(priority); exists {
			if item, ok := queue.Dequeue(); ok {
				return item, true
			}
		}
	}

	var zero T
	return zero, false
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	for _, priority := range pq.priorities {
		if queue, exists := pq.queues.Get(priority); exists {
			if item, ok := queue.Peek(); ok {
				return item, true
			}
		}
	}

	var zero T
	return zero, false
}

func (pq *PriorityQueue[T]) Len() int {
	total := 0
	pq.queues.ForEach(func(_ Priority, queue *Queue[T]) {
		total += queue.Len()
	})
	return total
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue[T]) PriorityLen(priority Priority) int {
	if queue, exists := pq.queues.Get(priority); exists {
		return queue.Len()
	}
	return 0
}

func (pq *PriorityQueue[T]) Clear() {
	pq.queues.ForEach(func(_ Priority, queue *Queue[T]) {
		queue.Clear()
	})
}
