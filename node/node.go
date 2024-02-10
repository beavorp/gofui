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
