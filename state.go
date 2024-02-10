package lib

import "github.com/beavorp/gofui/element"

type State[T any] struct {
	e element.Element
	v T
}

func (s *State[T]) Get() T {
	return s.v
}

func (s *State[T]) Set(v T) {
	s.v = v
	s.e.Render()
}
