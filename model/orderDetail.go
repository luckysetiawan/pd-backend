package controllers

type OrderDetail struct {
	ID         int     `form:"id" json:"id"`
	PizzaID    int     `form:"pizza_id" json:"pizza_id"`
	OrderID    int     `form:"order_id" json:"order_id"`
	Quantity   int     `form:"quantity" json:"quantity"`
	TotalHarga float64 `form:"total_harga" json:"total_harga"`
}

type OrderDetailResponse struct {
	Message string        `form:"message" json:"message"`
	Data    []OrderDetail `form:"data" json:"data"`
}
