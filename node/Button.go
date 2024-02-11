package node

import (
	"fmt"
	"syscall/js"

	"github.com/beavorp/gofui/element"
)

type Button struct {
	el    *js.Value
	Style *element.Style

	OnClick func()
}

func NewButton(html string) *Button {
	el := js.Global().Get("document").Call("createElement", "button")
	b := &Button{
		el:    &el,
		Style: element.NewStyle(el),
	}

	el.Set("innerHTML", html)
	el.Set("onclick", js.FuncOf(b.onClick))

	return b
}

func (b *Button) onClick(this js.Value, args []js.Value) any {
	if b.OnClick != nil {
		b.OnClick()
	}
	return nil
}

func (b *Button) Render() element.Element {
	return b
}

func (b *Button) Value() *js.Value {
	return b.el
}

func (b *Button) SetId(id string) {
	b.el.Set("id", id)
}

func (b *Button) SetDisabled(disabled bool) {
	fmt.Println("Setting disabled", disabled)
	b.el.Set("disabled", disabled)
}
