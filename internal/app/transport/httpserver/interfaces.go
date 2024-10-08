//go:generate mockery

package httpserver

import (
	"context"

	"github.com/KozlovNikolai/marketplace/internal/app/domain"
)

// IUserService is ...
type IUserService interface {
	CreateUser(context.Context, domain.User) (domain.User, error)
	GetUsers(context.Context, int, int) ([]domain.User, error)
	GetUserByID(context.Context, int) (domain.User, error)
	GetUserByLogin(context.Context, string) (domain.User, error)
	UpdateUser(context.Context, domain.User) (domain.User, error)
	DeleteUser(context.Context, int) error
}

// TokenService is a token service
type ITokenService interface {
	GenerateToken(user domain.User) (string, error)
	GetUser(token string) (domain.User, error)
}

// IProviderService is ...
type IProviderService interface {
	CreateProvider(context.Context, domain.Provider) (domain.Provider, error)
	GetProviders(context.Context, int, int) ([]domain.Provider, error)
	GetProvider(context.Context, int) (domain.Provider, error)
	UpdateProvider(context.Context, domain.Provider) (domain.Provider, error)
	DeleteProvider(context.Context, int) error
}

// IProductService is ...
type IProductService interface {
	CreateProduct(context.Context, domain.Product) (domain.Product, error)
	GetProducts(context.Context, int, int) ([]domain.Product, error)
	GetProduct(context.Context, int) (domain.Product, error)
	UpdateProduct(context.Context, domain.Product) (domain.Product, error)
	DeleteProduct(context.Context, int) error
}

// IOrderService is ...
type IOrderService interface {
	CreateOrder(context.Context, domain.Order) (domain.Order, error)
	GetOrders(context.Context, int, int, int) ([]domain.Order, error)
	GetOrder(context.Context, int) (domain.Order, error)
	GetOrdersByUserID(context.Context, int, int, int) ([]domain.Order, error)
	UpdateOrder(context.Context, domain.Order) (domain.Order, error)
	DeleteOrder(context.Context, int) error
}

// IOrderStateService is ...
type IOrderStateService interface {
	CreateOrderState(context.Context, domain.OrderState) (domain.OrderState, error)
	GetOrderStates(context.Context, int, int) ([]domain.OrderState, error)
	GetOrderState(context.Context, int) (domain.OrderState, error)
	UpdateOrderState(context.Context, domain.OrderState) (domain.OrderState, error)
	DeleteOrderState(context.Context, int) error
}

// IItemService is ...
type IItemService interface {
	CreateItem(context.Context, domain.Item) (domain.Item, error)
	GetItems(context.Context, int, int, int) ([]domain.Item, error)
	GetItem(context.Context, int) (domain.Item, error)
	UpdateItem(context.Context, domain.Item) (domain.Item, error)
	DeleteItem(context.Context, int) error
}
