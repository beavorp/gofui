package router

import (
	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/element"
)

type MiddleWareFunc func(next func())

type route struct {
	path       string
	e          element.Element
	middleware []MiddleWareFunc
}

var routes = make(map[string]*route, 0)

func Register(path string, e element.Element, middleware ...MiddleWareFunc) {
	routes[path] = &route{
		path:       path,
		e:          e,
		middleware: middleware,
	}
}

func PathName() string {
	return core.Location.Get("pathname").String()
}

func Dispatch() {
	path := PathName()
	match, ok := routes[path]
	if !ok {
		return
	}

	shouldContinue := false
	nextFunc := func() { shouldContinue = true }
	for _, m := range match.middleware {
		shouldContinue = false
		if m(nextFunc); !shouldContinue {
			return
		}
	}
	// We render the element after all middleware has been executed
	match.e.Render().Ref()
}

func Navigate(path string) {
	core.History.Call("pushState", nil, "", path)
	Dispatch()
}
