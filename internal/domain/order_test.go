package domain

import "testing"

func TestNewOrder(t *testing.T) {
	items := []OrderItem{
		{SKU: "TA", UnitPrice: 100, Quantity: 2},
		{SKU: "TB", UnitPrice: 200, Quantity: 1},
	}

	order := NewOrder("User1", items)

	if order.TotalAmount != 400 {
		t.Errorf("Expected total amount to be 400, got %.2f", order.TotalAmount)
	}
}
