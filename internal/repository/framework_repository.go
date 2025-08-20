package repository

import (
	"context"

	"github.com/aakarsh-kamboj/echo-practise/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FrameworkRepository interface {
	CreateFramework(ctx context.Context, name, description, version string, locked, editable bool, category []string) (db.Framework, error)
	GetFrameworkByID(ctx context.Context, id uuid.UUID) (db.Framework, error)
	UpdateFramework(ctx context.Context, id uuid.UUID, name, description, version string, locked, editable bool, category []string) (db.Framework, error)
	DeleteFramework(ctx context.Context, id uuid.UUID) (string, error)
	ListFrameworks(ctx context.Context) ([]db.Framework, error)
}

type frameworkRepository struct {
	q *db.Queries
}

func NewFrameworkRepository(pool *pgxpool.Pool) FrameworkRepository {
	q := db.New(pool)
	return &frameworkRepository{
		q: q,
	}
}

func (r *frameworkRepository) CreateFramework(ctx context.Context, name, description, version string, locked, editable bool, category []string) (db.Framework, error) {
	return r.q.CreateFramework(ctx, db.CreateFrameworkParams{
		Name:        name,
		Description: pgtype.Text{String: description},
		Version:     pgtype.Text{String: version},
		Locked:      locked,
		Editable:    editable,
		Category:    category,
	})
}

func (r *frameworkRepository) GetFrameworkByID(ctx context.Context, id uuid.UUID) (db.Framework, error) {
	return r.q.GetFramework(ctx, id)
}

func (r *frameworkRepository) UpdateFramework(ctx context.Context, id uuid.UUID, name, description, version string, locked, editable bool, category []string) (db.Framework, error) {
	return r.q.UpdateFramework(ctx, db.UpdateFrameworkParams{
		ID:          id,
		Name:        name,
		Description: pgtype.Text{String: description},
		Version:     pgtype.Text{String: description},
		Locked:      locked,
		Editable:    editable,
		Category:    category,
	})
}

func (r *frameworkRepository) DeleteFramework(ctx context.Context, id uuid.UUID) (string, error) {
	if err := r.q.DeleteFramework(ctx, id); err != nil {
		return "failed", err
	}
	return "success", nil
}

func (r *frameworkRepository) ListFrameworks(ctx context.Context) ([]db.Framework, error) {
	return r.q.ListFrameworks(ctx)
}
