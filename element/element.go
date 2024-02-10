package element

import "syscall/js"

type Element interface {
	Render() Element
	Value() *js.Value
}
