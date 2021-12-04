package controllers

type Payment struct {
	ID               int    `form:"id" json:"id"`
	OrderID          int    `form:"order_id" json:"order_id"`
	StatusPembayaran int    `form:"status_pembayaran" json:"status_pembayaran"`
	TotalPembayaran  int    `form:"total_pembayaran" json:"total_pembayaran"`
	WaktuPembayaran  string `form:"waktu_pembayaran" json:"waktu_pembayaran"`
}

type PaymentResponse struct {
	Message string    `form:"message" json:"message"`
	Data    []Payment `form:"data" json:"data"`
}

type PendapatanResponse struct {
	Message string `form:"message" json:"message"`
	Total   int    `form:"total" json:"total"`
}

type PizzaTerjualResponse struct {
	Message string `form:"message" json:"message"`
	Total   int    `form:"total" json:"total"`
}
