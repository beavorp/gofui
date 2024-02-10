package router

import (
	"github.com/beavorp/gofui/core"
	"github.com/beavorp/gofui/element"
)

var routes = make(map[string]element.Element)

func Register(path string, e element.Element) {
	routes[path] = e
}

func PathName() string {
	return core.Location.Get("pathname").String()
}

func Dispatch() {
	path := PathName()
	route, ok := routes[path]
	if !ok {
		return
	}

	route.Render().Value()
}

func Navigate(path string) {
	core.History.Call("pushState", nil, "", path)
	Dispatch()
}
