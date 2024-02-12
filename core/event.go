package core

import "syscall/js"

type (
	Event struct {
		ref js.Value
	}

	UIEvent struct {
		*Event
	}

	MouseEvent struct {
		*UIEvent
	}

	PointerEvent struct {
		*MouseEvent
	}
)

func NewEvent(ref js.Value) *Event {
	return &Event{
		ref: ref,
	}
}

func NewUIEvent(ref js.Value) *UIEvent {
	return &UIEvent{
		Event: NewEvent(ref),
	}
}

func NewMouseEvent(ref js.Value) *MouseEvent {
	return &MouseEvent{
		UIEvent: NewUIEvent(ref),
	}
}

func NewPointerEvent(ref js.Value) *PointerEvent {
	return &PointerEvent{
		MouseEvent: NewMouseEvent(ref),
	}
}

// GetClientX provides the horizontal coordinate within the application's viewport
// at which the event occurred (as opposed to the coordinate within the page).
// https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/clientX
func (m *MouseEvent) GetClientX() int {
	return m.ref.Get("clientX").Int()
}

// GetClientY provides the vertical coordinate within the application's viewport
// at which the event occurred (as opposed to the coordinate within the page).
// https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/clientY
func (m *MouseEvent) GetClientY() int {
	return m.ref.Get("clientY").Int()
}

// PreventDefault tells the user agent that if the event does not get explicitly handled,
// its default action should not be taken as it normally would be.
// https://developer.mozilla.org/en-US/docs/Web/API/Event/preventDefault
func (e *Event) PreventDefault() {
	e.ref.Call("preventDefault")
}

// StopPropagation prevents further propagation of the current event in the capturing and bubbling phases.
// It does not, however, prevent any default behaviors from occurring;
// for instance, clicks on links are still processed.
// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopPropagation
func (e *Event) StopPropagation() {
	e.ref.Call("stopPropagation")
}
