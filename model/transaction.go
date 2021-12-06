package controllers

type Transaction struct {
	// DataOrder   []Order       `form:"dataOrder" json:"dataOrder"`
	DataOrder   Order         `form:"dataOrder" json:"dataOrder"`
	DetailOrder []OrderDetail `form:"detailOrder" json:"detailOrder"`
}

type TransactionResponse struct {
	Message string        `form:"message" json:"message"`
	Data    []Transaction `form:"data" json:"data"`
}
