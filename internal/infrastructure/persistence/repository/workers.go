package repository

import "context"

/**
 * Workers
CREATE TABLE outbox_events (
	id SERIAL PRIMARY KEY,
	event_type TEXT NOT NULL,
	payload JSONB NOT NULL,
	created_at TIMESTAMP NOT NULL,
	processed BOOLEAN DEFAULT FALSE
);
*/

type WorkerRepository interface {
	Save(ctx context.Context, eventType string, payload []byte) error
	MarksProcessed(ctx context.Context, id int) error
	FetchUnprocessed(ctx context.Context) ([]string, error)
}
