package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/element"
)

type (
	H1 struct {
		// ref
		ref   *js.Value
		Style *element.Style
	}

	H2 struct {
		// ref
		ref   *js.Value
		Style *element.Style
	}
)

func NewH1(text string) *H1 {
	el := core.Document.Call("createElement", "h1")
	el.Set("innerText", text)
	return &H1{
		ref:   &el,
		Style: element.NewStyle(el),
	}
}

func NewH2(text string) *H2 {
	el := core.Document.Call("createElement", "h2")
	el.Set("innerText", text)
	return &H2{
		ref:   &el,
		Style: element.NewStyle(el),
	}
}

func (h *H1) Render() element.Element {
	return h
}
func (h *H1) Value() *js.Value {
	return h.ref
}
func (h *H1) SetInnerText(text string) {
	h.ref.Set("innerText", text)
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
