package headless

import (
	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/node"
)

type TextField struct {
	ref *node.Input
	// props
	onInput            func(event *core.Event)
	onValidationChange func(valid bool)

	rules []ValidationRule[string]
	valid bool
	error string
}

func NewTextField() (t *TextField) {
	t = &TextField{
		rules:              make([]ValidationRule[string], 0),
		ref:                node.NewInput().Type("text"),
		onInput:            func(event *core.Event) {},
		onValidationChange: func(valid bool) {},
	}

	t.ref.OnInput(func(e *core.Event) {
		t.onInput(e)
		t.Validate()
	})

	t.Validate()

	return
}

func (t *TextField) InputRef() *node.Input {
	return t.ref
}

func (t *TextField) Rules(rules ...ValidationRule[string]) *TextField {
	t.rules = rules
	return t
}

func (t *TextField) OnInput(f func(event *core.Event)) *TextField {
	if f == nil {
		return t
	}

	t.onInput = f
	return t
}

func (t *TextField) Validate() {
	v := t.Value()
	for _, r := range t.rules {
		if satisfied, msg := r(v); !satisfied {
			// Avoid unnecessary rendering
			if !t.valid && t.error == msg {
				return
			}

			t.valid = false
			t.error = msg
			t.onValidationChange(false)
			return
		}
	}

	// Avoid unnecessary rendering
	if t.valid {
		return
	}

	t.valid = true
	t.error = ""
	t.onValidationChange(true)

	return
}

func (t *TextField) Value() string {
	return t.ref.Ref().Get("value").String()
}

func (t *TextField) SetValue(value string) {
	t.ref.Ref().Set("value", value)
	t.Validate()
}

func (t *TextField) IsValid() bool {
	return t.valid
}

func (t *TextField) Error() string {
	return t.error
}

func (t *TextField) OnValidationChange(f func(valid bool)) *TextField {
	if f == nil {
		return t
	}

	t.onValidationChange = f

	return t
}
