package item

// 维护出现过的商品列表 [商品名] -> 商品模型
var ItemTable = map[string]ItemModel{}

// TODO: 使用高级的ID生成方式, 目前不适用该字段
// 创建新商品的同时维护出现过的商品列表
// 保证传入的时候, 商品不存在
func CreateNewItem(itemName string, itemPrice float32, itemDiscount float32) (ItemModel, error) {
	if item, exists := ItemTable[itemName]; exists {
		return item, nil
	}

	newItem := ItemModel{
		Id:       0,
		ItemName: itemName,
		Price:    itemPrice,
		Discount: itemDiscount,
	}

	ItemTable[itemName] = newItem

	return newItem, nil
}
