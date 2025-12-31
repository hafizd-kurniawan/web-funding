package event

import (
	"context"
	"funding/internal/application/event"
)

type InMemoryBus struct {
	handlers []event.Handler
}

var _ event.Publisher = (*InMemoryBus)(nil)

func NewInMemoryBus(handlers ...event.Handler) *InMemoryBus {
	return &InMemoryBus{handlers: handlers}
}

func (b *InMemoryBus) Publish(ctx context.Context, event any) error {
	for _, h := range b.handlers {
		if err := h.Handle(ctx, event); err != nil {
			return err
		}
	}
	return nil
}
