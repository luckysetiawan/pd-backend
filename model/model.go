package controllers

type User struct {
	ID     int    `form:"id" json:"id"`
	Nama   string `form:"nama" json:"nama"`
	Email  string `form:"email" json:"email"`
	NoTelp string `form:"notelp" json:"notelp"`
}

type UserResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}

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
