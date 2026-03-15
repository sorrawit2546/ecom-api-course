package products

import (
	"context"

	repository "github.com/sorrawit2546/internal/adapters/postgresql/sqlc"
)

type IService interface {
	ListProducts(ctx context.Context) ([]repository.Product, error)
}

type service struct {
	repository repository.Queries
}

func NewService(r repository.Queries) *service {
	return &service{
		repository: r,
	}
}

func (s *service) ListProducts(ctx context.Context) ([]repository.Product, error) {
	products, err := s.repository.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
