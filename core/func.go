package core

import "syscall/js"

func CreateElement(tag string) js.Value {
	return Document.Call("createElement", tag)
}
