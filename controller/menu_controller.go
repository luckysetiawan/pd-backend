package controllers

import (
	"log"
	"net/http"
	"strconv"

	model "pd-backend/model"

	"github.com/gin-gonic/gin"
)

// Get All Menus
func GetAllMenus(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM menus"

	nama := c.Query("nama")
	if nama != "" {
		query += " WHERE nama='" + nama + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var menu model.Menu
	var menus []model.Menu
	for rows.Next() {
		if err := rows.Scan(&menu.ID, &menu.Nama, &menu.Deskripsi, &menu.Harga, &menu.Gambar, &menu.Varian); err != nil {
			log.Fatal(err.Error())
		} else {
			menus = append(menus, menu)
		}
	}

	var response model.MenuResponse
	if err == nil {
		response.Message = "Get Menu Success"
		response.Data = menus
		sendMenuSuccessresponse(c, response)
	} else {
		response.Message = "Get Menu Query Error"
		sendMenuErrorResponse(c, response)
	}
}

// Insert Menu
func InsertMenu(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	deskripsi := c.PostForm("deskripsi")
	harga, _ := strconv.Atoi(c.PostForm("harga"))
	gambar := c.PostForm("gambar")
	varian := c.PostForm("varian")

	_, errQuery := db.Exec("INSERT INTO menus(nama, deskripsi, harga, gambar, varian) values (?,?,?,?,?)",
		nama,
		deskripsi,
		harga,
		gambar,
		varian,
	)

	var response model.MenuResponse
	if errQuery == nil {
		response.Message = "Insert Menu Success"
		sendMenuSuccessresponse(c, response)
	} else {
		response.Message = "Insert Menu Failed Error"
		sendMenuErrorResponse(c, response)
	}
}

// Update Menu
func UpdateMenu(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	deskripsi := c.PostForm("deskripsi")
	harga, _ := strconv.Atoi(c.PostForm("harga"))
	gambar := c.PostForm("gambar")
	varian := c.PostForm("varian")
	menuId := c.Param("menu_id")

	rows, _ := db.Query("SELECT * FROM menus WHERE id='" + menuId + "'")
	var menu model.Menu
	for rows.Next() {
		if err := rows.Scan(&menu.ID, &menu.Nama, &menu.Deskripsi, &menu.Harga, &menu.Gambar, &menu.Varian); err != nil {
			log.Fatal(err.Error())
		}
	}

	// Jika kosong dimasukkan nilai lama
	if nama == "" {
		nama = menu.Nama
	}

	if deskripsi == "" {
		deskripsi = menu.Deskripsi
	}

	if harga == 0 {
		harga = menu.Harga
	}

	if gambar == "" {
		gambar = menu.Gambar
	}

	if varian == "" {
		varian = menu.Varian
	}

	_, errQuery := db.Exec("UPDATE menus SET nama = ?, harga = ?, varian = ? WHERE id=?",
		nama,
		harga,
		varian,
		menuId,
	)

	var response model.MenuResponse
	if errQuery == nil {
		response.Message = "Update Menu Success"
		sendMenuSuccessresponse(c, response)
	} else {
		response.Message = "Update Menu Failed Error"
		sendMenuErrorResponse(c, response)
	}
}

// Delete Menu
func DeleteMenu(c *gin.Context) {
	db := connect()
	defer db.Close()

	menuId := c.Param("menu_id")

	_, errQuery := db.Exec("DELETE FROM menus WHERE id=?",
		menuId,
	)

	var response model.MenuResponse
	if errQuery == nil {
		response.Message = "Delete Menu Success"
		sendMenuSuccessresponse(c, response)
	} else {
		response.Message = "Delete Menu Failed Error"
		sendMenuErrorResponse(c, response)
	}
}

func sendMenuSuccessresponse(c *gin.Context, ur model.MenuResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendMenuErrorResponse(c *gin.Context, ur model.MenuResponse) {
	c.JSON(http.StatusBadRequest, ur)
}
