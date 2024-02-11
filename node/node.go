package node

import "syscall/js"

func replaceChildren(el js.Value, children []*js.Value) {
	vals := make([]interface{}, 0, len(children))
	for _, child := range children {
		if child == nil {
			continue
		}

		vals = append(vals, *child)
	}

	el.Call("replaceChildren", vals...)
}

func replaceChildrenV2(el js.Value, els ...*js.Value) {
	children := make([]interface{}, 0, len(els))
	for _, child := range els {
		if child == nil {
			continue
		}
		children = append(children, *child)
	}

	el.Call("replaceChildren", children...)
}
