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

func (s *Storage) GetItem(key string) string {
	return s.ref.Call("getItem", key).String()
}

func (s *Storage) SetItem(key, value string) {
	s.ref.Call("setItem", key, value)
}
