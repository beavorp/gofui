package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/element"
)

type Node struct {
	// ref
	ref   *js.Value
	Style *element.Style
}

func NewNode(v *js.Value) *Node {
	return &Node{
		ref:   v,
		Style: element.NewStyle(*v),
	}
}

func (n *Node) Render() element.Element {
	return n
}

func (n *Node) Value() *js.Value {
	return n.ref
}

func (n *Node) SetId(id string) {
	n.ref.Set("id", id)
}

func (n *Node) C(els ...*js.Value) *js.Value {
	children := make([]interface{}, 0, len(els))
	for _, child := range els {
		if child == nil {
			continue
		}
		children = append(children, *child)
	}

	n.ref.Call("replaceChildren", children...)

	return n.Value()
}

func (n *Node) Class(class string) *Node {
	n.Style.SetClassName(class)
	return n
}

func (n *Node) OnClick(fn func(e *core.PointerEvent)) *Node {
	n.ref.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(core.NewPointerEvent(args[0]))
		return nil
	}))

	return n
}

func (n *Node) Ref() *js.Value {
	return n.ref
}
