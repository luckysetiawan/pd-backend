package controllers

type Menu struct {
	ID        int    `form:"id" json:"id"`
	Nama      string `form:"nama" json:"nama"`
	Deskripsi string `form:"deskripsi" json:"deskripsi"`
	Harga     int    `form:"harga" json:"harga"`
	Gambar    string `form:"gambar" json:"gambar"`
}

type MenuResponse struct {
	Message string `form:"message" json:"message"`
	Data    []Menu `form:"data" json:"data"`
}
