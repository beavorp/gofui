package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/element"
)

type (
	H2 struct {
		// ref
		ref   *js.Value
		Style *element.Style
	}
)

func NewH2(text string) *H2 {
	el := core.Document.Call("createElement", "h2")
	el.Set("innerText", text)
	return &H2{
		ref:   &el,
		Style: element.NewStyle(el),
	}
}

func (h *H2) Render() element.Element {
	return h
}
func (h *H2) Value() *js.Value {
	return h.ref
}
func (h *H2) SetInnerText(text string) {
	h.ref.Set("innerText", text)
}
