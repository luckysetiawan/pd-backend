package controllers

type Order struct {
	ID            int    `form:"id" json:"id"`
	IdCustomer    int    `form:"idCustomer" json:"idCustomer"`
	IdOrderDetail int    `form:"idOrderDetail" json:"idOrderDetail"`
	Invoice       string `form:"invoice" json:"invoice"`
	Waktu         string `form:"waktu" json:"waktu"`
	Alamat        string `form:"alamat" json:"alamat"`
	Status        int    `form:"status" json:"status"`
}

type OrderDetail struct {
	ID           int     `form:"id" json:"id"`
	Menu         int     `form:"menu" json:"menu"`
	Rating       int     `form:"rating" json:"rating"`
	Quantity     int     `form:"quantity" json:"quantity"`
	TotalPesanan float64 `form:"totalPesanan" json:"totalPesanan"`
}

type OrderResponse struct {
	Message   string  `form:"message" json:"message"`
	DataOrder []Order `form:"dataOrder" json:"dataOrder"`
}

type OrderFullResponse struct {
	Message         string        `form:"message" json:"message"`
	DataOrder       []Order       `form:"dataOrder" json:"dataOrder"`
	DataUser        []User        `form:"dataUser" json:"dataUser"`
	DataOrderDetail []OrderDetail `form:"dataOrderDetail" json:"dataOrderDetail"`
	DataPayment     []Payment     `form:"dataPayment" json:"dataPayment"`
}

type OrderDetailResponse struct {
	Message string        `form:"message" json:"message"`
	Data    []OrderDetail `form:"data" json:"data"`
}
