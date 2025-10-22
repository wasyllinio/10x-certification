package events

import (
	"context"
)

// EventDispatcher represents the event dispatcher interface
type EventDispatcher interface {
	Dispatch(ctx context.Context, event BaseEvent) error
	Subscribe(eventType string, handler EventHandler)
}

// EventHandler represents an event handler interface
type EventHandler interface {
	Handle(ctx context.Context, event BaseEvent) error
}

// eventDispatcher represents the implementation of event dispatcher
type eventDispatcher struct {
	handlers map[string][]EventHandler
}

// NewEventDispatcher creates a new event dispatcher
func NewEventDispatcher() EventDispatcher {
	return &eventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

// Dispatch dispatches an event to all registered handlers
func (d *eventDispatcher) Dispatch(ctx context.Context, event BaseEvent) error {
	eventType := event.GetEventType()
	handlers := d.handlers[eventType]

	for _, handler := range handlers {
		if err := handler.Handle(ctx, event); err != nil {
			return err
		}
	}

	return nil
}

// Subscribe subscribes a handler to an event type
func (d *eventDispatcher) Subscribe(eventType string, handler EventHandler) {
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}
