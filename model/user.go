package controllers

type User struct {
	ID       int    `form:"id" json:"id"`
	Nama     string `form:"nama" json:"nama"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	NoTelp   string `form:"notelp" json:"notelp"`
	Position string `form:"position" json:"position"`
}

type UserResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
