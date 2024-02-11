package node

import "syscall/js"

func String(content string) *js.Value {
	v := js.ValueOf(content)
	return &v
}
