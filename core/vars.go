package core

import "syscall/js"

var (
	Document = js.Global().Get("document")
	Window   = js.Global().Get("window")
	Location = js.Global().Get("window").Get("location")
	History  = js.Global().Get("window").Get("history")
)
