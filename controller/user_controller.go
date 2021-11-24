package controllers

import (
	"log"
	"net/http"

	model "pd-backend/model"

	"github.com/gin-gonic/gin"
)

// Get All Users
func GetAllUsers(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	nama := c.Query("nama")
	if nama != "" {
		query += " WHERE nama='" + nama + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var user model.User
	var users []model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Nama, &user.Email, &user.NoTelp, &user.Password); err != nil {
			log.Fatal(err.Error)
		} else {
			users = append(users, user)
		}
	}

	var response model.UserResponse
	if err == nil {
		response.Message = "Get User Success"
		response.Data = users
		sendUserSuccessresponse(c, response)
	} else {
		response.Message = "Get User Query Error"
		sendUserErrorResponse(c, response)
	}
}

// Insert User
func Registrasi(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	email := c.PostForm("email")
	notelp := c.PostForm("notelp")
	password := c.PostForm("password")

	_, errQuery := db.Exec("INSERT INTO users(nama, email, notelp, password) values (?,?,?,?)",
		nama,
		email,
		notelp,
		password,
	)

	var response model.UserResponse
	if errQuery == nil {
		response.Message = "Insert User Success"
		sendUserSuccessresponse(c, response)
	} else {
		response.Message = "Insert User Failed Error"
		sendUserErrorResponse(c, response)
	}
}

// Update User
func UpdateUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	email := c.PostForm("email")
	notelp := c.PostForm("notelp")
	password := c.PostForm("password")
	userId := c.Param("user_id")

	rows, _ := db.Query("SELECT * FROM users WHERE id='" + userId + "'")
	var user model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Nama, &user.Email, &user.Password, &user.NoTelp); err != nil {
			log.Fatal(err.Error())
		}
	}

	// Jika kosong dimasukkan nilai lama
	if nama == "" {
		nama = user.Nama
	}

	if email == "" {
		email = user.Email
	}

	if notelp == "" {
		notelp = user.NoTelp
	}

	_, errQuery := db.Exec("UPDATE users SET nama = ?, email = ?, notelp = ?, password = ? WHERE id=?",
		nama,
		email,
		notelp,
		userId,
		password,
	)

	var response model.UserResponse
	if errQuery == nil {
		response.Message = "Update User Success"
		sendUserSuccessresponse(c, response)
	} else {
		response.Message = "Update User Failed Error"
		sendUserErrorResponse(c, response)
	}
}

// Delete User...
func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	userId := c.Param("user_id")

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		userId,
	)

	var response model.UserResponse
	if errQuery == nil {
		response.Message = "Delete User Success"
		sendUserSuccessresponse(c, response)
	} else {
		response.Message = "Delete User Failed Error"
		sendUserErrorResponse(c, response)
	}
}

func sendUserSuccessresponse(c *gin.Context, ur model.UserResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendUserErrorResponse(c *gin.Context, ur model.UserResponse) {
	c.JSON(http.StatusBadRequest, ur)
}
