package controllers

type Order struct {
	ID            int    `form:"id" json:"id"`
	CustomerEmail string `form:"customer_email" json:"customer_email"`
	Waktu         string `form:"waktu" json:"waktu"`
	Alamat        string `form:"alamat" json:"alamat"`
	Status        int    `form:"status" json:"status"`
	Rating        int    `form:"rating" json:"rating"`
}

type OrderResponse struct {
	Message   string  `form:"message" json:"message"`
	DataOrder []Order `form:"dataOrder" json:"dataOrder"`
}

// type OrderFullResponse struct {
// 	Message         string        `form:"message" json:"message"`
// 	DataOrder       []Order       `form:"dataOrder" json:"dataOrder"`
// 	DataUser        []User        `form:"dataUser" json:"dataUser"`
// 	DataOrderDetail []OrderDetail `form:"dataOrderDetail" json:"dataOrderDetail"`
// 	DataPayment     []Payment     `form:"dataPayment" json:"dataPayment"`
// }
