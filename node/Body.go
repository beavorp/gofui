package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/element"
)

type Body struct {
	ref   *js.Value
	Style *element.Style
}

func NewBody() *Body {
	el := js.Global().Get("document").Get("body")
	style := element.NewStyle(el)

	return &Body{
		ref:   &el,
		Style: style,
	}
}

func (b *Body) SetChildren(children []*js.Value) {
	replaceChildren(*b.ref, children)
}

func (b *Body) Value() *js.Value {
	return b.ref
}
