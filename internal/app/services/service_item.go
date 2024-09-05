package services

import (
	"context"

	"github.com/KozlovNikolai/marketplace/internal/app/domain"
)

// ItemService is a Item service
type ItemService struct {
	repo IItemRepository
}

// GetItem implements httpserver.IItemService.
func (s ItemService) GetItem(ctx context.Context, id int) (domain.Item, error) {
	return s.repo.GetItem(ctx, id)
}

// NewItemService creates a new Item service
func NewItemService(repo IItemRepository) ItemService {
	return ItemService{
		repo: repo,
	}
}

func (s ItemService) CreateItem(ctx context.Context, item domain.Item) (domain.Item, error) {
	orderChannels[item.OrderID()] <- struct{}{}
	return s.repo.CreateItem(ctx, item)
}

func (s ItemService) UpdateItem(ctx context.Context, item domain.Item) (domain.Item, error) {
	orderChannels[item.OrderID()] <- struct{}{}
	return s.repo.UpdateItem(ctx, item)
}

func (s ItemService) DeleteItem(ctx context.Context, id int) error {
	item, _ := s.GetItem(ctx, id)
	orderChannels[item.OrderID()] <- struct{}{}
	return s.repo.DeleteItem(ctx, id)
}

func (s ItemService) GetItems(ctx context.Context, limit, offset, orderid int) ([]domain.Item, error) {
	return s.repo.GetItems(ctx, limit, offset, orderid)
}
