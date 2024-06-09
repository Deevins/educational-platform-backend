package fixtures

import (
	"time"

	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
)

type OrderBuilder struct {
	instance *models.Order
}

func Order() *OrderBuilder {
	return &OrderBuilder{
		instance: &models.Order{},
	}
}

func (ob *OrderBuilder) ID(v models.OrderID) *OrderBuilder {
	ob.instance.ID = v
	return ob
}

func (ob *OrderBuilder) UserID(v models.UserID) *OrderBuilder {
	ob.instance.UserID = v
	return ob
}
func (ob *OrderBuilder) Quantity(v models.Quantity) *OrderBuilder {
	ob.instance.Quantity = v
	return ob
}
func (ob *OrderBuilder) ProductName(v models.ProductName) *OrderBuilder {
	ob.instance.ProductName = v
	return ob
}
func (ob *OrderBuilder) Status(v models.OrderStatus) *OrderBuilder {
	ob.instance.Status = v
	return ob
}
func (ob *OrderBuilder) OrderDate(v time.Time) *OrderBuilder {
	ob.instance.OrderDate = v
	return ob
}

func (ob *OrderBuilder) P() *models.Order {

	return ob.instance
}
func (ob *OrderBuilder) V() models.Order {
	return *ob.instance
}
func (ob *OrderBuilder) Valid() *OrderBuilder {
	return ob.
		ID(1).
		UserID(1).
		ProductName("Product").
		Status("approved").
		Quantity(1).
		OrderDate(time.Now())
}
