package collections

import "golang.org/x/exp/maps"

// Set emulates a mathematical set.
// It is backed by a map but intent is clearer when using named functions
// instead of map operations.
// Presence in the map indicates existence in the set.
// Not thread-safe
type Set[T comparable] struct {
	set map[T]struct{}
}

// NewSet assumes a set of strings
func NewSet(elements ...string) *Set[string] {
	return NewSetOf[string](elements...)
}

// NewSetOf permits a parameterized type
func NewSetOf[T comparable](elements ...T) *Set[T] {
	set := &Set[T]{
		set: make(map[T]struct{}),
	}

	set.Add(elements...)

	return set
}

// IsEmpty returns true if and only if no elements are in the set
func (s *Set[T]) IsEmpty() bool {
	return len(s.set) == 0
}

// Add an element to the set
func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		s.set[element] = struct{}{}
	}
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {
	return len(s.set)
}

// Contains indicates whether all elements are contained in the set
func (s *Set[T]) Contains(elements ...T) bool {
	for _, element := range elements {
		if _, contains := s.set[element]; !contains {
			return false
		}
	}

	return true
}

// Remove all elements from the set
func (s *Set[T]) Remove(elements ...T) {
	for _, element := range elements {
		delete(s.set, element)
	}
}

// AddAll performs a set union by modifying s
func (s *Set[T]) AddAll(other *Set[T]) {
	maps.Copy(s.set, other.set)
}

// KeepOnly performs a set intersection by modifying s
func (s *Set[T]) KeepOnly(other *Set[T]) {
	maps.DeleteFunc(s.set, func(k T, _ struct{}) bool {
		return !other.Contains(k)
	})
}

// Elements returns all the elements of the set as an array.
// Order within the array is not guaranteed.
func (s *Set[T]) Elements() []T {
	return maps.Keys(s.set)
}
