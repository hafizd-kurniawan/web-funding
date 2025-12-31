package postgres

import (
	"context"
	"funding/internal/infrastructure/persistance/repositoy"

	"github.com/jmoiron/sqlx"
)

type WorkerRepository struct {
	db *sqlx.DB
}

var _ repositoy.WorkerRepository = (*WorkerRepository)(nil)

func NewWorkerRepository(db *sqlx.DB) *WorkerRepository {
	return &WorkerRepository{db: db}
}

func (r *WorkerRepository) Save(ctx context.Context, eventType string, payload []byte) error {
	return nil
}

func (r *WorkerRepository) MarksProcessed(ctx context.Context, id int) error {
	return nil
}
