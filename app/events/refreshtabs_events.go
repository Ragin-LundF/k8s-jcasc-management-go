package events

import "time"

var RefreshTabs refreshTabs

// RefreshTabsPayload is the data for when a tab should be refreshed
type RefreshTabsPayload struct {
	Time time.Time
}

type refreshTabs struct {
	handlers []interface{ Handle(RefreshTabsPayload) }
}

// Register adds an event handler for this event
func (rt *refreshTabs) Register(handler interface{ Handle(RefreshTabsPayload) }) {
	rt.handlers = append(rt.handlers, handler)
}

// Trigger sends out an event with the payload
func (rt refreshTabs) Trigger(payload RefreshTabsPayload) {
	for _, handler := range rt.handlers {
		go handler.Handle(payload)
	}
}
