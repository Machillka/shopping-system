package domain

type OrderItem struct {
	SKU       string  // 标识
	UnitPrice float32 // 单价
	Quantity  int     // 购买数量
}

func (oi *OrderItem) TotalPrice() float32 {
	return oi.UnitPrice * float32(oi.Quantity)
}
