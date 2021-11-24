package controllers

type Menu struct {
	ID     int    `form:"id" json:"id"`
	Nama   string `form:"nama" json:"nama"`
	Harga  int    `form:"harga" json:"harga"`
	Varian string `form:"varian" json:"varian"`
}

type MenuResponse struct {
	Message string `form:"message" json:"message"`
	Data    []Menu `form:"data" json:"data"`
}
