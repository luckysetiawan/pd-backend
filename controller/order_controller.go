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

	query := "SELECT * FROM `order` WHERE status = 0"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var transaction model.Transaction
	var transactions []model.Transaction
	var order model.Order
	var orders []model.Order
	var orderdetails []model.OrderDetail

	for rows.Next() { // loop order
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
		} else {
			if order.Status == 0 {
				orders = append(orders, order)

				detailquery := "SELECT * FROM `orderdetail` WHERE order_id = '" + strconv.Itoa(order.ID) + "';"

				rows2, err2 := db.Query(detailquery)
				if err2 != nil {
					log.Println(err2)
				}

				for rows2.Next() { // loop orderdetail
					var orderdetail model.OrderDetail

					fmt.Println(order.ID, orderdetail.OrderID)
					if err := rows2.Scan(&orderdetail.ID, &orderdetail.PizzaID, &orderdetail.OrderID, &orderdetail.Quantity, &orderdetail.TotalHarga); err != nil {
						log.Fatal(err.Error())
					} else {
						fmt.Println(orderdetail)
						orderdetails = append(orderdetails, orderdetail)
					}
				}
			}

			fmt.Println()
			fmt.Println("id", order.ID, "orderdetails", orderdetails)
			transaction.DetailOrder = append(transaction.DetailOrder, orderdetails...)
			orderdetails = nil
			fmt.Println("trans", transaction.DetailOrder)
		}
		transaction.DataOrder = order
		transactions = append(transactions, transaction)
		transaction.DetailOrder = nil
	}

	var Response model.TransactionResponse
	if err == nil {
		Response.Message = "Get Active Order Success"
		Response.Data = transactions
		sendTransactionSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Active Order Query Error"
		fmt.Print(err)
		sendTransactionErrorResponse(c, Response)
	}
}

// Get Rincian Pesanan Customer(staff)
func GetOrderDetail(c *gin.Context) {
	db := connect()
	defer db.Close()

// 	OrderID := c.PostForm("order_id")
	OrderID := c.Param("order_id")

	query := "SELECT * FROM `orderdetail` WHERE order_id='" + OrderID + "';"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var orderDetail model.OrderDetail
	var orderDetails []model.OrderDetail
	for rows.Next() {
		if err := rows.Scan(&orderDetail.ID, &orderDetail.PizzaID, &orderDetail.OrderID,
			&orderDetail.Quantity, &orderDetail.TotalHarga); err != nil {
			log.Fatal(err.Error())
		} else {
			orderDetails = append(orderDetails, orderDetail)
		}
	}

	var Response model.OrderDetailResponse
	if err == nil {
		Response.Message = "Get Order Success"
		Response.Data = orderDetails
		sendOrderDetailSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Order Query Error"
		fmt.Print(err)
		sendOrderDetailErrorResponse(c, Response)
	}
}

// Get Status by Id
func GetStatus(c *gin.Context) {
	db := connect()
	defer db.Close()

	CustomerEmail := c.Request.URL.Query()["customer_email"][0]

	query := "SELECT status FROM `order` WHERE customer_email='" + CustomerEmail + "';"
	log.Println(CustomerEmail)

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
		fmt.Println("TotalHarga:", TotalHarga)
		fmt.Println("pizza.Harga:", pizza.Harga)
		fmt.Println("TempQuantity:", TempQuantity)
		fmt.Println("TotalPembayaran:", TotalPembayaran)
		fmt.Println(".")
		fmt.Println("PizzaArr[i]:", PizzaArr[i])
		fmt.Println("order.ID:", order.ID)
		fmt.Println("QuantityArr[i]:", QuantityArr[i])
		//Insert total harga ke DB
		_, errQuery = db.Exec("INSERT INTO orderdetail (pizza_id, order_id, quantity, total_harga) values (?,?,?,?)",
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
