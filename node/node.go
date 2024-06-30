package node

import (
	"syscall/js"

	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/element"
)

type Node[T element.Element] struct {
	// parent
	p T
	// ref
	ref   *js.Value
	Style *element.Style
}

func NewNode[T element.Element](parent T, v *js.Value) *Node[T] {
	return &Node[T]{
		p:     parent,
		ref:   v,
		Style: element.NewStyle(*v),
	}
}

func (n *Node[T]) Render() element.Element {
	return n
}

func (n *Node[T]) Ref() *js.Value {
	return n.ref
}

func (n *Node[T]) SetId(id string) {
	n.ref.Set("id", id)
}

func (n *Node[T]) C(els ...*js.Value) *js.Value {
	children := make([]interface{}, 0, len(els))
	for _, child := range els {
		if child == nil {
			continue
		}
		children = append(children, *child)
	}

	n.ref.Call("replaceChildren", children...)

	return n.Ref()
}

func (n *Node[T]) Class(class string) T {
	n.Style.Class(class)
	return n.p
}

func (n *Node[T]) Clsx(clsx map[string]bool) T {
	n.Style.Clsx(clsx)
	return n.p
}

func (n *Node[T]) OnClick(fn func(e *core.PointerEvent)) T {
	n.ref.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(core.NewPointerEvent(args[0]))
		return nil
	}))

	return n.p
}

func (n *Node[T]) OnInput(fn func(e *core.Event)) T {
	n.ref.Set("oninput", js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(core.NewEvent(args[0]))
		return nil
	}))

	return n.p
}

func (n *Node[T]) OnChange(fn func(e *core.Event)) T {
	n.ref.Call("addEventListener", "change", js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(core.NewEvent(args[0]))
		return nil
	}))

	return n.p
}

func (n *Node[T]) AddClass(className string) {
	n.Style.AddClassName(className)
}

func (n *Node[T]) RemoveClass(className string) {
	n.Style.RemoveClassName(className)
}

func (n *Node[T]) Attr(key string, value string) {
	n.ref.Set(key, value)
}

// P Returns the parent element where the node is attached to
// If you are using this on a component, it will return the component itself.
func (n *Node[T]) P() T {
	return n.p
}
