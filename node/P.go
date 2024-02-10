package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/element"
)

type P struct {
	// ref
	ref   *js.Value
	Style *element.Style
}

func NewP(text string) *P {
	el := js.Global().Get("document").Call("createElement", "p")
	el.Set("innerText", text)
	return &P{
		ref:   &el,
		Style: element.NewStyle(el),
	}
}

func (p *P) SetChildren(children []*js.Value) {
	replaceChildren(*p.ref, children)
}

func (p *P) Render() element.Element {
	return p
}

func (p *P) Value() *js.Value {
	return p.ref
}

func setInnerText(ref js.Value, text string) {
	ref.Set("innerText", text)
}
