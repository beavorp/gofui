package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/element"
)

type Div struct {
	ref   *js.Value
	Style *element.Style
}

func NewDiv() *Div {
	el := js.Global().Get("document").Call("createElement", "div")
	return &Div{
		ref:   &el,
		Style: element.NewStyle(el),
	}
}

func (d *Div) SetChildren(children []*js.Value) {
	replaceChildren(*d.ref, children)
}

func (d *Div) Render() element.Element {
	return d
}

func (d *Div) Value() *js.Value {
	return d.ref
}
