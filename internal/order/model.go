package order

type OrderModel struct {
	Id        uint   `json:"id"`
	ItemCount int    `json:"item_count`
	ItemName  string `json:"item_name"`
}
