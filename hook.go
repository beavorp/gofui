package lib

import "github.com/beavorp/gofui/element"

func UseState[T any](e element.Element, v T) *State[T] {
	return &State[T]{
		e: e,
		v: v,
	}
}

func UseStore[T any](c element.Element, s *Store[T]) *Store[T] {
	s.subscribe(c)
	return s
}
