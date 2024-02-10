package flow

import (
	"syscall/js"

	"github.com/beavorp/gofui/element"
)

func Show(when bool, a element.Element) *js.Value {
	if when {
		if a == nil {
			return nil
		}

		return a.Render().Value()
	}

	return nil
}
