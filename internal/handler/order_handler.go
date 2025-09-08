package handler

import (
	"github.com/machillka/shopping-system/internal/order"
)

// 处理订单相关操作
// TODO: 加入可持续化操作

// 创建订单
func CreateOrderHandler(itemName string, itemCount int) error {
	newOrder := order.CreateNewOrder(itemName, itemCount)
	err := RemoveItemHandler(itemName, itemCount)
	if err != nil {
		return err
	}
	CallForPayment(newOrder)
	return nil
}

func CancleOrderHandler() {

}
