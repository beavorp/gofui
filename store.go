package gofui

import "github.com/beavorp/gofui/element"

type Store[T any] struct {
	v         T
	consumers []element.Element
}

func NewStore[T any](v T) *Store[T] {
	return &Store[T]{
		v:         v,
		consumers: make([]element.Element, 0, 5), // 5 is an arbitrary number
	}
}

func (s *Store[T]) Get() T {
	return s.v
}

func (s *Store[T]) Set(v T) {
	s.v = v
	for _, c := range s.consumers {
		c.Render()
	}
}

func (s *Store[T]) subscribe(c element.Element) {
	s.consumers = append(s.consumers, c)
}
