package gofui

import "github.com/beavorp/gofui/element"

type Store[T any] struct {
	v         T
	consumers map[element.Element]bool
}

func NewStore[T any](v T) *Store[T] {
	return &Store[T]{
		v:         v,
		consumers: make(map[element.Element]bool, 5), // 5 is an arbitrary number
	}
}

func (s *Store[T]) Get() T {
	return s.v
}

func (s *Store[T]) Set(v T) {
	s.v = v
	for e, _ := range s.consumers {
		e.Render()
	}
}

func (s *Store[T]) subscribe(c element.Element) {
	s.consumers[c] = true
}
