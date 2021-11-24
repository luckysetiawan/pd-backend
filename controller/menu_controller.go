package controllers

import (
	"log"
	"net/http"
	"strconv"

	model "pd-backend/model"

	"github.com/gin-gonic/gin"
)

// Get All Pizzas
func GetAllPizzas(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM pizza"

	nama := c.Query("nama")
	if nama != "" {
		query += " WHERE nama='" + nama + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var pizza model.Pizza
	var pizzas []model.Pizza
	for rows.Next() {
		if err := rows.Scan(&pizza.ID, &pizza.Nama, &pizza.Harga, &pizza.Varian); err != nil {
			log.Fatal(err.Error)
		} else {
			pizzas = append(pizzas, pizza)
		}
	}

	var response model.PizzaResponse
	if err == nil {
		response.Message = "Get Pizza Success"
		response.Data = pizzas
		sendPizzaSuccessresponse(c, response)
	} else {
		response.Message = "Get Pizza Query Error"
		sendPizzaErrorResponse(c, response)
	}
}

// Insert Pizza
func InsertPizza(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	harga, _ := strconv.Atoi(c.PostForm("harga"))
	varian := c.PostForm("varian")

	_, errQuery := db.Exec("INSERT INTO pizza(nama, harga, varian) values (?,?,?)",
		nama,
		harga,
		varian,
	)

	var response model.PizzaResponse
	if errQuery == nil {
		response.Message = "Insert Pizza Success"
		sendPizzaSuccessresponse(c, response)
	} else {
		response.Message = "Insert Pizza Failed Error"
		sendPizzaErrorResponse(c, response)
	}
}

// Update Pizza
func UpdatePizza(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	harga, _ := strconv.Atoi(c.PostForm("harga"))
	varian := c.PostForm("varian")
	pizzaId := c.Param("pizza_id")

	rows, _ := db.Query("SELECT * FROM pizza WHERE id='" + pizzaId + "'")
	var pizza model.Pizza
	for rows.Next() {
		if err := rows.Scan(&pizza.ID, &pizza.Nama, &pizza.Harga, &pizza.Varian); err != nil {
			log.Fatal(err.Error())
		}
	}

	// Jika kosong dimasukkan nilai lama
	if nama == "" {
		nama = pizza.Nama
	}

	if harga == 0 {
		harga = pizza.Harga
	}

	if varian == "" {
		varian = pizza.Varian
	}

	_, errQuery := db.Exec("UPDATE pizza SET nama = ?, harga = ?, varian = ? WHERE id=?",
		nama,
		harga,
		varian,
		pizzaId,
	)

	var response model.PizzaResponse
	if errQuery == nil {
		response.Message = "Update Pizza Success"
		sendPizzaSuccessresponse(c, response)
	} else {
		response.Message = "Update Pizza Failed Error"
		sendPizzaErrorResponse(c, response)
	}
}

// Delete Pizza
func DeletePizza(c *gin.Context) {
	db := connect()
	defer db.Close()

	pizzaId := c.Param("pizza_id")

	_, errQuery := db.Exec("DELETE FROM pizza WHERE id=?",
		pizzaId,
	)

	var response model.PizzaResponse
	if errQuery == nil {
		response.Message = "Delete Pizza Success"
		sendPizzaSuccessresponse(c, response)
	} else {
		response.Message = "Delete Pizza Failed Error"
		sendPizzaErrorResponse(c, response)
	}
}

func sendPizzaSuccessresponse(c *gin.Context, ur model.PizzaResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendPizzaErrorResponse(c *gin.Context, ur model.PizzaResponse) {
	c.JSON(http.StatusBadRequest, ur)
}
