package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "CREATED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
	OrderStatusCompleted OrderStatus = "COMPLETED"
)

const (
// 错误定义
)

// 用户的购物车结算所下的总订单
type Order struct {
	ID          string
	UserID      string
	Items       []OrderItem
	TotalAmount float32
	Status      OrderStatus
	CreateAt    time.Time
	UpdateAt    time.Time
}

func (o *Order) calculateTotalAmount() {
	var total float32
	for _, item := range o.Items {
		total += item.TotalPrice()
	}
	o.TotalAmount = total
	o.touch() // update订单的更新时间
}

// 更新订单的更新时间戳
func (o *Order) touch() {
	o.UpdateAt = time.Now().UTC()
}

// 创建 Order 聚合的工厂函数
func NewOrder(userId string, items []OrderItem) *Order {
	now := time.Now()
	order := &Order{
		ID:       uuid.New().String(),
		UserID:   userId,
		Items:    items,
		Status:   OrderStatusCreated,
		CreateAt: now,
		UpdateAt: now,
	}
	order.calculateTotalAmount()
	return order
}

func (o *Order) Cancel() error {
	if o.Status != OrderStatusCompleted {
		return errors.New("订单尚未创建")
	}

	o.Status = OrderStatusCancelled
	o.touch()
	return nil
}

func (o *Order) Complete() error {
	if o.Status != OrderStatusCreated {
		return errors.New("订单尚未创建")
	}

	o.Status = OrderStatusCompleted
	o.touch()
	return nil
}
