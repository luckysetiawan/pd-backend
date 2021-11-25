package controllers

type Payment struct {
	Invoice          string  `form:"invoice" json:"invoice"`
	StatusPembayaran int     `form:"statusPembayaran" json:"statusPembayaran"`
	TotalHarga       float64 `form:"totalHarga" json:"totalHarga"`
}

type PaymentResponse struct {
	Message string    `form:"message" json:"message"`
	Data    []Payment `form:"data" json:"data"`
}
