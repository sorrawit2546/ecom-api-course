package products

import "context"

type IService interface {
	ListProducts(ctx context.Context) error
}

type service struct {
	repository IRepository
}

func NewService(r IRepository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) ListProducts(ctx context.Context) error {
	return  nil
}