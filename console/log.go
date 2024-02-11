package console

import "syscall/js"

var (
	console    = js.Global().Get("console")
	consoleLog = console.Get("log")
)

func Log(args ...interface{}) {
	consoleLog.Invoke(args...)
}
