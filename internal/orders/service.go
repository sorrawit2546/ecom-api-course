package orders

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	repository "github.com/sorrawit2546/internal/adapters/postgresql/sqlc"
)

var (
	ErrProductNotFound = errors.New("Product not found")
)

// interface
type IOrderService interface {
	placeOrder(ctx context.Context, tempOrder createOrderParams) (repository.Order, error)
}

// struct
type OrderService struct {
	repo repository.Queries
	db   *pgx.Conn
}

// constructor
func NewOrder(r repository.Queries, db *pgx.Conn) *OrderService {
	return &OrderService{
		repo: r,
		db: db,
	}
}

func (s *OrderService) placeOrder(ctx context.Context, tempOrder createOrderParams) (repository.Order, error) {
	if tempOrder.CustomerId == 0 {
		return repository.Order{}, fmt.Errorf("Customer ID is required")
	}
	if len(tempOrder.Items) == 0 {
		return repository.Order{}, fmt.Errorf("Order Item is required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repository.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerId)
	if err != nil {
		return repository.Order{}, err
	}

	for _, item := range tempOrder.Items {
		product, err := qtx.ListProduct(ctx, item.ProductID)
		if err != nil {
			return repository.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repository.Order{}, err
		}

		_, err = qtx.CreateOrderItem(ctx, repository.CreateOrderItemParams{
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			PriceCents: product.PriceInCents,
		})

		if err != nil {
			return repository.Order{}, err
		}
	}

	tx.Commit(ctx)
	return order, nil
}
