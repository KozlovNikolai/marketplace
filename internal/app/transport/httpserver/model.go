package httpserver

import (
	"time"

	"github.com/go-playground/validator/v10"
)

//const passwordLength = 6

// ProviderRequest is ...
type ProviderRequest struct {
	Name   string `json:"name" db:"name" example:"Microsoft" validate:"required"`
	Origin string `json:"origin" db:"origin" example:"Vietnam" validate:"required"`
}

func (p *ProviderRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(p)
}

type ProviderResponse struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Origin string `json:"origin" db:"origin"`
}

// #######################################################################################3
// ProductRequest is ...
type ProductRequest struct {
	Name       string  `json:"name" db:"name" example:"синхрофазотрон" validate:"required"`
	ProviderID int     `json:"provider_id" db:"provider_id" example:"1" validate:"gt=0"`
	Price      float64 `json:"price" db:"price" example:"1245.65" validate:"gte=0"`
	Stock      int     `json:"stock" db:"stock" example:"435" validate:"gte=0"`
}

func (p *ProductRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(p)
}

type ProductResponse struct {
	ID         int     `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	ProviderID int     `json:"provider_id" db:"provider_id"`
	Price      float64 `json:"price" db:"price"`
	Stock      int     `json:"stock" db:"stock"`
}

// #######################################################################################
// OrderStateRequest is ...
type OrderStateRequest struct {
	Name string `json:"name" db:"name" example:"в обработке" validate:"required"`
}

func (o *OrderStateRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(o)
}

type OrderStateResponse struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// #######################################################################################3
// ItemRequest is ...
type ItemRequest struct {
	ProductID int `json:"product_id" db:"product_id" example:"1" validate:"gt=0"`
	Quantity  int `json:"quantity" db:"quantity" example:"3" validate:"gt=0"`
	OrderID   int `json:"order_id" db:"order_id" example:"1" validate:"gt=0"`
}

func (i *ItemRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(i)
}

type ItemResponse struct {
	ID         int     `json:"id" db:"id"`
	ProductID  int     `json:"product_id" db:"product_id"`
	Quantity   int     `json:"quantity" db:"quantity"`
	TotalPrice float64 `json:"total_price" db:"total_price"`
	OrderID    int     `json:"order_id" db:"order_id"`
}

// #######################################################################################3
// OrderRequest is ...
type OrderRequest struct {
	UserID int `json:"user_id" db:"user_id" example:"1" validate:"gt=0"`
}

func (o *OrderRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(o)
}

type OrderResponse struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	StateID     int       `json:"state_id" db:"state_id"`
	TotalAmount float64   `json:"total_amount" db:"total_amount"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// #######################################################################################3
// UserRequest is ...
type UserRequest struct {
	Login    string `json:"login" db:"login" example:"cmd@cmd.ru" validate:"required,email"`
	Password string `json:"password" db:"password" example:"123456" validate:"required,min=6,max=32"`
}

func (u *UserRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(u)
}

type UserResponse struct {
	ID       int    `json:"id" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
	Token    string `json:"token" db:"token"`
}
