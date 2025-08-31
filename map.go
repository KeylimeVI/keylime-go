package kl

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() *Map[K, V] {
	m := make(Map[K, V])
	return &m
}

func NewMapWithCapacity[K comparable, V any](capacity int) *Map[K, V] {
	m := make(Map[K, V], capacity)
	return &m
}

func NewMapFrom[K comparable, V any](m map[K]V) *Map[K, V] {
	newMap := make(Map[K, V], len(m))
	for k, v := range m {
		newMap[k] = v
	}
	return &newMap
}

// Set or update a key-value pair
func (m *Map[K, V]) Set(key K, value V) {
	(*m)[key] = value
}

// Get a value by key (returns value, exists)
func (m *Map[K, V]) Get(key K) (V, bool) {
	value, exists := (*m)[key]
	return value, exists
}

// Delete a key-value pair
func (m *Map[K, V]) Delete(key K) {
	delete(*m, key)
}

func (m *Map[K, V]) Has(key K) bool {
	_, exists := (*m)[key]
	return exists
}

func (m *Map[K, V]) Len() int {
	return len(*m)
}

func (m *Map[K, V]) IsEmpty() bool {
	return m.Len() == 0
}

func (m *Map[K, V]) Clear() {
	*m = make(Map[K, V])
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	for key := range *m {
		keys = append(keys, key)
	}
	return keys
}

func (m *Map[K, V]) ListKeys() List[K] {
	return NewList[K](m.Keys()...)
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	for _, value := range *m {
		values = append(values, value)
	}
	return values
}

func (m *Map[K, V]) ListValues() List[V] {
	return NewList[V](m.Values()...)
}

func (m *Map[K, V]) ToNativeMap() map[K]V {
	result := make(map[K]V, m.Len())
	for k, v := range *m {
		result[k] = v
	}
	return result
}

// Merge with another map (overwrites existing keys)
func (m *Map[K, V]) Merge(other *Map[K, V]) {
	for key, value := range *other {
		m.Set(key, value)
	}
}

// ForEach iteration
func (m *Map[K, V]) ForEach(fn func(K, V)) {
	for key, value := range *m {
		fn(key, value)
	}
}

// Filter map based on predicate and return new map
func (m *Map[K, V]) Filter(predicate func(K, V) bool) *Map[K, V] {
	result := NewMap[K, V]()
	for key, value := range *m {
		if predicate(key, value) {
			result.Set(key, value)
		}
	}
	return result
}

// Clone the map
func (m *Map[K, V]) Clone() *Map[K, V] {
	return NewMapFrom(m.ToNativeMap())
}
