package kp

type Pair[A any, B any] struct {
	A any
	B any
}

func NewPair[A any, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{A: a, B: b}
}

func (p *Pair[A, B]) Unwrap() (A, B) {
	return p.A, p.B
}
