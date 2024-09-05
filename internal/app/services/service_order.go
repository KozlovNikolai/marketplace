package services

import (
	"context"
	"log"
	"time"

	"github.com/KozlovNikolai/marketplace/internal/app/domain"
	"github.com/KozlovNikolai/marketplace/internal/pkg/config"
)

var (
	orderChannels = make(map[int]chan struct{})
)

// OrderService is a Order service
type OrderService struct {
	repo IOrderRepository
}

// NewOrderService creates a new Order service
func NewOrderService(repo IOrderRepository) OrderService {
	return OrderService{
		repo: repo,
	}
}

func (s OrderService) GetOrder(ctx context.Context, id int) (domain.Order, error) {
	return s.repo.GetOrder(ctx, id)
}

func (s OrderService) GetOrdersByUserID(ctx context.Context, userID, limit, offset int) ([]domain.Order, error) {
	return s.repo.GetOrdersByUserID(ctx, userID, limit, offset)
}

func (s OrderService) CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	var newOrder = domain.NewOrderData{
		UserID:      order.UserID(),
		StateID:     domain.CreatedOrderStateID,
		TotalAmount: 0,
		CreatedAt:   time.Now(),
	}
	creatingOrder, err := domain.NewOrder(newOrder)
	if err != nil {
		return domain.Order{}, err
	}
	domainOrder, err := s.repo.CreateOrder(ctx, creatingOrder)
	if err == nil {
		orderChannels[domainOrder.ID()] = make(chan struct{})
		go s.startOrderTimer(ctx, domainOrder)
		log.Printf("create order %d\n", domainOrder.ID())
	}
	return domainOrder, err
}

func (s OrderService) UpdateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {

	return s.repo.UpdateOrder(ctx, order)
}

func (s OrderService) DeleteOrder(ctx context.Context, id int) error {
	return s.repo.DeleteOrder(ctx, id)
}

func (s OrderService) GetOrders(ctx context.Context, limit, offset, userid int) ([]domain.Order, error) {
	return s.repo.GetOrders(ctx, limit, offset, userid)
}

func (s OrderService) startOrderTimer(ctx context.Context, order domain.Order) {
	for {
		select {
		case <-time.After(config.Cfg.OrderDurationTime):
			log.Printf("change state order %d\n", order.ID())
			updatedOrder, _ := domain.NewOrder(domain.NewOrderData{
				ID:          order.ID(),
				UserID:      order.UserID(),
				StateID:     domain.InProgressOrderStateID,
				TotalAmount: order.TotalAmount(),
				CreatedAt:   order.CreatedAt(),
			})
			s.repo.UpdateOrder(ctx, updatedOrder)
			return
		case <-orderChannels[order.ID()]:
			log.Printf("update timer order %d\n", order.ID())
			continue
		}
	}
}
