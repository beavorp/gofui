package element

import "syscall/js"

type Element interface {
	Render() Element
	Ref() *js.Value
}
