package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/element"
)

type Body struct {
	ref   *js.Value
	style *element.Style
}

func NewBody() *Body {
	el := js.Global().Get("document").Get("body")
	style := element.NewStyle(el)

	style.H_PaMa0() // padding and margin 0
	style.SetOverflow("hidden")

	return &Body{
		ref:   &el,
		style: style,
	}
}

func (b *Body) SetChildren(children []*js.Value) {
	replaceChildren(*b.ref, children)
}

func (b *Body) Value() *js.Value {
	return b.ref
}
