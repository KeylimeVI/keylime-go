package kl

type PriorityLevel = int

const (
	Low PriorityLevel = iota
	Normal
	High
)

type PriorityQueue[T any] struct {
	queue      Queue[T]
	priorities List[PriorityLevel]
}
