package main

import (
	"syscall/js"

	"github.com/beavorp/gofui"
	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/element"
	"github.com/beavorp/gofui/flow"
	"github.com/beavorp/gofui/node"
)

func main() {
	// New body will not create a new body in the DOM,
	// it will get the reference to the existing body.
	body := node.NewBody()

	app := NewApp(body)
	app.Render()

	// This is a blocking call, it will keep the app running
	// and listening to events.
	<-make(chan interface{})
}

// App is the main component of the app.
// A component is a struct that implements the element.Element interface.
// It has a Render method that returns the element.Element and a Ref method that returns the *js.Value of the component.
type App struct {
	// ref to node.Body
	ref *node.Body

	displayMessage *gofui.State[bool]
	message        *node.Span
}

func NewApp(ref *node.Body) *App {
	a := &App{ref: ref}
	// Create a new state. We pass the component and the initial value of the state.
	// The state will automatically call the Render method of the component when the value changes.
	// The initial value of the state is false.
	a.displayMessage = gofui.UseState(a, false)

	// Create a new span node
	a.message = node.NewSpan()
	// We can set the children of the span node using the C method.
	// This span will never change, so we can set the children here.
	// Also, it will be hidden or showed based on the value of the state.
	a.message.C(node.String("Hello world"))

	return a
}

func (a *App) Render() element.Element {
	// Will update the children of the body
	a.ref.C(
		// Create a new button node
		node.NewButton().OnClick(func(e *core.PointerEvent) {
			a.displayMessage.Set(!a.displayMessage.Get())
		}).C(
			node.String("Click me"),
		),

		// Show the message if the state is true
		flow.ShowIf(a.displayMessage.Get(), a.message),
	)

	return a
}

func (a *App) Ref() *js.Value {
	return a.ref.Ref()
}
