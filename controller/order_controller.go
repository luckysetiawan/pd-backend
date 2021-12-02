package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	model "pd-backend/model"

	"github.com/gin-gonic/gin"
)

// Get All Orders
func GetAllOrders(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM `order`;"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order model.Order
	var orders []model.Order

	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error)
		} else {
			orders = append(orders, order)
		}
	}

	var Response model.OrderResponse
	if err == nil {
		Response.Message = "Get Order Success"
		Response.DataOrder = orders
		sendOrderSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Order Query Error"
		fmt.Print(err)
		sendOrderErrorOResponse(c, Response)
	}
}

// Get Orders by Id
func GetOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	ID := c.Param("order_id")

	query := "SELECT * FROM `order` WHERE id='" + ID + "';"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order model.Order
	var orders []model.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error)
		} else {
			orders = append(orders, order)
		}
	}

	var Response model.OrderResponse
	if err == nil {
		Response.Message = "Get Order Success"
		Response.DataOrder = orders
		sendOrderSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Order Query Error"
		fmt.Print(err)
		sendOrderErrorOResponse(c, Response)
	}
}

// Get Orders by Id
func GetStatus(c *gin.Context) {
	db := connect()
	defer db.Close()

	ID := c.Param("order_id")

	query := "SELECT status FROM `order` WHERE id='" + ID + "';"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order model.Order
	for rows.Next() {
		if err := rows.Scan(&order.Status); err != nil {
			log.Fatal(err.Error)
		}
	}

	// var order model.Order
	// if err := rows.Scan(&order.Status); err != nil {
	// 	log.Fatal(err.Error)
	// }

	var Response model.StatusResponse
	if err == nil {
		Response.Message = "Get Order Success"
		Response.Status = order.Status
		sendStatusSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Order Query Error"
		fmt.Print(err.Error())
		log.Fatal(err.Error())
		sendStatusErrorResponse(c, Response)
	}
}

// Get Response untuk fungsi lain
func GetDataResponse(ID string, c *gin.Context) []model.Order {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM `order` WHERE id='" + ID + "';"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order model.Order
	var orders []model.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error)
		} else {
			orders = append(orders, order)
		}
	}
	return orders
}

// Insert Order
func InsertOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	CustomerEmail := c.PostForm("customer_email")
	Nama := c.PostForm("nama")
	NoTelp := c.PostForm("no_telp")

	db.Exec("INSERT INTO customer (email, nama, no_telp) values (?,?,?)",
		CustomerEmail,
		Nama,
		NoTelp,
	)

	Pizza := c.PostForm("pizza_id")
	quantity := c.PostForm("quantity")
	pizzaArr := strings.Split(Pizza, ",")
	quantityArr := strings.Split(quantity, ",")

	for i := 0; i < len(pizzaArr); i++ {
		db.Exec("INSERT INTO orderdetail (pizza_id, quantity) values (?,?)",
			pizzaArr[i],
			quantityArr[i],
		)

	}

	ID, _ := strconv.Atoi(c.PostForm("id"))
	Waktu := time.Now()
	Alamat := c.PostForm("alamat")
	Status, _ := strconv.Atoi("0")
	Rating, _ := strconv.Atoi(c.PostForm("rating"))
	_, errQuery := db.Exec("INSERT INTO `order`(id, customer_email, waktu, alamat, status, rating) values (?,?,?,?,?,?)",
		ID,
		CustomerEmail,
		Waktu,
		Alamat,
		Status,
		Rating,
	)

	// var orderID model.Order
	// if c.PostForm("id") == "" {
	// 	getID, errID := db.Query("SELECT id FROM `order` WHERE waktu=" + time.Now() + ";")
	// 	if err := getID.Scan(&orderID.ID); errID != nil {
	// 		log.Fatal(err.Error)
	// 		log.Println(err.Error)
	// 		fmt.Println(ID)
	// 	}
	// } else {
	// 	orderID.ID = ID
	// }

	// _, errQuery := db.Exec("INSERT INTO payment(order_id,customer_email, status_pembayaran) values (?,?,0)",
	// 	orderID,
	// 	CustomerEmail,
	// )

	var response model.OrderResponse
	if errQuery == nil {
		response.Message = "Insert Order Success"
		response.DataOrder = GetDataResponse(strconv.Itoa(ID), c)
		sendOrderSuccessResponse(c, response)
	} else {
		response.Message = "Insert Order Failed"
		fmt.Print(errQuery)
		sendOrderErrorOResponse(c, response)
	}
}

// Update Order
func UpdateOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	ID := c.Param("order_id")
	CustomerEmail := c.PostForm("customer_email")
	Waktu := c.PostForm("waktu")
	Alamat := c.PostForm("alamat")
	Status := c.PostForm("status")
	Rating := c.PostForm("rating")

	rows, _ := db.Query("SELECT * FROM `order` WHERE id='" + ID + "'")
	var order model.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
		}
	}

	// Jika kosong dimasukkan nilai lama
	if ID == "" {
		ID = strconv.Itoa(order.ID)
	}

	if CustomerEmail == "" {
		CustomerEmail = order.CustomerEmail
	}

	if Waktu == "" {
		Waktu = order.Waktu
	}

	if Alamat == "" {
		Alamat = order.Alamat
	}

	if Status == "" {
		Status = strconv.Itoa(order.Status)
	}
	if Rating == "" {
		Rating = strconv.Itoa(order.Rating)
	}
	_, errQuery := db.Exec("UPDATE `order` SET customer_email = ?, waktu = ?, alamat = ?, status = ?, rating = ? WHERE id=?",
		CustomerEmail,
		Waktu,
		Alamat,
		Status,
		Rating,
		ID,
	)

	var response model.OrderResponse
	if errQuery == nil {
		response.Message = "Update Order Success"
		response.DataOrder = GetDataResponse(ID, c)
		sendOrderSuccessResponse(c, response)
	} else {
		response.Message = "Update Order Failed Error"
		fmt.Print(errQuery)
		sendOrderErrorOResponse(c, response)
	}
}

// Delete Order...
func DeleteOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	ID := c.Param("order_id")

	_, errQuery := db.Exec("DELETE FROM `order` WHERE id=?",
		ID,
	)

	var response model.OrderResponse
	if errQuery == nil {
		response.Message = "Delete Order Success"
		sendOrderSuccessResponse(c, response)
	} else {
		response.Message = "Delete Order Failed Error"
		fmt.Print(errQuery)
		sendOrderErrorOResponse(c, response)
	}
}
