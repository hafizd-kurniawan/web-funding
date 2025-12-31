package worker

import (
	"context"
	"funding/internal/application/event"
	"funding/internal/infrastructure/persistance/repositoy"
)

type Worker interface {
	Run()
}

type OutboxWorker struct {
	publisher event.Publisher
	repo      repositoy.WorkerRepository
}

var _ Worker = (*OutboxWorker)(nil)

func NewOutboxWorker() *OutboxWorker {
	return &OutboxWorker{}
}

func (w *OutboxWorker) Run() {
	ctx := context.Background()
	events := w.repo.FetchUnprocessed(ctx)
	for _, event := range events {
		w.publisher.Publish(ctx, event)
		w.repo.MarksProcessed(ctx, 1)
	}
}
