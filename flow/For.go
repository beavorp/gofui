package flow

import "syscall/js"

func For[T any](arr []T, f func(idx int, el T) *js.Value) []*js.Value {
	res := make([]*js.Value, 0, len(arr))

	for i, v := range arr {
		res = append(res, f(i, v))
	}

	return res
}
