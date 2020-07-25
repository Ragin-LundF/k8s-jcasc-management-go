package events

import "time"

var NamespaceCreated namespaceCreated

// NamespaceCreatedPayload is the data for when a namespace is created
type NamespaceCreatedPayload struct {
	Namespace string
	Time      time.Time
}

type namespaceCreated struct {
	handlers []interface{ Handle(NamespaceCreatedPayload) }
}

// Register adds an event handler for this event
func (namespaceCreated *namespaceCreated) Register(handler interface{ Handle(NamespaceCreatedPayload) }) {
	namespaceCreated.handlers = append(namespaceCreated.handlers, handler)
}

// Trigger sends out an event with the payload
func (namespaceCreated namespaceCreated) Trigger(payload NamespaceCreatedPayload) {
	for _, handler := range namespaceCreated.handlers {
		go handler.Handle(payload)
	}
}
