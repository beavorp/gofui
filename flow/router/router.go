package router

import (
	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/element"
)

type MiddleWareFunc func(next MiddleWareFunc)

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

	for i, m := range match.middleware {
		if i == len(match.middleware)-1 {
			m(func(next MiddleWareFunc) {
				match.e.Render().Value()
			})
			break
		}
		next := match.middleware[i+1]
		m(next)
	}
}

func Navigate(path string) {
	core.History.Call("pushState", nil, "", path)
	Dispatch()
}
