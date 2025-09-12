package domain

import "context"

// 定义持久化层对于 购物车订单的 接口
type OrderRepository interface {
	Save(ctx context.Context,o *Order) error
	FindById(ctx context.Context, id string) (*Order, error)
}