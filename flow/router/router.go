package router

import (
	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/element"
)

type route struct {
	path       string
	e          element.Element
	middleware []func()
}

var routes = make(map[string]*route, 0)

func Register(path string, e element.Element, middleware ...func()) {
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

	for _, m := range match.middleware {
		m()
	}

	match.e.Render().Value()
}

func Navigate(path string) {
	core.History.Call("pushState", nil, "", path)
	Dispatch()
}
