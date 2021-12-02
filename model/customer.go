package controllers

type Customer struct {
	Email      string `form:"email" json:"email"`
	Nama       string `form:"nama" json:"nama"`
	NoTelp     string `form:"no_telp" json:"no_telp"`
	JumlahBeli int    `form:"jumlah_pembelian" json:"jumlah_pembelian"`
}

type CustomerResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
