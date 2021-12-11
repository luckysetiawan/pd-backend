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
			log.Fatal(err.Error())
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
			log.Fatal(err.Error())
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

// Get Active Orders (status dimasak)
func GetActiveOrders(c *gin.Context) {
	db := connect()
	defer db.Close()
	fmt.Println("a")
	query := "SELECT * FROM `order` WHERE status='0';"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order model.Order
	var orders []model.Order

	var orderdetail model.OrderDetail
	var orderdetails []model.OrderDetail

	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
		} else {
			orders = append(orders, order)

			detailquery := "SELECT * FROM `orderdetail` WHERE order_id='" + strconv.Itoa(order.ID) + "';"

			rows2, err2 := db.Query(detailquery)
			if err2 != nil {
				log.Println(err2)
			}

			// var orderdetail model.OrderDetail
			// var orderdetails []model.OrderDetail
			for rows2.Next() {
				if err := rows2.Scan(&orderdetail.ID, &orderdetail.PizzaID, &orderdetail.OrderID, &orderdetail.Quantity, &orderdetail.TotalHarga); err != nil {
					log.Fatal(err.Error())
				} else {
					orderdetails = append(orderdetails, orderdetail)
				}
			}

		}
	}

	// query := "SELECT * FROM `orderdetail` WHERE order_id='" + order.ID + "';"

	var Response model.ActiveOrderResponse
	if err == nil {
		Response.Message = "Get Active Order Success"
		Response.ActiveOrder = orders
		Response.DetailOrder = orderdetails
		sendActiveOrderSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Order Query Error"
		fmt.Print(err)
		sendActiveOrderErrorResponse(c, Response)
	}
}

// Get Status by Id
func GetStatus(c *gin.Context) {
	db := connect()
	defer db.Close()

	CustomerEmail := c.PostForm("customer_email")

	query := "SELECT status FROM `order` WHERE customer_email='" + CustomerEmail + "';"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order int
	var orders []int
	for rows.Next() {
		if err := rows.Scan(&order); err != nil {
			log.Fatal(err.Error())
		} else {
			orders = append(orders, order)
		}
	}

	// var order model.Order
	// if err := rows.Scan(&order.Status); err != nil {
	// 	log.Fatal(err.Error)
	// }

	var Response model.StatusResponse
	if err == nil {
		Response.Message = "Get Order Success"
		Response.Status = orders
		sendStatusSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Order Query Error"
		fmt.Print(err.Error())
		log.Fatal(err.Error())
		sendStatusErrorResponse(c, Response)
	}
}

// Get Response untuk fungsi lain
func GetDataResponse(ID string, table string, c *gin.Context) []model.Order {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM `" + table + "` WHERE id='" + ID + "';"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order model.Order
	var orders []model.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
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

	//Insert Customer dari DB
	CustomerEmail := c.PostForm("customer_email")
	Nama := c.PostForm("nama")
	NoTelp := c.PostForm("no_telp")

	db.Exec("INSERT INTO customer (email, nama, no_telp) values (?,?,?)",
		CustomerEmail,
		Nama,
		NoTelp,
	)

	Waktu := time.Now()
	Alamat := c.PostForm("alamat")
	Status, _ := strconv.Atoi("0")
	_, errQuery := db.Exec("INSERT INTO `order`(customer_email, waktu, alamat, status, rating) values (?,?,?,?,0)",
		CustomerEmail,
		Waktu,
		Alamat,
		Status,
	)

	// Get ID order dari DB
	query := "SELECT ID FROM `order` WHERE customer_email='" + CustomerEmail + "' AND status=0;"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order model.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID); err != nil {
			log.Fatal(err.Error())
		}
	}

	Pizza := c.PostForm("pizza_id")
	Quantity := c.PostForm("quantity")
	PizzaArr := strings.Split(Pizza, ",")
	QuantityArr := strings.Split(Quantity, ",")
	var TotalHarga int
	TotalPembayaran := 0
	//Tambah Insert Total Harga
	for i := 0; i < len(PizzaArr); i++ {
		//Get harga dari DB
		query := "SELECT harga FROM pizza WHERE id=" + PizzaArr[i] + ";"

		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
		}

		var pizza model.Menu
		for rows.Next() {
			if err := rows.Scan(&pizza.Harga); err != nil {
				log.Fatal(err.Error())
			}
		}
		//Perhitungan Total Harga
		TempQuantity, _ := strconv.Atoi(QuantityArr[i])
		TotalHarga = pizza.Harga * int(TempQuantity)
		TotalPembayaran += TotalHarga
		fmt.Println(TotalHarga)
		fmt.Println(pizza.Harga)
		fmt.Println(TempQuantity)
		fmt.Println(TotalPembayaran)
		//Insert total harga ke DB
		db.Exec("INSERT INTO orderdetail (pizza_id, order_id, quantity, total_harga) values (?,?,?,?)",
			PizzaArr[i],
			order.ID,
			QuantityArr[i],
			TotalHarga,
		)

	}

	_, errQuery = db.Exec("INSERT INTO payment(order_id, status_pembayaran, total_pembayaran) values (?,0,?)",
		order.ID,
		TotalPembayaran,
	)

	var response model.OrderResponse
	if errQuery == nil {
		response.Message = "Insert Order Success"
		response.DataOrder = GetDataResponse(strconv.Itoa(order.ID), "order", c)
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
		response.DataOrder = GetDataResponse(ID, "order", c)
		sendOrderSuccessResponse(c, response)
	} else {
		response.Message = "Update Order Failed Error"
		fmt.Print(errQuery)
		sendOrderErrorOResponse(c, response)
	}
}

// // Update Rating
// func UpdateRating(c *gin.Context) {
// 	db := connect()
// 	defer db.Close()

// 	var order model.Order
// 	order.ID, _ = strconv.Atoi(c.Param("order_id"))
// 	order.Rating, _ = strconv.Atoi(c.PostForm("rating"))

// 	_, errQuery := db.Exec("UPDATE `order` SET rating = ? WHERE id=?",
// 		order.Rating,
// 		order.ID,
// 	)

// 	var response model.RatingResponse
// 	if errQuery == nil {
// 		response.Message = "Update Rating Success"
// 		response.Rating = order.Rating
// 		sendRatingSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Update Rating Failed Error"
// 		fmt.Print(errQuery)
// 		sendRatingErrorResponse(c, response)
// 	}
// }

// BUAT FUNC PAYMENT TAPI TERNYATA UDAH ADA JADI GA DIPAKE :(
// // Update Payment
// func UpdatePayment(c *gin.Context) {
// 	db := connect()
// 	defer db.Close()

// 	CustomerEmail := c.PostForm("customer_email")
// 	// Get ID order dari DB
// 	query := "SELECT ID FROM `order` WHERE customer_email='" + CustomerEmail + "' AND status=0;"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var order model.Order
// 	for rows.Next() {
// 		if err := rows.Scan(&order.ID); err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	}

// 	_, errQuery := db.Exec("UPDATE payment SET status_pembayaran = 1, waktu_pembayaran =?  WHERE id=?;",
// 		time.Now().Format("2006-01-02 15:04:05"),
// 		order.ID,
// 	)

// 	var response model.OrderResponse
// 	if errQuery == nil {
// 		response.Message = "Update Order Success"
// 		response.DataOrder = GetDataResponse(strconv.Itoa(order.ID), "payment", c)
// 		sendOrderSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Update Order Failed Error"
// 		fmt.Print(errQuery)
// 		sendOrderErrorOResponse(c, response)
// 	}
// }

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
