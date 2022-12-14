package v1

import (
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/orders/dtos"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreateOrderCommand struct {
	OrderID         uuid.UUID           `validate:"required"`
	ShopItems       []*dtos.ShopItemDto `validate:"required"`
	AccountEmail    string              `validate:"required,email"`
	DeliveryAddress string              `validate:"required"`
	DeliveryTime    time.Time           `validate:"required"`
	CreatedAt       time.Time           `validate:"required"`
}

func NewCreateOrderCommand(shopItems []*dtos.ShopItemDto, accountEmail, deliveryAddress string, deliveryTime time.Time) *CreateOrderCommand {
	return &CreateOrderCommand{OrderID: uuid.NewV4(), ShopItems: shopItems, AccountEmail: accountEmail, DeliveryAddress: deliveryAddress, DeliveryTime: deliveryTime, CreatedAt: time.Now()}
}
