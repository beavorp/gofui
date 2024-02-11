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

func (d *Div) SetChildren(children []*js.Value) *Div {
	replaceChildren(*d.ref, children)
	return d
}

func (d *Div) Render() element.Element {
	return d
}

func (d *Div) Value() *js.Value {
	return d.ref
}

func (d *Div) SetId(id string) {
	d.ref.Set("id", id)
}

// C is a shorthand for SetChildren
func (d *Div) C(els ...*js.Value) *js.Value {
	replaceChildrenV2(*d.ref, els...)
	return d.Value()
}

func (d *Div) Class(class string) *Div {
	d.Style.SetClassName(class)
	return d
}
