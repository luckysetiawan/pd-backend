package controllers

type User struct {
	ID       int    `form:"id" json:"id"`
	Nama     string `form:"nama" json:"nama"`
	Email    string `form:"email" json:"email"`
	NoTelp   string `form:"notelp" json:"notelp"`
	Password string `form:"password" json:"password"`
}

type UserResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
