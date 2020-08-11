package main

type event int

const (
	twoFiftySixSecondsPassed = iota
	talkToCountKidAfterFinished
	talkedToWellGuy
	talkedToOldGuy
	talkedToDog
)

var storyEventHandler = eventHandler{}

type eventHandler struct {
	handlers map[event][]func()
}

func initEventHandler() {
	storyEventHandler.handlers = make(map[event][]func())
}

func (h *eventHandler) sendEvent(e event) {
	for _, f := range h.handlers[e] {
		f()
	}
}

func (h *eventHandler) onEvent(e event, handler func()) {
	h.handlers[e] = append(h.handlers[e], handler)
}
