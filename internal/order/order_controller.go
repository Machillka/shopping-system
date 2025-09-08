package order

func CreateNewOrder(itemName string, itemCount int) OrderModel {
	newOrder := OrderModel{
		Id:        0,
		ItemCount: itemCount,
	}

	return newOrder
}
