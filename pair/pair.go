package kp

import kl "github.com/KeylimeVI/keylime-go/list"

// Pair is a generic type that holds two values
type Pair[A any, B any] struct {
	A A
	B B
}

// NewPair creates a new pair of two values
func NewPair[A any, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{A: a, B: b}
}

// Unwrap returns the values the pair is holding
func (p *Pair[A, B]) Unwrap() (A, B) {
	return p.A, p.B
}

func PairsToMap[K comparable, V any, S ~[]Pair[K, V]](slice S) map[K]V {
	result := make(map[K]V, len(slice))
	for _, pair := range slice {
		k, v := pair.Unwrap()
		result[k] = v
	}
	return result
}

func MapToPairs[K comparable, V any](m map[K]V) kl.List[Pair[K, V]] {
	result := kl.NewList[Pair[K, V]]()
	for k, v := range m {
		p := NewPair(k, v)
		result.Add(p)
	}
	return result
}
