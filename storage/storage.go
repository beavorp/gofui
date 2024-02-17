package storage

import (
	"syscall/js"
)

type Storage struct {
	ref js.Value
}

var (
	Local   = &Storage{ref: js.Global().Get("window").Get("localStorage")}
	Session = &Storage{ref: js.Global().Get("window").Get("sessionStorage")}
)

func (s *Storage) GetItem(key string) js.Value {
	return s.ref.Call("getItem", key)
}

func (s *Storage) SetItem(key, value any) {
	s.ref.Call("setItem", key, value)
}
