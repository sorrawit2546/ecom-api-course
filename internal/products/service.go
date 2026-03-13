package products

import (
	"context"

	repository "github.com/sorrawit2546/internal/adapters/postgresql/sqlc"
)

type IService interface {
	ListProducts(ctx context.Context) error
}

type service struct {
	repository repository.Queries
}

func NewService(r repository.Queries) *service {
	return &service{
		repository: r,
	}
}

func (s *service) ListProducts(ctx context.Context) error {
	_, err := s.repository.ListProducts(ctx)
	if err != nil {
		return err
	}
	return nil
}