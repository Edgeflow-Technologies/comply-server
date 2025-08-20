package service

import (
	"context"

	"github.com/aakarsh-kamboj/echo-practise/internal/db"
	"github.com/aakarsh-kamboj/echo-practise/internal/repository"
	"github.com/google/uuid"
)

type FrameworkService interface {
	CreateFramework(ctx context.Context, name, description, version string, locked, editable bool, category []string) (db.Framework, error)
	GetFrameworkByID(ctx context.Context, id uuid.UUID) (db.Framework, error)
	UpdateFramework(ctx context.Context, id uuid.UUID, name, description, version string, locked, editable bool, category []string) (db.Framework, error)
	DeleteFramework(ctx context.Context, id uuid.UUID) (string, error)
	ListFrameworks(ctx context.Context) ([]db.Framework, error)
}

type frameworkService struct {
	repo repository.FrameworkRepository
}

func NewFrameworkService(repo repository.FrameworkRepository) FrameworkService {
	return &frameworkService{
		repo: repo,
	}
}

func (s *frameworkService) CreateFramework(ctx context.Context, name, description, version string, locked, editable bool, category []string) (db.Framework, error) {
	return s.repo.CreateFramework(ctx, name, description, version, locked, editable, category)
}

func (s *frameworkService) GetFrameworkByID(ctx context.Context, id uuid.UUID) (db.Framework, error) {
	return s.repo.GetFrameworkByID(ctx, id)
}

func (s *frameworkService) UpdateFramework(ctx context.Context, id uuid.UUID, name, description, version string, locked, editable bool, category []string) (db.Framework, error) {
	return s.repo.UpdateFramework(ctx, id, name, description, version, locked, editable, category)
}

func (s *frameworkService) DeleteFramework(ctx context.Context, id uuid.UUID) (string, error) {
	return s.repo.DeleteFramework(ctx, id)
}

func (s *frameworkService) ListFrameworks(ctx context.Context) ([]db.Framework, error) {
	return s.repo.ListFrameworks(ctx)
}
