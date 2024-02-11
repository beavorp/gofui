package cycle

import "syscall/js"

var (
	documentReady          = false
	documentReadyCallbacks = make([]func(), 0, 10) // preallocate 10 slots
)

func init() {
	js.Global().Get("window").Set("gofui", js.FuncOf(_onDocumentReady))
}

func _onDocumentReady(this js.Value, args []js.Value) any {
	documentReady = true
	for _, callback := range documentReadyCallbacks {
		callback()
	}

	return nil
}

func OnDocumentReady(callback func()) {
	if documentReady {
		callback()
		return
	}

	documentReadyCallbacks = append(documentReadyCallbacks, callback)
}
