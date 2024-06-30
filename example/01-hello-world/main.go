package main

import (
	"syscall/js"

	"github.com/beavorp/gofui/element"
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
}

func NewApp(ref *node.Body) *App {
	return &App{ref: ref}
}

func (a *App) Render() element.Element {
	// Will update the children of the body
	a.ref.C(
		node.String("Hello world"),
	)

	return a
}

func (a *App) Ref() *js.Value {
	return a.ref.Ref()
}
