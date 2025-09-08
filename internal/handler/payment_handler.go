package handler

import "github.com/machillka/shopping-system/internal/order"

func CallForPayment(order order.OrderModel) error {
	totalPrice, err := CalculateTotalPrice(order.ItemName, order.ItemCount)
	if err != nil {
		return nil
	}
	// TODO: 模拟用户支付
	_ = totalPrice
	return nil
}
