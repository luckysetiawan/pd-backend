package controllers

type User struct {
	ID       int    `form:"id" json:"id"`
	Nama     string `form:"nama" json:"nama"`
	Email    string `form:"email" json:"email"`
	NoTelp   string `form:"notelp" json:"notelp"`
	Password string `form:"password" json:"password"`
}

type Pizza struct {
	ID     int    `form:"id" json:"id"`
	Nama   string `form:"nama" json:"nama"`
	Harga  int    `form:"harga" json:"harga"`
	Varian string `form:"varian" json:"varian"`
}

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
	Pizza        int     `form:"pizza" json:"pizza"`
	Alamat       string  `form:"alamat" json:"alamat"`
	Rating       int     `form:"rating" json:"rating"`
	Quantity     int     `form:"quantity" json:"quantity"`
	TotalPesanan float64 `form:"totalPesanan" json:"totalPesanan"`
}

type Payment struct {
	Invoice          string  `form:"invoice" json:"invoice"`
	StatusPembayaran int     `form:"statusPembayaran" json:"statusPembayaran"`
	TotalHarga       float64 `form:"totalHarga" json:"totalHarga"`
}

type UserResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}

type PizzaResponse struct {
	Message string  `form:"message" json:"message"`
	Data    []Pizza `form:"data" json:"data"`
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

type PaymentResponse struct {
	Message string    `form:"message" json:"message"`
	Data    []Payment `form:"data" json:"data"`
}
