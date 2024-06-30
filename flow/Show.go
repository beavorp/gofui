package flow

import (
	"syscall/js"

	"github.com/beavorp/gofui/element"
)

func ShowIf(when bool, a element.Element) *js.Value {
	if when {
		if a == nil {
			return nil
		}

		return a.Render().Ref()
	}

	return nil
}
