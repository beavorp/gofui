package core

import "syscall/js"

var (
	Document = js.Global().Get("document")
	Window   = js.Global().Get("window")
	Location = Window.Get("location")
	History  = Window.Get("history")
)
